#!/usr/bin/env bash
#
# This command is based on instructions from
# https://godoc.org/golang.org/x/mobile/cmd/gobind.
#

mkdir -p bindings/go_bindings
gobind -lang=go github.com/getlantern/flashlight/android/bindings > bindings/go_bindings/go_bindings.go
gobind -lang=java github.com/getlantern/flashlight/android/bindings > bindings/Flashlight.java
