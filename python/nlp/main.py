import grpc
import logging
import nlp_pb2
import nlp_pb2_grpc

from concurrent import futures
from natasha import (
    Doc,
    Segmenter,
    MorphVocab,
    NewsEmbedding,
    NewsNERTagger,
    NewsMorphTagger,
    NewsSyntaxParser,
)

emb = NewsEmbedding()
segmenter = Segmenter()
ner_tagger = NewsNERTagger(emb)
morph_vocab = MorphVocab()
morph_tagger = NewsMorphTagger(emb)
syntax_parser = NewsSyntaxParser(emb)


def to_nlp_pb2_doc_token(token):
    return nlp_pb2.DocToken(
        head_id=token.head_id,
        lemma=token.lemma,
        start=token.start,
        stop=token.stop,
        text=token.text,
        pos=token.pos,
        rel=token.rel,
        id=token.id,
    )


def to_nlp_pb2_doc_span(span):
    return nlp_pb2.DocSpan(
        tokens=map(to_nlp_pb2_doc_token, span.tokens),
        normal=span.normal,
        start=span.start,
        stop=span.stop,
        text=span.text,
        type=span.type,
    )


def to_nlp_pb2_doc_sent(sent):
    return nlp_pb2.DocSent(
        tokens=map(to_nlp_pb2_doc_token, sent.tokens),
        spans=map(to_nlp_pb2_doc_span, sent.spans),
        start=sent.start,
        stop=sent.stop,
        text=sent.text,
    )


class NatashaNLPServicer(nlp_pb2_grpc.NLPServicer):
    def Parse(self, request, context):
        doc = Doc(request.text)
        doc.segment(segmenter)
        doc.tag_ner(ner_tagger)
        doc.tag_morph(morph_tagger)
        doc.parse_syntax(syntax_parser)

        for token in doc.tokens:
            token.lemmatize(morph_vocab)

        for span in doc.spans:
            span.normalize(morph_vocab)

        return nlp_pb2.Result(
            tokens=map(to_nlp_pb2_doc_token, doc.tokens),
            sents=map(to_nlp_pb2_doc_sent, doc.sents),
            spans=map(to_nlp_pb2_doc_span, doc.spans),
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    nlp_pb2_grpc.add_NLPServicer_to_server(NatashaNLPServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("NLP gRPC server start on [::]:50051")
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
