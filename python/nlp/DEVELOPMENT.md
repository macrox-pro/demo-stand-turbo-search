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

# Generate gRPC 

```shell
python3.7 -m grpc_tools.protoc -I../../api/proto --python_out=. --pyi_out=. --grpc_python_out=. ../../api/proto/nlp.proto
```