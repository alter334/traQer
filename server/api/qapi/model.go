package qapi

import (
	"context"

	"github.com/traPtitech/go-traq"
)

type QapiHandler struct {
	auth   context.Context
	client *traq.APIClient
}

type UserMessages struct {
	User              traq.User `json:"user"`
	TotalMessageCount int64     `json:"TotalMessageCount"`
}
