#!/bin/bash
# **********************************************
# テストリポジトリのビルド実行
#
# $1 : テストリポジトリのビルドスクリプトのパス
#
# スクリプトを実行した同階層の `repos` ディレクトリに
# 引数のテスト用リポジトリを作成する
# **********************************************

# コマンドのエラーでビルドエラーとなるように設定
set -eu

# このスクリプトが存在するディレクトリを取得
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
# 生成したテストリポジトリを配置するディレクトリ
REPOS_DIR="${SCRIPT_DIR}/repos"
mkdir -p "${REPOS_DIR}"

# ファイル存在チェック
if [ ! -f $1 ]; then
    echo "File $1 no exists"
    exit 1
fi


# ビルドスクリプトからテストリポジトリ名の決定
FILE_PATH=$(readlink --canonicalize $1)
FILE_NAME=${FILE_PATH##*/}
TEST_NAME=${FILE_NAME%.*}

# ビルドスクリプトが"_"または"."から始まる場合は処理しない
if [[ "${TEST_NAME}" == _* ]]; then
    exit 0
fi

# テストリポジトリのディレクトリ作成
mkdir "$REPOS_DIR/$TEST_NAME"
cd "$REPOS_DIR/$TEST_NAME"

# テストリポジトリのビルドスクリプトの実行
$FILE_PATH
