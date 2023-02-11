generate:
	rm -rf ./gen
	mkdir -p ./gen
	swagger generate server \
		-t ./gen \
		-f ./spec.yml \
		-P dto.Principal \
		-m dto \
		--exclude-main \
		--strict-responders
	go mod tidy

context:
	docker compose -f deploy/docker-compose.context.yml up

compose:
	docker compose -f deploy/docker-compose.context.yml -f deploy/docker-compose.app.yml up

test:
	go test ./systest
