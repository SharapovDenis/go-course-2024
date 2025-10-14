package app

import (
	adrepoloc "homework8/internal/repositories/ad/local"
	userrepoloc "homework8/internal/repositories/user/local"
	adsvc "homework8/internal/services/ad"
	usersvc "homework8/internal/services/user"
)

type Service struct {
	Ad   adsvc.Service
	User usersvc.Service
}

func New() Service {
	adRepoLocal := adrepoloc.New()
	userRepoLocal := userrepoloc.New()
	AdService := adsvc.New(adRepoLocal)
	UserService := usersvc.New(userRepoLocal)
	return Service{Ad: AdService, User: UserService}
}
