serve:
	go run main.go --config=./config/config.yaml serve

client:
	go run main.go --config=./config/config.yaml client

redis:
	docker-compose up -d redis