package user

import "traQer/api"

// Handler
type UserHandler struct {
	api *api.ApiHandler
}

func NewUserHandler(api *api.ApiHandler) *UserHandler {
	return &UserHandler{api: api}
}
