init:
	docker compose exec server go mod init dice
start:
	docker compose exec server go mod tidy
	docker compose exec server go run .

proto:
	docker compose exec server protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative dice.proto
