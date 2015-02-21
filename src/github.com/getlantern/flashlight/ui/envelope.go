package ui

// Values for Envelope.Type
const (
	MessageTypeProxiedSites = `ProxiedSites`
)

// Envelope is a struct that wraps messages and associates them with a type.
type Envelope struct {
	Type    string
	Message interface{}
}
