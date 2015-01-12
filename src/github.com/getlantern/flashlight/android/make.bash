#!/usr/bin/env bash
#
# This command is based on instructions from
# https://godoc.org/golang.org/x/mobile/cmd/gobind.
#
# Copyright 2014 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e

if [ ! -f make.bash ]; then
  echo 'Error running make.bash'
	exit 1
fi

mkdir -p libs/armeabi-v7a src/go/flashlight
ANDROID_APP=$PWD
(cd $GOPATH/src/golang.org/x/mobile && cp $PWD/app/*.java $ANDROID_APP/src/go)
(cd $GOPATH/src/golang.org/x/mobile && cp $PWD/bind/java/Seq.java $ANDROID_APP/src/go)
CGO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 go build -ldflags="-shared" -o libgojni.so .
mv -f libgojni.so libs/armeabi-v7a
mv ./bindings/Flashlight.java src/go/flashlight
ant debug
