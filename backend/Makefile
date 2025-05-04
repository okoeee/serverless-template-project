.PHONY: build clean deploy deploy-dev deploy-stg local-api test

PROJECT_NAME = template-project

ENV ?= dev

ENV_DEV = dev
ENV_STG = stg
ENV_PROD = prod

build:
	sam build --use-container --cached

clean:
	rm -rf .aws-sam

deploy:
	sam deploy \
		--stack-name $(STACK_NAME) \
		--parameter-overrides Environment=$(ENV) \
		--capabilities CAPABILITY_IAM

deploy-dev:
	$(MAKE) deploy ENV=$(ENV_DEV) STACK_NAME=$(ENV_DEV)-$(PROJECT_NAME)

deploy-stg:
	$(MAKE) deploy ENV=$(ENV_STG) STACK_NAME=$(ENV_STG)-$(PROJECT_NAME)

local-api:
	make build
	sam local start-api \
		--env-vars local/dev.json \
		--docker-network sam-template-project 

local-dynamodb:
	docker run -d \
		--name dynamodb-sam-template-project \
		--network sam-template-project \
		-p 8000:8000 \
		amazon/dynamodb-local:latest

init-local-dynamodb:
	aws dynamodb create-table \
		--region ap-northeast-1 \
		--table-name dev-users \
		--attribute-definitions AttributeName=userId,AttributeType=S \
		--key-schema AttributeName=userId,KeyType=HASH \
		--billing-mode PAY_PER_REQUEST \
		--endpoint-url http://localhost:8000

delete-local-dynamodb:
	aws dynamodb delete-table \
		--region ap-northeast-1 \
		--table-name dev-users \
		--endpoint-url http://localhost:8000

test:
	go test ./...
