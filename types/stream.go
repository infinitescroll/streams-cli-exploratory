package types

import core "github.com/textileio/go-threads/core/db"

type StreamMeta struct {
	ID            core.InstanceID `json:"_id"`
	Name          string          `json:"name"`
	FFSID         string          `json:"ffsID"`
	FFSAuthToken  string          `json:"FFSAuthToken"`
	WalletAddress string          `json:"walletAddress"`
}

type StreamPointer struct {
	ID       core.InstanceID `json:"_id"`
	ThreadID string          `json:"threadID"`
	Name     string          `json:"name"`
}
