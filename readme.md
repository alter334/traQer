# Server パッケージ構成の見直し

## mainパッケージ
- 実行系のスクリプト
  - main.go
  - cron.go(定時実行関連)
  - root.go(ルーティング関連)
  - record.go(アプリ内情報保持関連)

## serviceパッケージ
- サービスそのものを書くイメージ 情報の整理,整形は基本このパッケージ
  - User.go(ユーザーごと情報取得,整理)
  - Channel.go
  - Stamp.go

## botパッケージ
- traQBOT関連のパッケージ
  - bot.go(セットアップ関連)
  - command.go(コマンド対応系)

## dbパッケージ
- db関連パッケージ
  - テーブル毎に処理を書く
    - e.g.)messagecount.go
  - dbセットアップもここ

## traQAPIパッケージ
- traQAPIを叩くパッケージ
  - どこまでを含めるかは要検討
  - メッセージならその取得まではこのパッケージでやりたい(serviceパッケージではこちらで定義したmodelですべて対応したい)
  - ID<->displayname変換などの処理もここ
  - BOTのメッセージ投稿も実際にはここにAPIを実装する(botパッケージではBotSimplePostのようなものを叩かせる)
