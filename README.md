# knowledge-like

## 動かし方

```
$ export GO111MODULE=on
$ go get
$ go run ./
```

### フロントエンド修正したら

```
$ cd frontend
$ npm run build
```

## 設計

### 大方針

- フロントエンドはvuejsで作る
- バックエンドはgolangでAPIサーバ作る
  - public はフロントエンドで作成した静的コンテンツを配置
  - API部分はJWT検証

### 採用技術

#### フロントエンド

- vue.js
  - **vuetify** :  UIコンポーネント
  - **axios** : HTTPクライアント

#### バックエンド

  - golang
    - **echo** Webフレームワーク
    - **gorm** ORマッパー
    - **go-jwt** JWT

### model

- user
  - id, name,password(暗号化して保存)
- knowledge
  - id, title, url, author
- like
  - user_id, knowledge_id, reason, comment

### controller

apiのハンドラーとして実装するもの

#### user

- サインアップ
  - /singup POST
- サインイン
  - /singin POST
  - JWT返す

#### knowledge

- knowledge一覧取得
  - /knowledge GET
- knowledge投稿
  - /knowledge POST

#### like

- 一覧取得
  - /like
- いいね取得
  - /like/{user_id} GET
- いいね投稿
  - /like/{user_id} POST

### view

画面遷移作るとわかりやすい

- サインアップ画面
- サインイン画面
- いいね投稿画面
  - knowledge一覧、いいね取得、いいね投稿
  - いいね投稿は３つまで
- いいね結果画面
  - いいね一覧取得
  - いいね詳細画面も

## Tips

- vue.jsとechoを繋げる

echo側はpublicフォルダを静的コンテンツとして設定する。

router.go
```
	// vue.js static files
	e.Static("/", "public/")
```

vue.js側は`npm run build`でpublicフォルダに出力するようにする。

vue.config.js
```
module.exports = {
  outputDir: '../public'
```
