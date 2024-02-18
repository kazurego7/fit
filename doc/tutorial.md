# チュートリアル

- [1. 注意事項](#1-注意事項)
- [2. fit の利用準備](#2-fit-の利用準備)
  - [2.1. github リモートリポジトリの作成](#21-github-リモートリポジトリの作成)
  - [2.2. git ローカルリポジトリの初期化](#22-git-ローカルリポジトリの初期化)
- [3. fit の利用の流れ](#3-fit-の利用の流れ)
  - [3.2. mainブランチからfeatureブランチを作成する](#32-mainブランチからfeatureブランチを作成する)
  - [3.3. 変更したファイルをインデックスにステージングする](#33-変更したファイルをインデックスにステージングする)
  - [3.4. コミットを作成する](#34-コミットを作成する)
  - [3.5. mainブランチの変更を現在のブランチに反映する](#35-mainブランチの変更を現在のブランチに反映する)
  - [3.6. マージコンフリクトを解消する](#36-マージコンフリクトを解消する)
  - [3.7. ブランチをリモートリポジトリへアップロードする](#37-ブランチをリモートリポジトリへアップロードする)
- [4. その他の利用方法](#4-その他の利用方法)
  - [4.1. ブランチ名を誤ってブランチを作成した ⇒ ブランチ名の変更](#41-ブランチ名を誤ってブランチを作成した--ブランチ名の変更)
  - [4.2. 違うファイルをステージングした ⇒ ステージングの解除](#42-違うファイルをステージングした--ステージングの解除)
  - [4.3. ファイルの変更を破棄したい](#43-ファイルの変更を破棄したい)
  - [4.4. 現在のコミットを修正したい ⇒ コミットのキャンセル](#44-現在のコミットを修正したい--コミットのキャンセル)


## 1. 注意事項

以下のチュートリアルは、bash および powershell での実行を想定しています。  
ただし、powershell の場合は、実行シェル上で以下の設定をしています。

```powershell
# デフォルトの文字コードをUTF8-BOM付きにする
$PSDefaultParameterValues['*:Encoding'] = 'utf8'
```

## 2. fit の利用準備

fit の利用の前に、ローカルに git リポジトリを作成し github にリモートリポジトリを作成しておく必要があります。

1. github リモートリポジトリの作成
2. git ローカルリポジトリの初期化

### 2.1. github リモートリポジトリの作成

以下の手順を参考に、リモートリポジトリを作成してください。

[リポジトリを作成する - GitHub Docs](https://docs.github.com/ja/get-started/quickstart/create-a-repo)

### 2.2. git ローカルリポジトリの初期化

```bash
# ホームディレクトリに移動
cd $HOME
```
```bash
# fit チュートリアル用のディレクトリを作成し、移動
mkdir fit_practice
cd fit_practice
```
```bash
# ローカルの git リポジトリの初期化
fit repository init
```
```bash
# github のリモートリポジトリとの非同期接続を設定する
fit repository connect [githubのリポジトリのURL]
```
```bash
# リモートリポジトリの設定が完了したことを確認する
fit repository remote
```

## 3. fit の利用の流れ

GitLab Flow の場合、以下のような流れで fit を利用します。

1. mainブランチからfeatureブランチを作成する
2. 変更したファイルをインデックスにステージングする
3. コミットを作成する
4. mainブランチの変更を現在のブランチに反映する
5. マージコンフリクトを解消する
6. ブランチをリモートリポジトリへアップロードする

### 3.2. mainブランチからfeatureブランチを作成する

```bash
# ブランチ一覧から現在のブランチが main となっていることを確認
fit branch list
```
```bash
# feat-hoge ブランチを作成
fit branch create feat-hoge
```
```bash
# feat-hoge ブランチが作成され、現在のブランチが feat-hoge であることを確認
fit branch list
```

### 3.3. 変更したファイルをインデックスにステージングする

```bash
# ワークツリーにテキストファイルを作成する
mkdir hoge
echo "1st text contets dayo" > ./hoge/first.txt
echo "2nd text contets dayo" > ./hoge/second.txt
```
```bash
# 変更のあるファイルの一覧を確認
fit change list
```
```bash
# ./hoge 配下のファイルをインデックスにステージングする
fit change stage ./hoge
```
```bash
# first.txt と second.txt がステージングされたことを確認
fit change list
```

### 3.4. コミットを作成する

```bash
# 現在のブランチと、コミットの履歴を確認
fit commit list
```
```bash
# インデックスの変更からコミットを作成
fit commit create "最初のコミットだよー"
```
```bash
# 新しいコミットが作られたことを確認
fit commit list
```

### 3.5. mainブランチの変更を現在のブランチに反映する

準備として、他の人が main ブランチにコミットを作成した状態を再現します。  
```bash
# mainブランチに hoge/first.txt が追加されたコミットを作成する
fit branch switch main
mkdir hoge
echo "Other's chages dayo" > ./hoge/first.txt
fit change stage ./hoge/first.txt
fit commit create "他の人のコミットだよー"
# featureブランチに移動する
fit branch switch feat-hoge
```

mainブランチを現在のブランチにマージします。
```bash
# main ブランチにコミットが追加されていることを確認する
fit commit list
```
```bash
# main ブランチの指すコミットに含まれるファイルを表示する
fit commit list --details main
```
```bash
# main ブランチを現在の feat-hoge ブランチにマージする
fit commit merge main
```

### 3.6. マージコンフリクトを解消する

```bash
# マージコンフリクトが発生しているファイルを確認する
fit conflict list
```
```bash
# first.txt のマージコンフリクトを解消する
# 好きなエディタで first.txt のコンフリクトマーカーを消す
```
```bash
# first.txt のマージコンフリクトを解消し、ステージングする
fit change stage hoge/first.txt
```
```bash
# マージコンフリクトを解消し、マージコミットを作成する
fit conflict resolve
# マージコミットのメッセージ編集エディタが開くため、メッセージを編集しエディタを閉じる
```
```bash
# マージコミットが作成されたことを確認する
fit commit list
```

### 3.7. ブランチをリモートリポジトリへアップロードする
```bash
# リモートリポジトリからブランチ・タグ・コミットをダウンロードし、ローカルリポジトリの状態を最新にする
fit branch follow

# ブランチをリモートリポジトリへアップロード
fit branch upload
```

## 4. その他の利用方法

### 4.1. ブランチ名を誤ってブランチを作成した ⇒ ブランチ名の変更

```bash
# 現在のブランチ名を変更
fit branch rename feat-foo_bar
```
```bash
# 現在のブランチ名が feat-foo_bar に変更されたことを確認
fit branch list
```

### 4.2. 違うファイルをステージングした ⇒ ステージングの解除

```bash
# second.txt ファイルのステージングをやめる
fit change unstage hoge/second.txt
```
```bash
# second.txt の変更がワークツリーに戻っていることを確認
fit change list
```

### 4.3. ファイルの変更を破棄したい

```bash
# second.txt の変更を破棄する
fit change delete hoge/second.txt
```
```bash
# second.txt の変更が破棄されたことを確認する
fit change list
```

### 4.4. 現在のコミットを修正したい ⇒ コミットのキャンセル

```bash
# コミットをキャンセルし、現在のブランチを一つ前のコミットに戻す
fit commit back
```
```bash
# 現在のブランチが一つ前のコミットに戻っていることを確認
fit commit list
```

