package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Jeffail/gabs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WriteAPIData to mongoDB
func WriteAPIData(JSONData *gabs.Container) string {

	record := make([]string, 6)
	record[0], _ = JSONData.Path("url").Data().(string)
	record[1], _ = JSONData.Path("product.name").Data().(string)
	record[2], _ = JSONData.Path("product.imageURL").Data().(string)
	record[3], _ = JSONData.Path("product.description").Data().(string)
	record[4], _ = JSONData.Path("product.price").Data().(string)
	totalReviews, _ := JSONData.Path("product.totalReviews").Data().(float64)
	record[5] = fmt.Sprintf("%g", totalReviews)

	//getting current time
	currentTime := time.Now()
	timestamp := currentTime.Format("2006.01.02 15:04:05")

	//connecting to mongoDB
	mongoURI := "mongodb+srv://sellerappapi:qwertyawsd@cluster0.c58ny.mongodb.net/sellerappDB?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return "errorinConnection"
	}
	defer client.Disconnect(ctx)

	//selecting db
	mongoDatabase := client.Database("sellerappDB")
	//selecting collection
	sellerappCollection := mongoDatabase.Collection("sellerappCollection")

	//update variable
	update := bson.M{
		"$set": bson.M{
			"url": record[0],
			"product": bson.M{
				"name":         record[1],
				"imageURL":     record[2],
				"description":  record[3],
				"price":        record[4],
				"totalReviews": record[5],
			},
			"dateUpdated": timestamp,
		},
	}

	//filter variable to find the record in mongoDB
	filter := bson.M{
		"url": record[0],
	}

	var urlmongo bson.M
	sellerappCollection.FindOne(ctx,
		filter,
	).Decode(&urlmongo)

	//if no record found then new record will be created in mongoDB
	if urlmongo == nil {
		result, err := sellerappCollection.InsertOne(ctx, bson.D{
			{Key: "url", Value: record[0]},
			{Key: "product", Value: bson.D{
				{Key: "name", Value: record[1]},
				{Key: "imageURL", Value: record[2]},
				{Key: "description", Value: record[3]},
				{Key: "price", Value: record[4]},
				{Key: "totalReviews", Value: record[5]},
			},
			},
			{Key: "dateCreated", Value: timestamp},
			{Key: "dateUpdated", Value: timestamp},
		})

		if err != nil {
			return "error"
		}

		log.Println(result.InsertedID)

	} else {

		//updating the record of the record exists in the DB
		_, err := sellerappCollection.UpdateOne(ctx,
			filter,
			update,
		)

		if err != nil {
			return "error"
		}

		return "update"

	}

	return "success"

}
