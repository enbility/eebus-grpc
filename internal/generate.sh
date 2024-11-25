#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright Pionix GmbH and Contributors to EVerest


PROTO_PATH="../protobuf"
OUT="../rpc_services"
mkdir -p "$OUT"

# Common Types
protoc \
    --proto_path="$PROTO_PATH" \
    --go_out="$OUT" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$OUT" \
    --go-grpc_opt=paths=source_relative \
    "$PROTO_PATH/common_types/types.proto"

# Control Service
protoc \
    --proto_path="$PROTO_PATH" \
    --go_out="$OUT" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$OUT" \
    --go-grpc_opt=paths=source_relative \
    "$PROTO_PATH/control_service/control_service.proto" \
    "$PROTO_PATH/control_service/messages.proto" \
    "$PROTO_PATH/control_service/types.proto"


# EG LPC Service
protoc \
    --proto_path="$PROTO_PATH" \
    --go_out="$OUT" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$OUT" \
    --go-grpc_opt=paths=source_relative \
    "$PROTO_PATH/usecases/eg/lpc/service.proto" \
    "$PROTO_PATH/usecases/eg/lpc/messages.proto"

# CS LPC Service
protoc \
    --proto_path="$PROTO_PATH" \
    --go_out="$OUT" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$OUT" \
    --go-grpc_opt=paths=source_relative \
    "$PROTO_PATH/usecases/cs/lpc/service.proto" \
    "$PROTO_PATH/usecases/cs/lpc/messages.proto"
