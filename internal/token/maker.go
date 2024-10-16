package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// Create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}