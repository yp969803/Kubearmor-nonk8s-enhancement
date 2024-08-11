
build:
	@go build -o bin/build 


generatePolicy:
	@ ./bin/build -mode=generate

generatePolicyWithoutBlank:
	@ ./bin/build -mode=generateWithoutBlank

enforcePolicy:
	@ ./bin/build -mode=enforce 

format:
	@gofmt -w .

help:
	@echo "Available commands:"
	@echo "  make generatePolicy      - Generate the policy for the docker container in outpolicy folder"
	@echo "  make generatePolicyWithoutBlank      - Generate the policy for the docker container in outpolicy folder without blank fields"
	@echo "  make enforcePolicy       - Enforce the policy on the docker container"
	@echo "  make test                - Test the application"
	@echo "  make format               - Format the application"
	@echo "  make help                - Show this help message"