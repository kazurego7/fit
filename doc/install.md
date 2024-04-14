# インストール方法

## 必須要件

- git version 2.25.1 以降 [インストール手順](https://git-scm.com/downloads)
- go version 1.20.0 以降 [インストール手順](https://go.dev/doc/install)

## インストール方法

fit には本体のほかに、オプションで以下の便利機能をインストールすることができます。
- fit コマンド補完機能
- gitステータスバー表示ツール

### fit本体のインストール

最新のfitをインストールする。
```
go install github.com/kazurego7/fit/cmd/fit@latest
```

### fit コマンド補完機能のインストール
利用するシェルに合わせて、以下のインストール方法から選択してください。

PowerShell:
```powershell
fit setting completion --powershell > "$(Split-Path -Path $PROFILE)/fit_completion.ps1"
echo '. "$(Split-Path -Path $PROFILE)/fit_completion.ps1"' >> $PROFILE
# この機能を有効にするためには、シェルを再起動する必要があります。
```

Bash: Linux
```bash
fit setting completion --bash > /etc/bash_completion.d/fit_completion
```

Bash: macOS
```bash
fit setting completion --bash > $(brew --prefix)/etc/bash_completion.d/fit_completion
```

Zsh:
```bash
fit setting completion --zsh > "${fpath[1]}/_fit_completion"
# この機能を有効にするためには、シェルを再起動する必要があります。
```

fish:
```bash
fit setting completion --fish > ~/.config/fish/completions/fit_completion.fish
```

#### コマンド補完機能の動作確認

- 操作： "fit b"まで入力後、TABキーを1回押す
- 確認： "fit branch"と補完されることを確認する
```bash
fit b
↓
fit branch
```

### gitステータスバー表示ツールのインストール
利用するシェルに合わせて、以下のインストール方法から選択してください。

PowerShell:  
[posh-git](https://github.com/dahlbyk/posh-git)のインストール方法を以下に記載する
**Powershellを管理者として実行する**
```powershell
Install-Module posh-git -Scope CurrentUser -Force
Add-PoshGitToProfile
```

Bash:  
[posh-git-bash](https://github.com/lyze/posh-git-sh)のインストール方法を以下に記載する
```bash
wget https://raw.githubusercontent.com/lyze/posh-git-sh/master/git-prompt.sh ~/.git-prompt.sh
cat << EOF >> ~/.bashrc
source ~/.git-prompt.sh
export PROMPT_COMMAND='__posh_git_ps1 "\\[\[\e[0;32m\]\u@\h \[\e[0;33m\]\w" " \[\e[1;34m\]\n\$\[\e[0m\] ";'$PROMPT_COMMAND
EOF
# この機能を有効にするためには、シェルを再起動する必要があります。
```

#### gitステータスバー表示ツールの動作確認

- 操作： 適当なローカルリポジトリをカレントディレクトリとする
- 確認： gitステータスバーが表示されていることを確認する
![gitステータスバーの表示画像](./doc/image/readme/git-prompt.png)

