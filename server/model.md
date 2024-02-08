# サーバ設計関連資料

## パッケージ
- [readme](../README.md)参照
![](../IMG_7693.jpeg)

### 命名規則
- パッケージ名は全て小文字
- ハンドラ名,struct名はパスカルケース
- 変数名はキャメルケース
- メソッド名はパスカルケース
- DB,API等の略語はパッケージ名の時以外は大文字とする

## Botパッケージ
- botのコマンドやbotからのサーバへのアクセスを管理
- 他のパッケージからは独立したものとして扱う(他のパッケージのメソッドにBotパッケージの関数は入らない)
  - Botからサーバのメソッドを呼び出す イメージとしては`service.User.BotHoge()`のように記述する
  - Service系パッケージ内にBOTから呼び出されるメソッドを実装する

### Bot
```
type Bot struct {
  bot *traqwsbot.Bot // traqBot
}
```

## APIパッケージ
- db,qapiパッケージを格納するパッケージ
- 設計上は最下層の一つ上に位置する
- サーバの各サービスごとにインスタンス化(?)される
- サーバの持つデータもここに格納され、そのデータを読み取るメソッドが生やされる

### API
```
type ApiHandler struct{
  db *db.DB // db関連
  qapi *qapi.Qapi // traqApi関連
  serverData ServerData // サーバの持つデータ
}
```

### ServerData
```
type ServerData struct {
  lastTrackMessage traq.Message // 最後に取得したメッセージ
  lastTrackTime time.Time // 最後の取得日時
  // 増えたらここに書く
}
```

## DBパッケージ
- dbのセットアップ及びdb操作をするメソッドの集まったパッケージ
- 設計上の最下層

### DB
```
type DB struct{
  db *sqlx.DB // sqlx.db
}
```

## Qapiパッケージ
- traQのAPIラッパーを用いてAPIを叩くメソッドの集まったパッケージ
- 設計上の最下層

### Qapi
```
type Qapi struct {
  auth context.Context // 認証関連
  client *traq.APIClient // traqAPIClient
}
```



## Serviceパッケージ
- メソッドは略
- 各パッケージにhandler.goを用意しそこにインスタンス化を定義 model.goも用意

### Service
```
type Service struct{
  User *user.UserHandler
  Channel *channel.ChannelHandler
  Stamp *stamp.StampHandler
}
```
## Userパッケージ

### UserHandler
```
type UserHandler struct{
  Api *api.APIHandler
}
```

### ChannelHandler
```
type ChannelHandler struct{
  Api *api.APIHandler
}
```

### StampHandler
```
type StampHandler struct{
  Api *api.APIHandler
}
```

## setupパッケージ
- bot,serviceの起動を行う(Handlerの設定)
- cronの設定,echoの設定を行う

### Setup
```
type Setup struct{
  bot *bot.Bot
  service *service.Service
}
```

## mainパッケージ
- 実行用


