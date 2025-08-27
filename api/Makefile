PROTOC = protoc
PROTO_DIR = proto
OUT_DIR = pb

all: generate

generate:
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto
