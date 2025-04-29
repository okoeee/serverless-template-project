.PHONY: build clean deploy deploy-dev deploy-stg local-api test

PROJECT_NAME = template-project

ENV ?= dev

ENV_DEV = dev
ENV_STG = stg
ENV_PROD = prod

build:
	sam build

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
	sam local start-api --env-vars local/dev.json

test:
	go test ./...
