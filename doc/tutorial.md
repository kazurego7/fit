# チュートリアル.md

- [1. fit の目的](#1-fit-の目的)
- [2. fit のインストール](#2-fit-のインストール)
  - [2.1. fit 本体のインストール](#21-fit-本体のインストール)
  - [2.2. fit 補助機能のインストール](#22-fit-補助機能のインストール)
- [3. fit の利用方法](#3-fit-の利用方法)
  - [3.1. fit の利用準備](#31-fit-の利用準備)
    - [3.1.1. github リモートリポジトリの作成](#311-github-リモートリポジトリの作成)
    - [3.1.2. git ローカルリポジトリの初期化](#312-git-ローカルリポジトリの初期化)
    - [3.1.3. fit コマンド補完機能の利用](#313-fit-コマンド補完機能の利用)
  - [3.2. fit のワークフロー](#32-fit-のワークフロー)
    - [3.2.1. ブランチの作成](#321-ブランチの作成)
    - [3.2.2. ファイルの変更をステージングする](#322-ファイルの変更をステージングする)
    - [3.2.3. コミットの作成](#323-コミットの作成)
    - [3.2.4. mainブランチの変更を現在のブランチに反映](#324-mainブランチの変更を現在のブランチに反映)
    - [3.2.5. ブランチをリモートリポジトリへアップロード](#325-ブランチをリモートリポジトリへアップロード)
  - [3.3. その他の利用方法](#33-その他の利用方法)


## 1. fit の目的

git は幅広く使われているが、git の CLI は扱いづらい
fit は git で最もよく使われるコマンドを厳選し、予測しやすいものに改名し、コマンド体系を再編成したもの
つまり fit は単なる git のファサードコマンド

## 2. fit のインストール

### 2.1. fit 本体のインストール

### 2.2. fit 補助機能のインストール

## 3. fit の利用方法

以下の利用方法は、bash および powershell での実行を想定している
ただし、powershell の場合は、実行シェル上で以下の設定をしておく
```powershell
# デフォルトの文字コードをUTF8-BOM付きにする
$PSDefaultParameterValues['*:Encoding'] = 'utf8'
```

### 3.1. fit の利用準備

fit の利用の前に、ローカルに git リポジトリを作成し、github にリモートリポジトリを作成する

1. github リモートリポジトリの作成
2. git ローカルリポジトリの初期化
3. fit コマンド補完機能の利用

#### 3.1.1. github リモートリポジトリの作成

#### 3.1.2. git ローカルリポジトリの初期化

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
# git リポジトリの初期化
fit repository init
```

#### 3.1.3. fit コマンド補完機能の利用

### 3.2. fit のワークフロー

1. ブランチの作成
2. ファイルの変更をステージングする
3. コミットの作成
4. mainブランチの変更を現在のブランチに反映
5. ブランチをリモートリポジトリへアップロード

#### 3.2.1. ブランチの作成

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
```bash
# おっと間違えた！feat-foo_bar ブランチを作るんだった
```
```bash
# 現在のブランチ名を変更
fit branch rename feat-foo_bar
```
```bash
# 現在のブランチ名が feat-foo_bar に変更されたことを確認
fit branch list
```

#### 3.2.2. ファイルの変更をステージングする

```bash
# テキストファイルを作成する
mkdir hoge
echo "1st text contets dayo" > ./hoge/first.txt
echo "2nd text contets dayo" > ./hoge/second.txt
```
```bash
# 変更があったファイルの一覧を確認
fit change list
```
```bash
# hogeディレクトリ配下のファイルの変更をステージングする
fit change stage hoge
```
```bash
# 2つのファイルがステージングされたことを確認
fit change list
```
```bash
# おっと間違えた！second.txt をステージングしなくて良かったんだった
```
```bash
# second.txt ファイルのステージングをやめる
fit change unstage hoge/second.txt
```
```bash
# second.txt の変更がワークツリーに戻っていることを確認
fit change list
```

#### 3.2.3. コミットの作成

```bash
# 現在のブランチとコミットの履歴を確認
fit revision log
```
```bash
# インデックスの変更からコミットを作成
fit revision commit "最初のコミッ……"
```
```bash
# 新しいコミットが作られたことを確認
fit revision log
```
```bash
# おっと間違えた！コミットメッセージを修正しなきゃ
```
```bash
# コミットをキャンセルし、現在のブランチを一つ前のコミットに戻す
fit revision uncommit
```
```bash
# 現在のブランチが一つ前のコミットに戻っていることを確認
fit revision log
```
```bash
# コミットを再作成
fit revision commit "最初のコミットだよー"
```

#### 3.2.4. mainブランチの変更を現在のブランチに反映

他の人がリポジトリを更新しmainブランチが進んだ状態になった時、mainブランチの変更を現在のブランチに取り込みたいことがある
mainブランチを現在のブランチにマージすることで変更を取り込む

他の人が main ブランチにコミットした状態を再現する
```bash
# mainブランチに hoge/first.txt の追加をコミットする
fit revision switch main
echo "Other's chages dayo" > ./hoge/first.txt
fit change stage ./hoge/first.txt
fit revision commit "他の人のコミットだよー"
fit revision switch feat-foo_bar
```

mainブランチを現在のブランチにマージする
```bash
# main ブランチの状態を確認する
fit revision log
```
```bash
# main ブランチの指すコミットに含まれるファイルを表示する
fit revision show main
```
```bash
# mainブランチを現在のfeat-foo_barブランチにマージする
fit revision merge main
```
```bash
# おっと間違えた！second.txt の変更が残ったままだった
```
```bash
# second.txt の変更を破棄する
fit change delete hoge/second.txt
# コマンド実行後、バックアップの確認プロンプトが出るため "yes" を入力し Enter
```
```bash
# second.txt の変更が破棄されたことを確認する
fit change list
```
```bash
# mainブランチを現在のfeat-foo_barブランチにマージする
fit revision merge main
```
```bash
# おっと間違えた！マージコンフリクトが発生したぞ
```
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
```
```bash
# マージコミットが作成されたことを確認する
fit revision log
```

#### 3.2.5. ブランチをリモートリポジトリへアップロード
```bash
# リモートリポジトリからブランチ・タグ・コミットをダウンロードし、ローカルリポジトリの状態を最新にする
fit revision download

# ブランチをリモートリポジトリへアップロード
fit branch upload
```

### 3.3. その他の利用方法
