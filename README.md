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

- [git](https://git-scm.com/downloads) version 2.25.1 以降

## インストール方法

[最新のリリース](https://github.com/kazurego7/fit/releases/latest)

### Windows

以下、Powershell での操作

ログインユーザーのコマンドの保存先ディレクトリを作成し、そこへ最新のfitをダウンロードする
```powershell
mkdir ~/bin
Invoke-WebRequest -Uri https://github.com/kazurego7/fit/releases/latest/download/fit.exe -OutFile ~/bin/fit.exe
```

保存先ディレクトリへのパスを通す
```powershell
$new_path = [Environment]::GetEnvironmentVariable("Path", "User")
$new_path += ";$HOME/bin"
[Environment]::SetEnvironmentVariable("Path", $new_path, "User")
# powershell を再起動する
```

### Linux

以下、Bash での操作

ログインユーザーのコマンドの保存先ディレクトリを作成し、そこへ最新のfitをダウンロードする
```bash
mkdir ~/bin
wget https://github.com/kazurego7/fit/releases/latest/download/fit -P ~/bin/
```

ログインユーザーがfitを利用できるように、パーミッションを修正する
```bash
chmod u+x ~/bin/fit
```

## 利用方法

```bash
fit --help
```

## ライセンス

fit は MIT ライセンスです。
[LICENSE](LICENSE)を確認してください。