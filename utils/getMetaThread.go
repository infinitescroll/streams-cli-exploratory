package utils

import "github.com/textileio/go-threads/core/thread"

func GetMetaThread() thread.ID {
	id, err := thread.Decode("bafk2pukjgfvfgantvjqk7ggtv7h2brji2iw74ztfcfsq5so6kx6alkq")
	CheckErr(err)

	return id
}
