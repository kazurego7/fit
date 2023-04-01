# README.md

## 目的
- テストを何度も繰り返すため、自動化したい
- テスト用の git リポジトリを素早く準備したい
- テスト用の git リポジトリの修正も git で管理したい

## 手法

### 方針
1. コンテナ上で、テスト用の git リポジトリをビルドスクリプトによって作成し、 docker イメージとして固める
2. テストのたびにコンテナを起動、コンテナ内でテストを実行する

### 詳細
- ビルド時間短縮のためステージを分ける([マルチステージビルドの利用](https://matsuand.github.io/docs.docker.jp.onthefly/develop/develop-images/multistage-build/))
  - テストリポジトリのビルドステージ
  - テスト実行ステージ

## テストまでの流れ
1. テストリポジトリのビルドスクリプトを配置
2. 各ステージごとのイメージをビルド
3. テストリポジトリが作成できたかの確認
4. テスト実行

## 1. テストリポジトリのビルドスクリプトを配置

ビルドスクリプトを `build/build_script` に配置する  
※"_"始まりのスクリプトは無視される

ビルドスクリプト `hogehoge.sh` の例
```bash
#!/bin/bash 
set -eu # エラーおよび未定義変数があれば停止

git init

echo "hogehoge" > hoge.txt
git add -A
git commit -m "hogeeee"

echo "fugafuga" > fuga.txt
git add -A
git commit -m "fugaaaa"
```

## 2. 各ステージごとのイメージをビルド
テストリポジトリのビルドステージ、テスト実行ステージごとにイメージをビルドする
```
docker-compose build test-repo-build test-run
```

## 3. テストリポジトリが作成できたかの確認
コンテナ上の bash に接続し、テストリポジトリが作成されているかを確認する
```
docker-compose run test-repo-build
```

## 4. テスト実行
コンテナ上の node でテストが実行される
```
docker-compose run test-run
```