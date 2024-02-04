package qapi

import (
	"context"

	"github.com/traPtitech/go-traq"
)

func NewQapiHandler(auth context.Context, client *traq.APIClient) *Qapi {
	return &Qapi{auth: auth, client: client}
}
