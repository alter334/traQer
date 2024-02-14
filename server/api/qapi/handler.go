package qapi

import (
	"context"
	"os"

	"github.com/traPtitech/go-traq"
)

type QapiHandler struct {
	auth   context.Context
	client *traq.APIClient
}

func NewQapiHandler(auth context.Context, client *traq.APIClient) *QapiHandler {
	return &QapiHandler{auth: auth, client: client}
}

func QapiSetup() *QapiHandler {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, os.Getenv("TRAQ_TOKEN"))

	qapi := NewQapiHandler(auth, client)

	return qapi
}
