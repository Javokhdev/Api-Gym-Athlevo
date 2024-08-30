#!/bin/bash

CURRENT_DIR=$1
GENPROTO_DIR="${CURRENT_DIR}/genproto"

# Remove the old genproto directory
rm -rf ${GENPROTO_DIR}

# Create the genproto directory
mkdir -p ${GENPROTO_DIR}

# Generate proto files
for x in $(find ${CURRENT_DIR}/protos/gym_protos/* -type d); do
  echo "Processing directory: ${x}"
  echo "Running protoc command:"
  echo "protoc -I=${x} -I=${CURRENT_DIR}/protos/gym_protos -I /usr/local/go --go_out=${GENPROTO_DIR} --go-grpc_out=${GENPROTO_DIR} ${x}/*.proto"
  
  protoc -I=${x} -I=${CURRENT_DIR}/protos/gym_protos -I /usr/local/go \
    --go_out=${GENPROTO_DIR} --go-grpc_out=${GENPROTO_DIR} ${x}/*.proto

  echo "Contents of ${GENPROTO_DIR}:"
  ls -l ${GENPROTO_DIR}
done
