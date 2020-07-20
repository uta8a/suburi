#!/bin/sh

CLIENT_OUTDIR=client/src/proto/user
SERVER_OUTPUT_DIR=server/proto/user

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=internal/proto/user user.proto \
    --js_out=import_style=commonjs:${CLIENT_OUTDIR} \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTDIR} \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
sed -i '1i/* eslint-disable */' ${CLIENT_OUTDIR}/user_pb.js

CLIENT_OUTDIR=client/src/proto/challenge
SERVER_OUTPUT_DIR=server/proto/challenge

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=internal/proto/challenge challenge.proto \
    --js_out=import_style=commonjs:${CLIENT_OUTDIR} \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTDIR} \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
sed -i '1i/* eslint-disable */' ${CLIENT_OUTDIR}/challenge_pb.js

CLIENT_OUTDIR=client/src/proto/check
SERVER_OUTPUT_DIR=server/proto/check

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=internal/proto/check check.proto \
    --js_out=import_style=commonjs:${CLIENT_OUTDIR} \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTDIR} \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
sed -i '1i/* eslint-disable */' ${CLIENT_OUTDIR}/check_pb.js

CLIENT_OUTDIR=client/src/proto/id
SERVER_OUTPUT_DIR=server/proto/id

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=internal/proto/id id.proto \
    --js_out=import_style=commonjs:${CLIENT_OUTDIR} \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTDIR} \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
sed -i '1i/* eslint-disable */' ${CLIENT_OUTDIR}/id_pb.js

CLIENT_OUTDIR=client/src/proto/scoreboard
SERVER_OUTPUT_DIR=server/proto/scoreboard

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=internal/proto/scoreboard scoreboard.proto \
    --js_out=import_style=commonjs:${CLIENT_OUTDIR} \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTDIR} \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
sed -i '1i/* eslint-disable */' ${CLIENT_OUTDIR}/scoreboard_pb.js
