# traQer
traQ Database viewer
traQでの発言履歴,発言数,スタンプ数など"話題性"につながる指標を抽出し、traQにおける発言の傾向および勢いを分析するためのデータを提供する。
収集されたデータを見やすい形でまとめ,利用しやすい形で提供することも目標とする。

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
