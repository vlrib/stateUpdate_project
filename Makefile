postgres:
	docker rm --force projectdb
	docker run -d --name projectdb -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=your-password -e POSTGRES_DB=teste_db postgres:11-alpine

proto:
	sudo protoc -I stateUpdate/ stateUpdate/stateUpdate.proto --go_out=plugins=grpc:stateUpdate
