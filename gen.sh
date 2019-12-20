#!/usr/bin/env bash
protoc -I todo/ todo/todo.proto \
--go_out=plugins=grpc:todo \
--js_out=import_style=commonjs,binary:todo-client/src/ \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:todo-client/src/
