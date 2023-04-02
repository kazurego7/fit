#!/bin/bash 

# Description
#
#

set -eux # エラーおよび未定義変数があれば停止

## Main
source $(dirname "$0")"/[コピーするファイル名]"
# source $(dirname "$0")"/_create_template.sh"

## Success
echo "success"
