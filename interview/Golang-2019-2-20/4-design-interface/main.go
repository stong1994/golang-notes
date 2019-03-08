package main

type Amimal interface {
	Name()
	Sex()
}

type Person struct {
}

func (p *Person) Name() {

}

func (p *Person) Sex() {

}
