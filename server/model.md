# サーバ設計関連資料

## パッケージ
- [readme](../README.md)参照

## struct関連
- メソッドは略
- 各パッケージにhandler.goを用意しそこにインスタンス化を定義 model.goも用意

### Server
```
type Server struct{
  bot *Bot // Bot関連
  db *DB // db関連
  service *Service // 具体的機能
  traQapi *traQapi // traqApi関連
  serverData ServerData // 保持しておくデータ
}
```

### ServerData
```
type ServerData struct {
  lastTrackId traq.Message // 最後に取得したメッセージ
  lastTrackTime time.Time // 最後の取得日時
  // 増えたらここに書く
}
```

### Bot
```
type Bot struct {
  bot *traqwsbot.Bot // traqBot
}
```

### DB
```
type DB struct{
  db *sqlx.DB // sqlx.db
}
```

### traQapi
```
type traQapi struct {
  auth context.Context // 認証関連
  client *traq.APIClient // traqAPIClient
}
```

