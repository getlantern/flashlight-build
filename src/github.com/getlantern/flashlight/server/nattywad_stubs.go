// This file is used to disable nattywad on android hosts.
//
// +build android

package server

// startNattywad is not implemented on android.
func (server *Server) startNattywad(waddellAddr string) {
	log.Debugf("startNattywad is not implemented.")
}

// stopNattywad is not implemented on android.
func (server *Server) stopNattywad() {
	log.Debugf("stopNattywad is not implemented.")
}
