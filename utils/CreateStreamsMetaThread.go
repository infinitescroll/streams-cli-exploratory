package utils

import (
	"context"
	"fmt"

	"github.com/openworklabs/streams-cli/v2/types"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/go-threads/util"
)

func CreateStreamsMetaThread(client *client.Client) {
	// TODO: READ THIS IN FROM CONFIG - viper
	id, err := thread.Decode("bafk2pukjgfvfgantvjqk7ggtv7h2brji2iw74ztfcfsq5so6kx6alkq")
	CheckErr(err)
	fmt.Println(id)
	// client.DeleteDB(context.Background(), id)
	client.NewDB(context.Background(), id)

	// move to setup func
	client.NewCollection(context.Background(), id, db.CollectionConfig{
		Name:   "Organization",
		Schema: util.SchemaFromInstance(&types.OwnerPointer{}, false),
	})

	client.NewCollection(context.Background(), id, db.CollectionConfig{
		Name:   "Individual",
		Schema: util.SchemaFromInstance(&types.OwnerPointer{}, false),
	})
}
