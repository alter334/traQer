package qapi

import (
	"context"

	"github.com/traPtitech/go-traq"
)

func NewQapiHandler(auth context.Context, client *traq.APIClient) *QapiHandler {
	return &QapiHandler{auth: auth, client: client}
}
