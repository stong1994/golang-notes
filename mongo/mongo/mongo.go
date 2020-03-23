package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	MaxAggregateTime = 15 * time.Second
)

var mongodb *mongo.Database

func Init() {
	opts := &options.ClientOptions{}
	//connString := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", user, password, host, port, dbname)
	connString := "mongodb://127.0.0.1:27017"
	opts.ApplyURI(connString)
	opts.SetMaxPoolSize(30) //设置连接池大小
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	//Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	// Set MongoDB database
	mongodb = client.Database("test")
	if mongodb == nil {
		panic("db is not init")
	}
}

func GetDB() *mongo.Database {
	return mongodb
}

func FindOne(collection string, filter bson.M, result interface{}) error {
	c := GetDB().Collection(collection)
	err := c.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

// limit skip无约束 传入0或者小于0即可 descFields 1为增序 -1 为逆序
func Find(collection string, filter bson.D, limit int64, skip int64, sort SortCond, result interface{}) error {
	ctx := context.Background()
	c := GetDB().Collection(collection)
	option := options.Find()
	if limit > 0 {
		option.SetLimit(limit)
	}
	if skip > 0 {
		option.SetSkip(skip)
	}
	if sort != nil && sort.IsValid() {
		option.SetSort(convertSort(sort))
	}
	if filter == nil {
		filter = bson.D{}
	}
	cur, err := c.Find(ctx, filter, option)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	return cur.All(ctx, result)
}

func InsertOne(collection string, obj interface{}) error {
	c := GetDB().Collection(collection)
	insertResult, err := c.InsertOne(context.Background(), obj)
	if err != nil {
		return err
	}
	fmt.Println("id of the inserted data is", insertResult.InsertedID)
	return nil
}

func InsertMany(collection string, obj []interface{}) error {
	c := GetDB().Collection(collection)
	result, err := c.InsertMany(context.TODO(), obj)
	if err != nil {
		return err
	}
	fmt.Println("id of the inserted data is", result.InsertedIDs)
	return nil
}

// filter should be bson.D or bson.M
func Delete(collection string, filter interface{}) error {
	c := GetDB().Collection(collection)
	result, err := c.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Println("deleted count is", result.DeletedCount)
	return nil
}

// update should be bson.D or bson.M
func Update(collection string, filter interface{}, update interface{}) error {
	c := GetDB().Collection(collection)
	result, err := c.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("matched count is %d and modified count is %d\n", result.MatchedCount, result.ModifiedCount)
	return nil
}

func Count(collection string, filter interface{}) (int64, error) {
	c := GetDB().Collection(collection)
	if filter == nil {
		filter = bson.D{}
	}
	return c.CountDocuments(context.Background(), filter)
}

// ==========================================  Aggregation start ==========================================
func FindWithAggregate(collection string, matchStage, projectStage bson.D, limit, skip int, sort SortCond, result interface{}) error {
	c := GetDB().Collection(collection)
	ctx := context.Background()
	pip := mongo.Pipeline{matchStage, projectStage}
	opts := options.Aggregate().SetMaxTime(MaxAggregateTime)
	if sort != nil {
		pip = append(pip, bson.D{
			{"$sort", convertSort(sort)},
		})
	}
	if skip > 0 {
		pip = append(pip, bson.D{{"$skip", skip}})
	}
	if limit > 0 {
		pip = append(pip, bson.D{{"$limit", limit}})
	}
	cursor, err := c.Aggregate(ctx, pip, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func AggregateGroup(collection string, matchStage, groupStage, projectStage bson.D, limit, skip int, sort SortCond, result interface{}) error {
	c := GetDB().Collection(collection)
	ctx := context.Background()
	pip := mongo.Pipeline{matchStage, groupStage, projectStage}
	opts := options.Aggregate().SetMaxTime(MaxAggregateTime)
	if sort.IsValid() {
		pip = append(pip, bson.D{
			{"$sort", convertSort(sort)},
		})
	}
	if skip > 0 {
		pip = append(pip, bson.D{{"$skip", skip}})
	}
	if limit > 0 {
		pip = append(pip, bson.D{{"$limit", limit}})
	}
	cursor, err := c.Aggregate(ctx, pip, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func CountWithAggregate(collection string, matchStage bson.D) (int64, error) {
	c := GetDB().Collection(collection)
	ctx := context.Background()
	pip := mongo.Pipeline{matchStage}
	opts := options.Aggregate().SetMaxTime(MaxAggregateTime)
	count := bson.D{
		{"$count", "count"},
	}
	project := bson.D{
		{"$project", bson.D{
			{"count", 1},
		}},
	}
	pip = append(pip, count, project)
	cursor, err := c.Aggregate(ctx, pip, opts)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)
	type model struct {
		Count int64 `bson:"count"`
	}
	var m model
	for cursor.Next(ctx) {
		err = cursor.Decode(&m)
		break
	}
	return m.Count, err
}
// ==========================================  Aggregation end ==========================================
// delete collection
func DropCollection(collection string) error {
	coll := GetDB().Collection(collection)
	return coll.Drop(context.Background())
}

func convertSort(cond SortCond) bson.D {
	sort := bson.D{}
	if !cond.IsValid() {
		return sort
	}
	switch cond.(type) {
	case sortConditions:
		for _, cond := range cond.(sortConditions) {
			value := -1
			if !cond.(SingleSortCondition).Desc {
				value = 1
			}
			sort = append(sort, bson.E{cond.(SingleSortCondition).Key, value})
		}
		return sort
	case SingleSortCondition:
		value := -1
		if !cond.(SingleSortCondition).Desc {
			value = 1
		}
		return bson.D{{cond.(SingleSortCondition).Key, value}}
	default:
		panic("not valid type of sortCond")
	}
}