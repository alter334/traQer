# traQer
traQ Database viewer
traQでの発言履歴,発言数,スタンプ数など"話題性"につながる指標を抽出し、traQにおける発言の傾向および勢いを分析するためのデータを提供する。
収集されたデータを見やすい形でまとめ,利用しやすい形で提供することも目標とする。
[サーバ](./server/model.md)
![](./IMG_7693.jpeg)


## Server パッケージ構成の見直し

## mainパッケージ(もしくはsetup)
- 実行系のスクリプト
  - main.go
  - cron.go(定時実行関連)
  - root.go(ルーティング関連)
  - record.go(アプリ内情報保持関連)

## serverパッケージ
- サービスそのものを書く 情報の整理,整形は基本このパッケージ
  - User.go(ユーザーごと情報取得,整理)
  - Channel.go
  - Stamp.go
- serviceの配下に以下各種パッケージがある

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

## 今後具体的なクラス相当の設計を考える

## やりたいこと

### traQ発言数量データの収集>>traQAPIの利用 フロント実装だけでいけると思う

- 投稿数
  - 個人投稿数(Daily,Weekly,Total)
  - チャンネル毎投稿数(Daily,Weekly,Total)
  - リアルタイム流速(5分単位毎にアクティビティみたいな感じで盛り上がってるチャンネルを表示)
- スタンプ数
  - スタンプ毎数ランキング
  - スタンプユーザーランキング
- 話題  
  - トレンドワード表示

### 収集データのtraQ利用体験向上への利用

- 収集データのまとめサイト
  - データを表,グラフを用いてなんかいい感じに表示させたい
  - フロントエンド実装頑張る
- 話題,traQ活発化通知
  - gazerみたいにBOTとか使っていい感じにやりたいよね
  - 一定流速以上を観測すると通知,特定チャンネルでの特定のワードで反応するetc  
