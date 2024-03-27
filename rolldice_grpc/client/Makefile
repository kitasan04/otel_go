init:
	docker compose exec client go mod init dice
install:
	docker compose exec client go install .
start:
	docker compose exec client go mod tidy
	docker compose exec client go run .

proto:
	docker compose exec client protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative dice.proto
