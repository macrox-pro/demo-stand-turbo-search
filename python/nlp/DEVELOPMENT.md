# Prepare

Install gRPC Tools

```shell
pip3.7 install grpc_tools
```

Install Natasha NLP Tools and Yargy

```shell
pip3.7 install natasha
pip3.7 install yargy
```

Install Rasa

```shell
pip3.7 install rasa
pip3.7 install --upgrade aio-pika
```

# Train NLU Rasa

```shell
rasa train nlu
```

## Test NLU

```shell
rasa shell nlu
```

# Generate gRPC 

```shell
python3.7 -m grpc_tools.protoc -I../../api/proto --python_out=. --pyi_out=. --grpc_python_out=. ../../api/proto/nlp.proto
```