.PHONY: build

build:
	GOOS=linux GOARCH=arm64 go build -C ./cmd/lambda -o ../../dist/lambda/bootstrap
plan:
	cd ./cmd/lambda/terraform/ && terraform plan -var-file=terraform.tfvars -out="../../../dist/lambda/deployment.tfplan"
