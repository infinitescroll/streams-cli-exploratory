package types

import (
	core "github.com/textileio/go-threads/core/db"
)

type Owner struct {
	ID       core.InstanceID `json:"_id"`
	ThreadID string          `json:"threadID"`
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
