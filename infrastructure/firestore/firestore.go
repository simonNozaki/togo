package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"log"
)

func InitializeFirestore() (*firestore.Client, error) {
	// APIキーや秘密鍵といった秘匿系情報は環境変数 GOOGLE_APPLICATION_CREDENTIALS から暗黙的に吸い取られる
	// 環境変数から吸い取るのはfirebase推奨なのでそちらに則る: https://firebase.google.com/docs/admin/setup?hl=ja#initialize_the_sdk_in_non-google_environments
	var ctx = context.Background()
	var config = &firebase.Config{
		ProjectID: "cords-5647f",
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing client: %v\n", err)
	}
	return client, err
}
