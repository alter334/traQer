package qapi

import (
	"context"

	"github.com/traPtitech/go-traq"
)

type Qapi struct {
	auth   context.Context
	client *traq.APIClient
}
