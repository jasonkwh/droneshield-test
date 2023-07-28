serve:
	go run main.go --config=./config/config.yaml serve

client:
	go run main.go --config=./config/config.yaml client

# For more information about test flags:
# https://pkg.go.dev/cmd/go/internal/test
test-integration:
	go test -count=1 -p=1 -tags=integration -v ./...

redis:
	docker-compose up -d redis