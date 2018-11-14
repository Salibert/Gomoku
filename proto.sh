#!/bin/sh

PLUGIN_FRONT=protoc-gen-grpc=bin_tools/tools/macosx_x64/grpc_csharp_plugin
PROTO=bin_tools/protoc/macosx_x64/protoc
PATH_OUT_FRONT=front/Assets/scripts

rm -rf back/server/pb/buffer.pb.go
$PROTO --proto_path=proto --go_out=plugins=grpc:back/server/pb/. buffer.proto

rm -rf $PATH_OUT_FRONT/Buffer.cs
rm -rf $PATH_OUT_FRONT/BufferGRPC.cs
$PROTO --proto_path=proto --plugin=$PLUGIN_FRONT --csharp_out $PATH_OUT_FRONT/. --plugin=$PLUGIN_FRONT --grpc_out $PATH_OUT_FRONT buffer.proto