#!/bin/sh

if [ "$#" -ne 1 ]; then
    echo "lib_grpc: only 1 argument is required."
    echo "usage: ./lib_grpc unity_game"
else
    FILE=$1

    if mkdir -p $FILE/Contents/Frameworks/MonoEmbedRuntime/osx; then
        if cp bin_tools/tools/macosx_x64/libgrpc_csharp_ext.bundle $FILE/Contents/Frameworks/MonoEmbedRuntime/osx/ ; then
            echo "Lib successfully imported"
        fi
    fi
fi