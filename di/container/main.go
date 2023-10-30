package main

import (
	"context"
	"golang-learning/di/container/di"
)

const (
	keyRepo       = "repo"
	keyCache      = "cache"
	keyCompanyAPI = "companyAPI"
	keyPublisher  = "publisher"
)

type Repo interface{}

func NewUserRepo() Repo {
	return 1
}

type Cache interface{}

func NewUserCache() Cache {
	return 2
}

type CompanyAPI interface{}

func NewCompanyAPI() CompanyAPI {
	return 3
}

type Publisher interface{}

func NewPublisher() Publisher {
	return 4
}

type Service struct {
	userRepo       Repo
	userCache      Cache
	companyAPI     CompanyAPI
	eventPublisher Publisher
}

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

func main() {
	container := di.New()
	container.AddSingleton(keyRepo, func(c di.Container) (any, error) {
		return NewUserRepo(), nil
	})
	container.AddSingleton(keyCache, func(c di.Container) (any, error) {
		return NewUserCache(), nil
	})
	container.AddSingleton(keyCompanyAPI, func(c di.Container) (any, error) {
		return NewCompanyAPI(), nil
	})
	container.AddSingleton(keyPublisher, func(c di.Container) (any, error) {
		return NewPublisher(), nil
	})

	service := newService(container)
	_ = service

	ctx := context.Background()
	ctx = container.Scoped(ctx)

}

func newService(container di.Container) *Service {
	service := Service{
		userRepo:       container.Get(keyRepo).(Repo),
		userCache:      container.Get(keyCache).(Cache),
		companyAPI:     container.Get(keyCompanyAPI).(CompanyAPI),
		eventPublisher: container.Get(keyPublisher).(Publisher),
	}
	return &service
}
