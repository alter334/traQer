package user

import "traQer/api"

func NewUserHandler(api *api.ApiHandler) *UserHandler {
	return &UserHandler{api: api}
}
