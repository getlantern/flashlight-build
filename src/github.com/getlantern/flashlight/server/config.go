package server

type ServerConfig struct {
	// ListenWithOpenSSL: whether to use OpenSSL (as opposed to crypto/tls)
	// for listening.
	ListenWithOpenSSL bool

	// Country: 2 letter country code
	Country string

	// Portmap: if non-zero, server will attempt to map this port on the UPnP or
	// NAT-PMP internet gateway device
	Portmap int

	// AdvertisedHost: FQDN that is guaranteed to hit this server
	AdvertisedHost string

	// WaddellAddr: Address at which to connect to waddell for signaling
	WaddellAddr string
}
