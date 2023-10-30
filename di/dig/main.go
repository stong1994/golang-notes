package main

import (
	"fmt"
	"go.uber.org/dig"
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

func (svc *Service) DoSomething() {
	if svc.userRepo == nil ||
		svc.userCache == nil ||
		svc.companyAPI == nil ||
		svc.eventPublisher == nil {
		panic("init failed")
	}
	fmt.Println("init success")
}

func main() {
	container := dig.New()
	mustProvide(container, NewUserRepo)
	mustProvide(container, NewUserCache)
	mustProvide(container, NewCompanyAPI)
	mustProvide(container, NewPublisher)
	mustProvide(container, NewService)

	err := container.Invoke(func(svc *Service) {
		svc.DoSomething()
	})
	mustNoErr(err)
}

func mustProvide(container *dig.Container, constructor interface{}) {
	err := container.Provide(constructor)
	mustNoErr(err)
}

func mustNoErr(err error) {
	if err != nil {
		panic(err)
	}
}
