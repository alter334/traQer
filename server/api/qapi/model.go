package qapi

import (
	"context"

	"github.com/traPtitech/go-traq"
)

type QapiHandler struct {
	auth   context.Context
	client *traq.APIClient
}
