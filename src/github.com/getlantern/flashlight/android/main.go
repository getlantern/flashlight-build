package main

import (
	_ "github.com/getlantern/flashlight/android/bindings/go_bindings"
	_ "golang.org/x/mobile/app"
	_ "golang.org/x/mobile/bind/java"
)

func main() {
	app.Run(app.Callbacks{})
}
