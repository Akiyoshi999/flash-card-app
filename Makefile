include .env
deploy:
	cd infra && cdk deploy --profile $(AWS_PROFILE)

diff:
	cd infra && cdk diff --profile $(AWS_PROFILE)

destroy:
	cd infra && cdk destroy --profile $(AWS_PROFILE)

fe-build:
	cd frontend && npm run build

