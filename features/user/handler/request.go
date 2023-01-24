package handler

import "ecommerce/features/user"

type RegisterRequest struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address string `json:"address" form:"address"`
}

type UpdateRequest struct {
	Avatar string `json:"avatar" form:"avatar"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address string `json:"address" form:"address"`
}

type LoginRequest struct {
	Email string `json:"email"  form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Address = cnv.Address
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Avatar = cnv.Avatar
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Address = cnv.Address
	default:
		return nil
	}

	return &res
}