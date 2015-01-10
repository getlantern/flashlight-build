// Not supported in android.
// +build android

package server

func (server *Server) startNattywad(waddellAddr string) {
	log.Debugf("startNattywad is not implemented.")
}

func (server *Server) stopNattywad() {
	log.Debugf("stopNattywad is not implemented.")
}
