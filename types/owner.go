package types

import (
	core "github.com/textileio/go-threads/core/db"
)

type OwnerPointer struct {
	ID       core.InstanceID `json:"_id"`
	ThreadID string          `json:"threadID"`
	Name     string          `json:"name"`
}

type OwnerMetadata struct {
	ID    core.InstanceID `json:"_id"`
	Name  string          `json:"name"`
	Email string          `json:"email"`
}

type OwnerToken struct {
	// resource like GitHub
	Resource     string `json:"resource"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
