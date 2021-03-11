package mongodb

import (
	"amiera/src/domain/access_token"
	"amiera/src/utils/utils_errors"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
func ConnectToAccessToken() (*mongo.Client, context.Context, *utils_errors.RestErr) {
	uri := "mongodb://localhost:27017/?connect=direct"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, utils_errors.CustomInternalServerError("Cannot connect to database")
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	fmt.Println("was executed")
	//dbUsers := client.Database("db_oauth")
	//accessTokenCollection := dbUsers.Collection("access_token")

	//return accessTokenCollection, ctx, nil
	return client, ctx, nil
}
*/

func GetAllAccessToken() ([]access_token.AccessToken, *utils_errors.RestErr) {

	uri := "mongodb://localhost:27017/?connect=direct"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, utils_errors.CustomInternalServerError(err.Error())
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	accessTokenCollection := client.Database("db_oauth").Collection("access_token")
	//dbUsers := client.Database("db_oauth")
	//accessTokenCollection := dbUsers.Collection("access_token")

	cursor, err := accessTokenCollection.Find(ctx, access_token.AccessToken{})
	if err != nil {
		fmt.Println("Err Find data: ", err)
		return nil, utils_errors.CustomInternalServerError(err.Error())
	}

	defer cursor.Close(ctx)

	var accessToken []access_token.AccessToken
	if err = cursor.All(ctx, &accessToken); err != nil {
		fmt.Println("Err Read Cursor ", err)
		return nil, utils_errors.CustomInternalServerError(err.Error())
	}

	//example to read record
	/*
		for idx, strRecord := range accessToken {
			fmt.Print("idx# ", idx)
			fmt.Println(" -> ", strRecord)
		}
	*/
	//loop manual cursor
	/*
		for cursor.Next(ctx) {
			var accToken access_token.AccessToken
			if err = cursor.Decode(&accToken); err != nil {
				log.Fatal(err)
			}
			fmt.Println("accToken ", accToken)
		}
	*/

	return accessToken, nil
}

func GetAccessTokenById(acId int64) (*access_token.AccessToken, *utils_errors.RestErr) {
	uri := "mongodb://localhost:27017/?connect=direct"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, utils_errors.CustomInternalServerError(err.Error())
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	accessTokenCollection := client.Database("db_oauth").Collection("access_token")

	//var accTokenUnique bson.M
	var accTokenUnique access_token.AccessToken
	if err = accessTokenCollection.FindOne(ctx, access_token.AccessToken{UserId: acId}).Decode(&accTokenUnique);
		err != nil {
		fmt.Println("Error ", err)
	}

    if accTokenUnique.UserId <= 0 {
		return nil, utils_errors.CustomNotFoundError(fmt.Sprintf("Access Token Id: %d not found.", acId))
	}
	return &accTokenUnique, nil
}
