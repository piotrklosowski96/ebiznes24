package storage

import "go.mongodb.org/mongo-driver/bson"

func setIfNotNil[T any](bsonDocument *bson.M, name string, value *T) {
	if value != nil {
		(*bsonDocument)[name] = *value
	}
}
