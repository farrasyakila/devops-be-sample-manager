build-dev:
	docker build -t dev-api-go -f Dockerfile-dev .

build-staging:
	docker build -t stg-api-go -f Dockerfile-stg .

build-prod:
	docker build -t prod-api-go -f Dockerfile-prod .

.PHONY: build-dev build-staging build-prod