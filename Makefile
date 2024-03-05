generate-api:
	docker-compose -f docker-compose.api.yml run --rm openapi-generator-cli

build-linux-amd64:
	docker build --platform=linux/amd64 -t yinsiyu/yugod-backend:v1.1 .

launch-app:
	docker-compose -f docker-compose.app.yml up -d

image-push:
	docker push yinsiyu/yugod-backend:v1.1