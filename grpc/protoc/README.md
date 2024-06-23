### build protoc docker image

``` shell 
docker buildx build -t protoc:0.1 .
```

buildx plugin is install by default with Docker Desktop, but if you are using lima or other docker solutions instead of Docker Desktop, you need to install buildx plugin first, please refer to [this document](https://github.com/docker/buildx?tab=readme-ov-file#linux-packages).

### run protoc

``` shell
cat gen-api.sh | docker run -i --rm -v $(PWD)/../..:/root/kakuyasu-services -e SRC_ROOT=/root/kakuyasu-services protoc:0.1 sh -s
```