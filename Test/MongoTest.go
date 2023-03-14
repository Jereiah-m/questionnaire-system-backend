package Test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"questionnaire-system-backend/models"
	"testing"
	"time"
)

func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "root",
		Password: "WZnpXs77ny2rj6L0",
	}).ApplyURI("mongodb://nat.qwq.trade:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("MHQ")
	user := new(models.User)
	err = db.Collection("user").FindOne(context.Background(), bson.D{}).Decode(user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("user ==>", user)
}
