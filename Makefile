# For more information about test flags:
# https://pkg.go.dev/cmd/go/internal/test
test-integration:
	go test -count=1 -p=1 -tags=integration -v ./...

test-unit:
	go test -count=1 ./...

test-clean-cache:
	go clean -testcache

serve:
	go run main.go --config=./config/config.yaml serve

client:
	go run main.go --config=./config/config.yaml client

redis:
	docker-compose up -d redis

mocks:
	mockgen -package mocks -source vendor/github.com/gomodule/redigo/redis/redis.go -destination test/mocks/mock_redis_interfaces.go
	mockgen -package mocks -source internal/server/interfaces.go -destination test/mocks/mock_server_interfaces.go