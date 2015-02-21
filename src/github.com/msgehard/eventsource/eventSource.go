/*
Package eventsouce provides a basic interface for
serving server sent events (aka EventSource).

It is patterned after the Go Websocket library:
https://code.google.com/p/go/source/browse?repo=net#hg%2Fwebsocket

For more destails about Server Sent Events see:
http://html5doctor.com/server-sent-events/
http://www.html5rocks.com/en/tutorials/eventsource/basics/
http://cjihrig.com/blog/the-server-side-of-server-sent-events/
*/

package eventsource

import (
	"fmt"
	"net/http"
)

type Conn struct {
	writer http.ResponseWriter
	http.Flusher
	http.CloseNotifier
}

func (c Conn) Write(message string) {
	c.writer.Write([]byte(fmt.Sprintf("data: %s", message)))
	c.writer.Write([]byte("\n\n"))
	c.Flush()
}

type Handler func(*Conn)

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f, f_ok := w.(http.Flusher)

	if !f_ok {
		panic("ResponseWriter is not a Flusher")
	}

	cn, cn_ok := w.(http.CloseNotifier)

	if !cn_ok {
		panic("ResponseWriter is not a CloseNotifier")
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")

	f.Flush()

	h(&Conn{w, f, cn})
}
