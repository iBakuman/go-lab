SRC_ROOT=${SRC_ROOT:-$(dirname $0)/../..}

cd ${SRC_ROOT}

PROTO_ROOT=internal/api
PROTO_FILES=$(find ./${PROTO_ROOT} -name *.proto)
PROTO_IMPORTS="-I=$PROTO_ROOT -I=api"
PROTO_OUT=internal/apipb

for PROTO_FILE in ${PROTO_FILES}
do 
    protoc ${PROTO_IMPORTS} \
    --go_out=./${PROTO_OUT} \
    --go_opt=paths=source_relative \
    --go-grpc_out=./${PROTO_OUT} \
    --go-grpc_opt=paths=source_relative \
    --validate_out="lang=go,paths=source_relative:./${PROTO_OUT}" \
    ${PROTO_FILE}
done
