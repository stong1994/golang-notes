//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

type Service struct {
	userRepo       Repo
	userCache      Cache
	companyAPI     CompanyAPI
	eventPublisher Publisher
}

var svcSet = wire.NewSet(
	NewUserRepo,
	NewUserCache,
	NewCompanyAPI,
	NewPublisher,
)

func NewService(
	userRepo Repo,
	userCache Cache,
	companyAPI CompanyAPI,
	eventPublisher Publisher,
) *Service {
	return &Service{
		userRepo:       userRepo,
		userCache:      userCache,
		companyAPI:     companyAPI,
		eventPublisher: eventPublisher,
	}
}

func InitService() (*Service, error) {
	panic(wire.Build(svcSet, NewService))
}
