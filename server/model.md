# サーバ設計関連資料

## パッケージ
- [readme](../README.md)参照

## struct関連
- メソッドは略

### Server
```
type struct Server{
  bot *Bot // Bot関連
  db *DB // db関連
  service *Service // 具体的機能
  traQapi *traQapi // traqApi関連
  serverData ServerData // 保持しておくデータ
}
```

### ServerData
```
type struct ServerData{
  lastTrackId traq.Message // 最後に取得したメッセージ
  // 増えたらここに書く
}
```

### Bot
type struct Bot{
  bot *traqwsbot.Bot // traqBot
}

### traQapi
type struct traQapi{
  auth context.Context // 認証関連
  client *traq.APIClient // traqAPIClient
}
