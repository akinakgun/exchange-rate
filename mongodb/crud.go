package mongodb

import (
	"context"
	"exchange-rate/domain/exchangerate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(task exchangerate.ExchangeRate) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(db).Collection(collection)

	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}

	return nil
}

func CreateMany(list []exchangerate.ExchangeRate) error {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}

	client, err := GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(db).Collection(collection)

	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}

	return nil
}

func GetById(currencyCode string) (exchangerate.ExchangeRate, error) {
	rate := exchangerate.ExchangeRate{}
	filter := bson.D{primitive.E{Key: "CurrencyCode", Value: currencyCode}}

	client, err := GetMongoClient()
	if err != nil {
		return rate, err
	}

	collection := client.Database(db).Collection(collection)
	err = collection.FindOne(context.TODO(), filter).Decode(&rate)
	if err != nil {
		return rate, err
	}

	return rate, nil
}

func GetAll() ([]exchangerate.ExchangeRate, error) {
	filter := bson.D{{}}
	var rates []exchangerate.ExchangeRate

	client, err := GetMongoClient()
	if err != nil {
		return rates, err
	}

	collection := client.Database(db).Collection(collection)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return rates, findError
	}

	for cur.Next(context.TODO()) {
		var t exchangerate.ExchangeRate
		err := cur.Decode(&t)
		if err != nil {
			return rates, err
		}
		rates = append(rates, t)
	}

	cur.Close(context.TODO())
	if len(rates) == 0 {
		return rates, mongo.ErrNoDocuments
	}
	return rates, nil
}
