serve:
	go run . --config=./config/config.yaml

redis:
	docker-compose up -d redis