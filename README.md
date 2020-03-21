# 2929BE

とりあえずつくってみる

## Description

- DDDライクな構成を模索する（まだ）
- GraphQLサーバーを習得する

## 技術

- Go 1.14.0
    - Go Modules
- GraphQL
    - gqlgen v0.11.3
- docker
    - multi stage build
    - docker-compose
- Make（まだ）

## Demo

ユーザー登録

- `docker-compose up`
    - user,mysqlコンテナ起動
- `localhost:8080` にアクセス

### リクエスト

```
# ユーザー認証
# ユーザー登録
# ユーザー更新
# ユーザー削除
```

※curlだと以下

```
$ 
```

### 改修方法

```
# 事前準備
## ローカルでGraphQLサーバーのコード生成するため、gqlgenをインストール
$ go mod download
```

```
# リゾルバ修正後
$ go run github.com/99designs/gqlgen generate
```

## Tutorial

本リポジトリの利用技術を理解するために、各チュートリアルを実施する

参考記事は独断と偏見によるものなので、周辺知識は別途ググる

### 0. Go

[環境構築]
https://qiita.com/yosemite2307/items/08dce692894c92ae08ee

[Go Modules]
https://blog.mmmcorp.co.jp/blog/2019/10/10/go-mod/

`export GO111MODULE=on` が出来ていればOK

### 1. gqlgen

https://gqlgen.com/getting-started/

### 2. GraphQL

https://qiita.com/SiragumoHuin/items/cc58f456bc43a1be41b4

### 3. database/sql

https://noumenon-th.net/programming/2019/09/20/go-sql-driver/

## わからなかったこと

### GraphQL

- `ID：一意の識別子であることを示す以外はStringと同じ` らしいが、同じ文字列を与えてもバリデーションエラーにならない
- GraphQLのエラーの出し方がわからない

## Todo

- `gqlgen generate` Make化
- GraphQL本を読もう...
