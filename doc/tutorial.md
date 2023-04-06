<!-- no toc -->
# チュートリアル.md

- [1. fit の目的](#1-fit-の目的)
- [fit のインストール](#fit-のインストール)
  - [fit 本体のインストール](#fit-本体のインストール)
  - [fit 補助機能のインストール](#fit-補助機能のインストール)
- [2. fit の利用方法](#2-fit-の利用方法)
  - [2.1. git の利用準備](#21-git-の利用準備)
    - [2.1.1. git ローカルリポジトリの初期化](#211-git-ローカルリポジトリの初期化)
  - [2.2. fit のワークフロー](#22-fit-のワークフロー)
    - [ブランチの作成](#ブランチの作成)
    - [ファイルの変更をステージングする](#ファイルの変更をステージングする)


## 1. fit の目的

git は幅広く使われているが、git の CLI は扱いづらい
fit は git で最もよく使われるコマンドを厳選し、予測しやすいものに改名し、コマンド体系を再編成したもの
つまり fit はただの git のラッパーコマンド

## fit のインストール

### fit 本体のインストール

### fit 補助機能のインストール

## 2. fit の利用方法

### 2.1. git の利用準備

fit の利用の前に、ローカルに git リポジトリを作成し、github にリモートリポジトリを作成する

1. git ローカルリポジトリの初期化
2. github リモートリポジトリの作成

#### 2.1.1. git ローカルリポジトリの初期化

```bash
# ホームディレクトリに移動
cd $HOME

# fit チュートリアル用のディレクトリを作成
mkdir fit_practice

# git リポジトリの初期化
fit repository init
```



### 2.2. fit のワークフロー

1. ブランチの作成
2. ファイルの変更をステージングする
3. コミットの作成
4. ローカルでのマージ
5. ブランチをリモートリポジトリへアップロード

#### ブランチの作成

```bash
# ブランチ一覧から現在のブランチが main となっていることを確認
fit branch list

# feat-hoge ブランチを作成
fit branch create feat-hoge

# feat-hoge ブランチが作成され、現在のブランチが feat-hoge であることを確認
fit branch list

# おっと間違えた、feat-foo ブランチを作るんだった……

# 現在のブランチ名を変更
fit branch rename feat-foo
```

#### ファイルの変更をステージングする

```bash
# ファイルを編集
echo "first, file" > ./first.txt
echo "second file" > ./second.txt

# 変更があったファイルの一覧を確認
fit change list

# 変更があったファイル内容の差分を確認
fit change show first.txt

# ファイルの変更をステージングする
fit change stage ./

# ファイルがステージングされたことを確認
fit change list

# あらら、second.txt をステージングしなくて良かったのに……

# second.txt ファイルのステージングをやめる
fit change unstage second.txt
```