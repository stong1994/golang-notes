package main

import (
	"fmt"
	"nothing/mongo"
	"testing"
)

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertAndGet(t *testing.T) {
	mongo.Init()
	handleErr(t, mongo.DropCollection(CollectionName))

	handleErr(t, Insert())
	count, err := AllCount()
	handleErr(t, err)
	if count != 3 {
		t.Fatalf("insert %d, but got %d", 3, count)
	}
	all, err := FindAll()
	handleErr(t, err)
	for _, v := range all {
		fmt.Printf("%v\n", v)
	}
	c1, err := CountByAge(20, 40)
	handleErr(t, err)
	c2, err := CountByAggregate(20, 40)
	handleErr(t, err)

	t.Log("FindByAge desc")
	p1, err := FindByAge(20, 40, 0, 0)
	handleErr(t, err)
	p4, err := FindByAge(20, 40, 0, 1)
	handleErr(t, err)
	p5, err := FindByAge(20, 40, 1, 0)
	handleErr(t, err)
	if len(p1) != 2 || c1 != 2 || c2 != 2 || len(p4) != 1 || len(p5) != 1{
		t.Fatal("find failed")
	}
	for _, v := range p1 {
		t.Logf("\t%+v", v)
		if v.Age > 40 || v.Age < 20 {
			t.Fatal("filter not valid")
		}
	}
	p3, err := FindWithAggregate(20,40, 0, 0)
	handleErr(t, err)
	if len(p1) != len(p3) {
		t.Fatal("FindWithAggregate failed")
	}

	p2, err := FindOne("alice")
	handleErr(t, err)
	if p2.Name != "alice" || p2.Age != 18 || p2.Weight != 50 {
		t.Fatalf("find one failed: %v", p2)
	}

	err = AddWeightByName("alice", 1)
	handleErr(t, err)
	p2, err = FindOne("alice")
	handleErr(t, err)
	if p2.Weight != 51 {
		t.Fatalf("add weight failed: want %f got %f", 51.0, p2.Weight)
	}
	t.Log("AddWeightByName success")

	err = UpdateWeightByName("alice", 60)
	handleErr(t, err)
	p2, err = FindOne("alice")
	if p2.Weight != 60 {
		t.Fatalf("add weight failed: want %f got %f", 60.0, p2.Weight)
	}
	t.Log("UpdateWeightByName success")

	err = UpdateWeightByAge(20, 40, 60)
	handleErr(t, err)
	p1, err = FindByAge(20, 40, 0, 0)
	handleErr(t, err)
	for _, v := range p1 {
		if v.Weight != 60 {
			t.Fatal("UpdateWeightByAge failed")
		}
	}
	t.Log("UpdateWeightByAge success")

	handleErr(t, mongo.InsertOne(CollectionName, People{
		Name:   "bigBob",
		Age:    25,
		Weight: 70,
	}))

	t.Log("GroupAge desc")
	group, err := GroupAge(20, 40)
	handleErr(t, err)
	for _, v := range group {
		fmt.Printf("%+v\n", v)
	}
	t.Log("GroupAge success")

	handleErr(t, DeleteByName("alice"))
	p2, err = FindOne("alice")
	if p2.Name != "" {
		t.Fatal("deleteByName failed")
	}
	t.Log("DeleteByName success")
}
