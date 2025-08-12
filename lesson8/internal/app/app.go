package app

import (
	"homework8/internal/repositories/adrepo"
	"homework8/internal/repositories/userrepo"
	"homework8/internal/service"
	"homework8/internal/service/adsvc"
	"homework8/internal/service/usersvc"
)

func New() service.Service {
	adRepo := adrepo.New()
	userRepo := userrepo.New()
	AdService := adsvc.New(adRepo)
	UserService := usersvc.New(userRepo)
	return service.New(AdService, UserService)
}
