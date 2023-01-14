all: generate context run

generate:
	rm -rf ./gen
	mkdir -p ./gen
	swagger generate server -t ./gen -f ./spec.yml --exclude-main --strict-responders
	go mod tidy

run:
	go run .

context:
	docker compose up -d
