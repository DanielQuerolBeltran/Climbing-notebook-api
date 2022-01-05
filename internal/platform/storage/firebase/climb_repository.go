package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	entities "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	collectionKeyId          = "id"
	collectionKeyDescription = "description"
	collectionKeyDate        = "date"
	collectionKeyGrade       = "grade"
	collectionKeyArea        = "area"
)

type repository struct {
	projectId      string
	collectionName string
}

func NewRepository() entities.ClimbRepository {
	return &repository{
		projectId:      "go-testing-b6275",
		collectionName: "climbs",
	}
}

func getConnection(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsJSON([]byte(`
	{
	  }
	`))
	
	//sa := option.WithCredentials("go-testing-firebase.json")
	app, _ := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(ctx)

	return client, err
}

func (r *repository) Save(ctx context.Context, climb entities.Climb) error {
	client, err := getConnection(ctx)

	if err != nil {
		return fmt.Errorf("error trying to connecto to firebase: %v", err)
	}

	defer client.Close()

	_, _, err = client.Collection(r.collectionName).Add(ctx, map[string]interface{}{
		collectionKeyId:          climb.Id().String(),
		collectionKeyDescription: climb.Description().String(),
		collectionKeyArea:        climb.Area().String(),
		collectionKeyGrade:       climb.Grade().String(),
		collectionKeyDate:        climb.Date().String(),
	})

	if err != nil {
		return fmt.Errorf("error trying to persist climb on database: %v", err)

	}

	return nil
}

func (r *repository) Get(ctx context.Context, id entities.ClimbId) ([]entities.Climb, error) {
	client, err := getConnection(ctx)

	if err != nil {
		return nil, fmt.Errorf("error trying to connecto to firebase: %v", err)

	}

	defer client.Close()

	iter := client.Collection(r.collectionName).Documents(ctx)
	var climbs []entities.Climb

	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error trying to create a Climb: %v", err)

		}

		climb, err := entities.NewClimb(
			doc.Data()[collectionKeyId].(string),
			doc.Data()[collectionKeyDate].(string),
			doc.Data()[collectionKeyGrade].(string),
			doc.Data()[collectionKeyDescription].(string),
			doc.Data()[collectionKeyArea].(string),
		)

		if err != nil {
			return nil, fmt.Errorf("error trying to fetch climbs from database: %v", err)
		}

		climbs = append(climbs, climb)
	}

	return climbs, nil
}
