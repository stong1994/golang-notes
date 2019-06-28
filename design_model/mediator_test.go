package design_model

import "testing"

func TestMediator(t *testing.T) {
	m := &Mediator{}
	u1 := &User1{m, "andy"}
	u2 := &User2{m, "chris"}
	m.u1 =  u1
	m.u2 = u2
	u1.SendMsg("How Are You")
	u2.SendMsg("I'm Fine!")

}
