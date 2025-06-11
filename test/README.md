# テスト環境

## 目的
- テストを何度も繰り返すため、自動化したい
- テスト用の git リポジトリを素早く準備したい
- テスト用の git リポジトリの修正も git で管理したい

## 手法

### 方針
1. `setup_script` に配置したシェルスクリプトからテスト用の git リポジトリを作成する
2. `repo_setup.sh` (本ディレクトリ直下) を通じて各スクリプトを実行し、`repos` ディレクトリにリポジトリを生成する

### 詳細
- スクリプト名がそのまま作成されるリポジトリ名となる
- "_" で始まるスクリプトは無視される
- 生成されたリポジトリは `repos` ディレクトリに保存され、`git` 管理対象外となる

## テストまでの流れ
1. テストリポジトリのビルドスクリプトを `setup_script` に配置
2. `repo_setup.sh <スクリプトパス>` を実行し `repos` 配下にリポジトリを生成
3. 生成されたリポジトリを確認
4. 必要に応じて `go test` などでテストを実行

### ビルドスクリプト例
```bash
#!/bin/bash
set -eu

git init

echo "hogehoge" > hoge.txt
git add -A
git commit -m "hogeeee"

echo "fugafuga" > fuga.txt
git add -A
git commit -m "fugaaaa"
```
