api-dev:
	cd cmd/api && go run main.go

api-prod:
	cd cmd/api && go build -o teste-build