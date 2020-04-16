package responsibility_chain

import (
	"fmt"
	"testing"
)

func Login(request *Request){
	username := request.PostForm["username"][0]
	password := request.PostForm["password"][0]

	fmt.Println(username)
	fmt.Println(password)
}

func TestResponsibilityChain(t *testing.T) {
	mux := NewMux()

	mux.Handle("login",Login)


	req := NewRequest()
	req.Method="POST"
	req.SetValues("username","111")
	req.SetValues("password","222")
	req.Post("login",mux)

}