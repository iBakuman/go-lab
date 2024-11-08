export PWD = $(shell pwd)

COLOR := "\e[1;36m%s\e[0m\n"

.PHONY: build-image
build-image:
	@docker image inspect protoc:1.0 > /dev/null 2>&1 && \
	printf $(COLOR) "Image already exists..." || \
	(echo "Image does not exist, building..." && \
	docker buildx build -t protoc:1.0 -f .docker/Dockerfile . && \
	printf $(COLOR) "Image built successfully...")

.PHONY: compile
compile: build-image
	@printf $(COLOR) "Compile proto files..."
	@cat .docker/gen.sh | docker run -i --rm -v $(PWD):/root/app -e SRC_ROOT=/root/app protoc:1.0 sh -s

.PHONY: compile-with-buf
compile-with-buf:
	@buf dep update
	@buf generate