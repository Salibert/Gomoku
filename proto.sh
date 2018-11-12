#!/bin/sh

PATH_GRPC_FRONT=$PWD/front/Packages/Grpc.Tools.1.16.0/tools/macosx_x64
PLUGIN_FRONT=protoc-gen-grpc=$PATH_GRPC_FRONT/grpc_csharp_plugin
PATH_OUT_FRONT=front/Assets/scripts

rm -rf back/server/pb/buffer.pb.go
protoc --proto_path=proto --go_out=plugins=grpc:back/server/pb/. buffer.proto

rm -rf $PATH_OUT_FRONT/Buffer.cs
rm -rf $PATH_OUT_FRONT/BufferGRPC.cs
$PATH_GRPC_FRONT/protoc --proto_path=proto --plugin=$PLUGIN_FRONT/. --csharp_out $PATH_OUT_FRONT --plugin=$PLUGIN_FRONT/. --grpc_out $PATH_OUT_FRONT buffer.proto 