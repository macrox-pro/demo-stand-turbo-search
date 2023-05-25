import grpc
import logging
import asyncio

import nlp_pb2
import nlp_pb2_grpc

from natasha import Doc, MorphVocab, Segmenter, NewsEmbedding, NewsNERTagger, NewsMorphTagger, NewsSyntaxParser

from rasa.model import get_local_model
from rasa.core.agent import Agent
from rasa.shared.data import TrainingType
from rasa.shared.utils.cli import print_error
from rasa.shared.constants import DEFAULT_MODELS_PATH
from rasa.engine.storage.local_model_storage import LocalModelStorage


class ExampleNLPServicer(nlp_pb2_grpc.NLPServicer):
    def __init__(self, model_path):
        self.emb = NewsEmbedding()
        self.agent = Agent.load(model_path)
        self.segmenter = Segmenter()
        self.morph_vocab = MorphVocab()
        self.ner_tagger = NewsNERTagger(self.emb)
        self.morph_tagger = NewsMorphTagger(self.emb)
        self.syntax_parser = NewsSyntaxParser(self.emb)

    def normalize(self, text):
        doc = Doc(text)
        doc.segment(self.segmenter)
        doc.tag_morph(self.morph_tagger)
        doc.parse_syntax(self.syntax_parser)
        for token in doc.tokens:
            token.lemmatize(self.morph_vocab)

        doc.tag_ner(self.ner_tagger)

        for span in doc.spans:
            span.normalize(self.morph_vocab)

        simple_tokens = []
        for token in doc.tokens:
            if len(doc.spans) > 0:
                span = next(span for span in doc.spans
                            if span.stop == token.stop and
                            span.text == token.text)

                if span:
                    simple_tokens.append(span.normal)
                    continue

            if token.lemma:
                simple_tokens.append(token.lemma)
            else:
                simple_tokens.append(token.text)

        return " ".join(simple_tokens)

    async def Parse(self, request, context):
        d = await self.agent.parse_message(request.text)
        res = nlp_pb2.Result(text=d.get("text", request.text))

        intent = d.get('intent', None)
        if intent is not None:
            res.intent.CopyFrom(nlp_pb2.Intent(name=intent.get("name"),
                                               confidence=intent.get("confidence", 0.0)))

        entities = d.get('entities', [])
        for entity in entities:
            value = entity.get("value")
            entity_type = entity.get("entity")

            if entity_type == "person":
                value = value.lower().title()

            normal_value = self.normalize(value)

            res.entities.append(nlp_pb2.Entity(end=entity.get("end"),
                                               start=entity.get("start"),
                                               type=entity_type,
                                               value=value,
                                               normal_value=normal_value))

        return res


async def serve():
    path = get_local_model(DEFAULT_MODELS_PATH)
    metadata = LocalModelStorage.metadata_from_archive(path)
    if metadata.training_type == TrainingType.CORE:
        print_error(
            "No NLU model found. Train a model before running the "
            "server using `rasa train nlu`."
        )
        return

    server = grpc.aio.server()
    nlp_pb2_grpc.add_NLPServicer_to_server(ExampleNLPServicer(path), server)
    server.add_insecure_port('[::]:50051')
    await server.start()
    print("NLP gRPC server start on [::]:50051")

    await server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    asyncio.run(serve())
