#!/bin/sh

protoc --proto_path=. --go_out=../back/server/. buffer.proto
protoc --proto_path=. --csharp_out=../front/Assets/scripts/. buffer.proto