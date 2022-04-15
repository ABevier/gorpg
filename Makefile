run:
	go run cmd/main.go

test:
	go test -cover ./...

cover:
	go test -coverprofile cover.out  ./... && go tool cover -html=cover.out
