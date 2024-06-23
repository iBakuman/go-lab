SRC_ROOT=${SRC_ROOT:-$(dirname $0)/../..}

cd ${SRC_ROOT}

PROTO_ROOT=api
PROTO_FILES=$(find ./${PROTO_ROOT} -not -path "./${PROTO_ROOT}/validate/*" -name *.proto)
PROTO_IMPORTS=-I=$PROTO_ROOT
PROTO_OUT=apipb

for PROTO_FILE in ${PROTO_FILES}
do 
    protoc ${PROTO_IMPORTS} \
    --go_out=./${PROTO_OUT} \
    --go_opt=paths=source_relative \
    --go-grpc_out=./${PROTO_OUT} \
    --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./${PROTO_OUT} \
    --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt=generate_unbound_methods=true \
    --openapiv2_out=./${PROTO_ROOT}/openapiv2 \
    --openapiv2_opt=generate_unbound_methods=true \
    --openapiv2_opt=openapi_naming_strategy=fqn \
    --validate_out="lang=go,paths=source_relative:./${PROTO_OUT}" \
    ${PROTO_FILE}
done