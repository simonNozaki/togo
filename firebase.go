package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"time"
	"togo-web/domain/data"
	"togo-web/infrastructure/firestore"
)

func main() {
	var client, _ = firestore.InitializeFirestore()
	iter := client.Collection("todos").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO エラー処理
		}

		createdAt, err := time.Parse("2006-01-02", doc.Data()["createdAt"].(string))
		updatedAt, err := time.Parse("2006-01-02", doc.Data()["updatedAt"].(string))
		if err != nil {
			panic(fmt.Sprintf("updateAtがパースできません: %s", err))
		}
		var v = data.Todo{
			Id:          doc.Data()["id"].(string),
			UserId:      doc.Data()["userId"].(string),
			Title:       doc.Data()["title"].(string),
			Description: doc.Data()["description"].(string),
			State:       data.GetState(doc.Data()["state"].(string)),
			CreatedAt:   createdAt,
			UpdateAt:    updatedAt,
		}
		var data, _ = json.Marshal(v)
		fmt.Printf("%s\n", data)
	}
}
