package storage

import (
	"context"
	"fmt"

	"github.com/Feinot/authorization-service/internal/entity"
	"github.com/Feinot/authorization-service/internal/modules/hash"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckDb(refreshToken string, guid string, t *entity.Tokens) (bool, error) {
	err := t.DB.Ping(context.Background(), nil)
	if err != nil {
		return false, fmt.Errorf("cannot pinget db %v", err)
	}
	collection := t.DB.Database("UserSession").Collection("sessions")
	if err != nil {
		return false, fmt.Errorf("cannot delete entry: %v", err)
	}
	cursor, err := collection.Find(context.Background(), bson.M{"user": bson.M{"guid": guid}})
	if err != nil {
		return false, fmt.Errorf("cannot find entry: %v", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			return false, fmt.Errorf("cannot decodding: %v", err)
		}
		str := fmt.Sprintf("%v", episode["refreshtoken"])
		res := hash.CheckTokenHash(refreshToken, str)
		if res {
			id := episode["_id"]
			_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
			if err != nil {
				return false, fmt.Errorf("cannot delete entry: %v", err)
			}
			return true, nil
		}
	}
	return false, nil
}
func CreateDb(u entity.User, at entity.AuthToken, t *entity.Tokens) error {
	err := t.DB.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("cannot ping data base: %v", err)
	}
	collection := t.DB.Database("UserSession").Collection("sessions")
	var s entity.Session
	s.User = u
	s.RefreshToken, err = hash.HashToken(at.RefreshToken)
	if err != nil {
		return fmt.Errorf("cannot add refreshToken: %v", err)
	}
	_, err = collection.InsertOne(context.Background(), s)
	if err != nil {
		return fmt.Errorf("cannot insert in coolection: %v", err)
	}
	return nil
}
