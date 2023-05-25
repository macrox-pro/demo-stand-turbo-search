Create a new virtual environment by choosing a Python interpreter and making a ./venv directory to hold it:

```shell
source ./venv/bin/activate
```

# Prepare

Install Rasa

```shell
pip install rasa==3.5.9
```

Install Natasha

```shell
pip install natasha
```

Install grpcio tools (without deps)

```shell
pip install --no-deps grpcio-tools==1.48.2
```

## Generate gRPC 

```shell
python -m grpc_tools.protoc -I../../api/proto --python_out=. --grpc_python_out=. ../../api/proto/nlp.proto
```


## Train NLU Rasa

```shell
rasa train nlu
```

## Test NLU Rasa

```shell
rasa shell nlu
```