import grpc
import logging
import asyncio

from rasa.core.agent import Agent
from rasa.shared.data import TrainingType
from rasa.shared.utils.cli import print_error

import nlp_pb2
import nlp_pb2_grpc

from rasa.model import get_local_model
from rasa.shared.constants import DEFAULT_MODELS_PATH
from rasa.engine.storage.local_model_storage import LocalModelStorage


class ExampleNLPServicer(nlp_pb2_grpc.NLPServicer):
    def __init__(self, model_path):
        self.agent = Agent.load(model_path)

    async def Parse(self, request, context):
        d = await self.agent.parse_message(request.text)
        res = nlp_pb2.Result(text=d.get("text", request.text))

        intent = d.get('intent', None)
        if intent is not None:
            res.intent.CopyFrom(nlp_pb2.Intent(name=intent.get("name"),
                                               confidence=intent.get("confidence", 0.0)))

        entities = d.get('entities', [])
        for entity in entities:
            res.entities.append(nlp_pb2.Entity(end=entity.get("end"),
                                               name=entity.get("entity"),
                                               start=entity.get("start"),
                                               value=entity.get("value")))

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
