.PHONY: \
	proto \
	run \

proto:
	protoc -I ./proto --go_out ./internal/controller/grpc/spec \
	--go_opt paths=source_relative \
	--go-grpc_out ./internal/controller/grpc/spec \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./internal/controller/grpc/spec \
	 --grpc-gateway_opt paths=source_relative ./proto/grpc_test_app.proto ./proto/grpc_test_app_service.proto

run:
	docker-compose up --build


