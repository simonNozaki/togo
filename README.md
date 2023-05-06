# ToGo

## セットアップ
### Firestoreの設定
Firestore推奨の設定に則り、設定した環境変数を参照して接続を試みる。アプリ内では個別の設定を入れていないので注意する。

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/json/cords-5647f-firebase-adminsdk-kkvgd-6670459692.json"
```

参考: [Google 以外の環境で SDK を初期化する](https://firebase.google.com/docs/admin/setup?hl=ja#initialize_the_sdk_in_non-google_environments)

### ローカル開発
**todo**

登録 `POST /todo`

```json
{
    "id": "3",
    "userId": "1",
    "title": "task",
    "description": "from postman",
    "state": "unprocessed",
    "createdAt": "2023-05-06",
    "createdBy": "postman",
    "updatedAt": "2023-05-06",
    "updatedBy": "postman"
}
```

1件取得 `GET /todo/:id`

全県取得 `GET /todo`
