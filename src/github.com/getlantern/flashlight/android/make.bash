#!/usr/bin/env bash
# Copyright 2014 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e

if [ ! -f make.bash ]; then
  echo 'Error running make.bash'
	exit 1
fi

mkdir -p bindings/go_bindings
gobind -lang=go github.com/getlantern/flashlight/android/bindings > bindings/go_bindings/go_bindings.go
gobind -lang=java github.com/getlantern/flashlight/android/bindings > bindings/Bindings.java

mkdir -p libs/armeabi-v7a src/go/flashlight
ANDROID_APP=$PWD
(cd $GOPATH/src/golang.org/x/mobile && ln -sf $PWD/app/*.java $ANDROID_APP/src/go)
(cd $GOPATH/src/golang.org/x/mobile && ln -sf $PWD/bind/java/Seq.java $ANDROID_APP/src/go)
CGO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 \
	go build -ldflags="-shared" .
mv -f bindings libs/armeabi-v7a/libgojni.so
ant debug
