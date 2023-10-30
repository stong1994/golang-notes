package main

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
