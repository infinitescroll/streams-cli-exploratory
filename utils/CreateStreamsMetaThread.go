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
	err = client.NewDB(context.Background(), id)
	CheckErr(err)

	// move to setup func
	err = client.NewCollection(context.Background(), id, db.CollectionConfig{
		Name:   "Organization",
		Schema: util.SchemaFromInstance(&types.Owner{}, false),
	})
	CheckErr(err)

	err = client.NewCollection(context.Background(), id, db.CollectionConfig{
		Name:   "Individual",
		Schema: util.SchemaFromInstance(&types.Owner{}, false),
	})
	CheckErr(err)
}
