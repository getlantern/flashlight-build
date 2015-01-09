#!/bin/bash
mkdir -p bindings/go_bindings
gobind -lang=go github.com/getlantern/flashlight/android/bindings > bindings/go_bindings/go_bindings.go
gobind -lang=java github.com/getlantern/flashlight/android/bindings > bindings/Bindings.java
