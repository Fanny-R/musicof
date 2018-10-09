IMAGE_NAME=fannyrottee/musicof

.PHONY: test
test:
	@go test -race -cover -timeout=5s ./...

.PHONY: package
package:
	@docker build -t ${IMAGE_NAME} .

.PHONY: publish_package
publish_package: package
	@docker push ${IMAGE_NAME}
