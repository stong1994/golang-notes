package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"nothing/mongo"
)

type People struct {
	Name string
	Age  int
	Weight float64
}

const (
	CollectionName = "people"
)

func Insert() error {
	alice := People{
		Name:   "alice",
		Age:    18,
		Weight: 50,
	}
	err := mongo.InsertOne(CollectionName, &alice)
	if err != nil {
		return err
	}
	bob := People{
		Name:   "bob",
		Age:    25,
		Weight: 75,
	}
	cat := People{
		Name:   "cat",
		Age:    30,
		Weight: 60,
	}
	peoples := []interface{}{bob, cat}
	err = mongo.InsertMany(CollectionName, peoples)
	if err != nil {
		return err
	}
	return nil
}

func FindAll() ([]People, error) {
	var peoples []People
	err := mongo.Find(CollectionName, nil, 0, 0, nil, &peoples)
	return peoples, err
}

func AllCount() (int64, error) {
	return mongo.Count(CollectionName, nil)
}

func CountByAge(minAge, maxAge int) (int64, error) {
	filter := bson.D{
		{"age", bson.D{
			{"$gte", minAge},
			{"$lte", maxAge},
		}},
	}
	return mongo.Count(CollectionName, filter)
}

func FindOne(name string) (People, error) {
	m := bson.M{"name": name}
	var p People
	err := mongo.FindOne(CollectionName, m, &p)
	return p, err
}

func FindByAge(minAge, maxAge int, limit, skip int64) ([]People, error) {
	m := bson.D{
		bson.E{
			Key: "age",
			Value: bson.D{
				{"$gte", minAge},
				{"$lte", maxAge},
			},
		},
	}
	sort := mongo.NewEmptySortCond()
	sortCond := sort.And(mongo.CondExpr("age", true))
	var p []People
	err := mongo.Find(CollectionName, m, limit, skip, sortCond, &p)
	return p, err
}

func AddWeightByName(name string, weightChange float64) error {
	filter := bson.M{"name": name}
	update := bson.D{
		{"$inc", bson.D{
				{"weight", weightChange},},
		},
	}
	return mongo.Update(CollectionName, filter, update)
}

func UpdateWeightByName(name string, weight float64) error {
	filter := bson.M{"name": name}
	update := bson.D{
		{"$set", bson.M{"weight": weight}},
	}
	return mongo.Update(CollectionName, filter, update)
}

func UpdateWeightByAge(minAge, maxAge int, weight float64) error {
	filter := bson.D{
		{"age", bson.D{
			{"$gte", minAge},
			{"$lte", maxAge},
		}},
	}
	update := bson.D{
		{"$set", bson.M{"weight": weight}},
	}
	return mongo.Update(CollectionName, filter, update)
}

func DeleteByName(name string) error {
	filter := bson.M{"name": name}
	return mongo.Delete(CollectionName, filter)
}

type groupAgeModel struct {
	Age int `json:"age" bson:"age"`
	AvgWeight float64 `json:"avgWeight" bson:"avgWeight"`
	Count int `json:"count" bson:"count"`
}

// 统计每个年龄段的体重平均值
func GroupAge(minAge, maxAge int) ([]groupAgeModel, error) {
	matchStage := bson.D{
		{"$match", bson.D{
			{"age", bson.D{
				{"$gte", minAge},
				{"$lte", maxAge},
			}},
		}},
	}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", bson.D{
				{"age", "$age"},
			}},
			{"avgWeight", bson.D{
				{"$avg", "$weight"},
			}},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}

	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"age", "$_id.age"},
			{"avgWeight", 1},
			{"count", 1},
		}},
	}
	sort := mongo.NewEmptySortCond()
	sortCond := sort.And(mongo.CondExpr("age", true))
	var result []groupAgeModel
	err := mongo.AggregateGroup(CollectionName, matchStage, groupStage, projectStage, 0, 0, sortCond, &result)
	return result, err
}

func CountByAggregate(minAge, maxAge int) (int64, error) {
	matchStage := bson.D{
		{"$match", bson.D{
			{"age", bson.D{
				{"$gte", minAge},
				{"$lte", maxAge},
			}},
		}},
	}

	return mongo.CountWithAggregate(CollectionName, matchStage)
}

func FindWithAggregate(minAge, maxAge int64, limit, skip int) ([]People, error){
	matchStage := bson.D{
		{"$match", bson.D{
			{"age", bson.D{
				{"$gte", minAge},
				{"$lte", maxAge},
			}},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"age", 1},
			{"name", 1},
			//{"weight", 0}, 除_id外，不能显示定义不显示的字段，默认是0
		}},
	}
	var people []People
	err := mongo.FindWithAggregate(CollectionName, matchStage, projectStage, limit, skip, nil, &people)
	return people, err
}