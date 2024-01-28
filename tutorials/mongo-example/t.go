package main

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// convert bson to json
func covertExample() {
	now := time.Now().UTC()
	bsonDocument := bson.D{{"hello", "world"}, {"now", now}, {"number", 12.5}}

	jsonBytes, err := bson.MarshalExtJSON(bsonDocument, true, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
	jsonBytes1, _ := bson.Marshal(bsonDocument)
	r := bson.M{}
	_ = bson.Unmarshal(jsonBytes1, &r)
	h, _ := r["hello"]
	fmt.Println(h.(string))
}

func main() {
	covertExample()
}
