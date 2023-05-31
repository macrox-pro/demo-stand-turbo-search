generate:
	go generate ./...
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import api/proto/nlp.proto

clean:
	rm internal/grpc/nlp/*.pb.go

server.run:
	go run $(CURDIR)/cmd/server

server.index.init:
	go run $(CURDIR)/cmd/server index:init

server.index.sync:
	go run $(CURDIR)/cmd/server index:sync

server.index.counts:
	go run $(CURDIR)/cmd/server index:counts > .index.counts.yaml