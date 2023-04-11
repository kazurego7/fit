# fit

![fit のヘルプ画像](./doc/image/readme/fit-title.png)

ユーザーフレンドリーな git CLI

## 概要

> ある初心者がマスター・ギットの元で学んでいた。レッスンの終わりに彼はノートに目を通して言った。  
「師匠、いくつか質問があるんです。聞いてもいいですか？」  
マスター・ギットは頷いた。  
「すべてのタグのリストを見るにはどうしたらいいですか？」  
「git tag」と、マスター・ギットは答えた。  
「すべてのリモートのリストを見るにはどうしたらいいですか？」  
「git remote -v」と、マスター・ギットは答えた。  
「すべてのブランチのリストを見るにはどうしたらいいですか？」  
「git branch -a」と、マスター・ギットは答えた。

途中省略  

>初心者はしばらく考えてから尋ねた。  
「この中のいくつかはもっと一貫性を持たせて、コーディングの最中に思い出しやすくすることができるんじゃないですか？」  
マスター・ギットは指を鳴らした。ホブゴブリンが部屋に入ってきて、初心者を生きたまま食べてしまった。あの世で、初心者は悟りを開いた。  

[Git Koans - The Hobgoblin](https://stevelosh.com/blog/2013/04/git-koans/#s4-the-hobgoblin)

git のコマンドは初心者殺しです。  
fit は、 git で最もよく使うコマンドのみを厳選し、予測しやすいものに改名し、コマンド体系を再編成した CLI です。

## 必須要件

- git version 2.25.1 以降
- go version 1.20.0 以降

## インストール方法

[gitのインストール](https://git-scm.com/downloads)

[goのインストール](https://go.dev/doc/install)

最新のfitをインストールする。
```
go install github.com/kazurego7/fit/fit@latest
```

## 利用方法

```bash
fit --help
```

[ユーザーマニュアル](./doc/manual.md)も併せてご確認ください。

## ライセンス

fit は MIT ライセンスです。
[LICENSE](LICENSE)を確認してください。