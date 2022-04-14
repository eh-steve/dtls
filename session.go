package dtls

import "time"

// Session store data needed in resumption
type Session struct {
	// ID store session id
	ID []byte
	// Secret store session master secret
	Secret []byte
	// Optional expiry based on a verified certificate NotAfter date, otherwise empty
	Expiry time.Time
}

// SessionStore defines methods needed for session resumption.
type SessionStore interface {
	// Set save a session.
	// For client, use server name as key.
	// For server, use session id.
	Set(key []byte, s Session) error
	// Get fetch a session.
	Get(key []byte) (Session, error)
	// Del clean saved session.
	Del(key []byte) error
}
