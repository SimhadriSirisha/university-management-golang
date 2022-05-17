bind_migrations:
	go-bindata -pkg migrations -o db/migrations.go -prefix='db/migrations/' db/migrations/

Proto:
	protoc --go-grpc_out=protoclient --go_out=protoclient university-management.proto

docker:
	docker-compose up -d

open_db:
	 psql  -p 5436 -U postgres -h localhost

server:
	go run server/main/main.go

client:
	go run client/main/main.go