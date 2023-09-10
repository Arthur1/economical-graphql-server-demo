db.sqlite:
	$(RM) db.sqllite
	go run ./cmd/gen-db-file

bootstrap:
	CGO_ENABLED=0 GO_OS=linux GOARCH=arm64 go build -o ./bootstrap ./cmd/server-lambda

server-lambda.zip: db.sqlite bootstrap
	zip server-lambda.zip bootstrap db.sqlite

clean:
	$(RM) db.sqllite bootstrap server-lambda.zip
