generate:
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import api/proto/nlp.proto

clean:
	rm internal/grpc/nlp/*.pb.go