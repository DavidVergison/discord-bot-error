package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
)

type DiscordCommandValidator struct {
	publicKey string
}

func (d DiscordCommandValidator) verifyRequest(signature string, timestamp string, body string) error {
	// Decode the public key
	pubKey, err := hex.DecodeString(d.publicKey)
	if err != nil {
		return fmt.Errorf("failed to decode public key: %v", err)
	}

	// Decode the public key
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("failed to decode sig: %v", err)
	}

	// Verify the signature
	if !ed25519.Verify(pubKey, []byte(timestamp+body), sig) {
		return fmt.Errorf("invalid request signature")
	}

	return nil
}

func (d DiscordCommandValidator) IsPing(cmdType int) bool {
	return cmdType == 1
}
