.PHONY: protoc
protoc:
	@printf $(COLOR) "Compile gateway proto files..."
	cat protoc/gen-api.sh | docker run -i --rm -v $(PWD):/root/kakuyasu-services -e SRC_ROOT=/root/kakuyasu-services protoc:0.2 sh -s
