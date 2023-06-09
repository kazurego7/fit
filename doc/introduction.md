# git の紹介

- [1. git の目的](#1-git-の目的)
  - [1.1. コミットの作成](#11-コミットの作成)
  - [1.2. コミットの復元](#12-コミットの復元)
  - [1.3. コミットのマージ](#13-コミットのマージ)
- [2. git の機能](#2-git-の機能)
  - [2.1. リポジトリ](#21-リポジトリ)
  - [2.2. ワークツリー・インデックス](#22-ワークツリーインデックス)
  - [2.3. ブランチ](#23-ブランチ)
  - [2.4. ローカルリポジトリ・リモートリポジトリ](#24-ローカルリポジトリリモートリポジトリ)

## 1. git の目的

コミットは、ある瞬間のファイル・ディレクトリの状態のこと。  
git の目的は、コミットに対する以下の3つの操作ができること。  
- コミットの作成
- コミットの復元
- コミットのマージ

### 1.1. コミットの作成

コミットの作成とは、ある瞬間のファイル・ディレクトリの状態を保存し1つのコミットとして記録すること。  
例えば、ファイルを修正したいが古いものをバックアップにとっておきたいことがある。  
その際にバックアップとして、ファイルとディレクトリ全体をコミットとして保存することができる。  

![commit-create](image/introducion/commit-create.drawio.svg)

### 1.2. コミットの復元

コミットの復元とは、ファイルやディレクトリを特定のコミットの状態に復元すること。  
例えば、変更前のファイルの内容を確認したいことがある。  
その際に、過去のコミットに含まれるファイルとディレクトリの状態を復元することができる。  

![commit-restore](image/introducion/commit-restore.drawio.svg)

### 1.3. コミットのマージ

コミットのマージとは、複数のコミットを元にその内容をマージした1つのコミット作成すること。  
例えば、他人のファイルの変更と自分のファイルの変更をマージしたいことがある。  
その際に、それぞれの変更をコミットとして作成し、マージすることができる。  

![commit-merge](image/introducion/commit-merge.drawio.svg)

## 2. git の機能

### 2.1. リポジトリ

リポジトリは、git のデータの保管庫。  
コミット・ブランチなどの git のデータを保持している。(ブランチについては後述)  
特定のディレクトリをリポジトリにすることによって、そのディレクトリ配下でコミットの作成・マージ・復元を行うことができるようになる。  

![repository](image/introducion/repository.drawio.svg)

### 2.2. ワークツリー・インデックス

ワークツリーは、コミットの復元先のディレクトリ全体。  
デフォルトでは、リポジトリとして指定したディレクトリ全体がワークツリーとなる。  
コミットが変わる際、変わった先のコミットのファイルやディレクトリがワークツリーに復元される。  

![worktree](image/introducion/worktree.drawio.svg)

インデックスは、ファイルの変更を一時保存する場所。  
ワークツリーと異なり実際のディレクトリ上には見えず裏で管理されている。  
コミットが変わる際、変わった先のコミットのファイルやディレクトリがインデックスに復元される。  

![index](image/introducion/index.drawio.svg)

ワークツリーとインデックスの活用方法の1つとして、「コミットに保存するファイルの選別」がある。  
実際の開発では、複数の作業を一度に行って後からコミットを作成したいことがある。  
その際にまず、ワークツリー上の1作業分のファイルの変更をインデックスに登録する。  
その後、インデックスをコミットとして保存することで、作業ごとにファイルを選別してコミットを作成することができる。  

![index-commit](image/introducion/index-commit.drawio.svg)

また、ワークツリー上のファイルをインデックスに登録することを **「ステージング」** と呼ぶ。

### 2.3. ブランチ

ブランチは、開発の目印。  
ブランチは特定のコミットを指し示し、開発者はブランチ名からどのような機能を開発しているかを判別することができる。  
またコミットの作成時に、現在のブランチは作成されたコミットへと移動するため、ブランチは最新の開発コミットを指し示すことができる。  

![branch](image/introducion/branch.drawio.svg)

ワークツリー・インデックスの状態を特定のコミットに変更したい場合は、ブランチの切り替えを行う。  
現在のブランチを切り替えることで、そのブランチの指すコミットがワークツリー・インデックス上に復元される。  
コミットの変更は、ブランチの切り替えによって行うことが多い。  

![branch-switch](image/introducion/branch-switch.drawio.svg)


### 2.4. ローカルリポジトリ・リモートリポジトリ

ローカルリポジトリは、手元で作業するためのリポジトリ。  
リモートリポジトリは、ローカルリポジトリとは別の場所に配置されたリポジトリ。  
ローカルリポジトリでリモートリポジトリを設定することで、2つのリポジトリ間で git データ(コミット・ブランチなど)を共有することができる。  

![remote-download](image/introducion/remote-download.drawio.svg)

例えば、チームで各々のローカルリポジトリに共有のリモートリポジトリを設定することで、チーム間でコミットやブランチを共有することができる。  

![remote-upload](image/introducion/remote-upload.drawio.svg)