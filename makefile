all: generate context run

generate:
	rm -rf ./gen
	mkdir -p ./gen
	swagger generate server -t ./gen -f ./spec.yml -P models.Principal --exclude-main --strict-responders
	go mod tidy

run:
	go run .

context:
	docker compose up -d

cloc:
	cloc --exclude-dir=gen .

test:
	(cd systest && godog run .)

docker-image:
	docker build -t settingstore -f deploy/dockerfile .
