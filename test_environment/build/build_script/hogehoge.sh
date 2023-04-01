#!/bin/bash 
set -eu # エラーおよび未定義変数があれば停止

git init

echo "hogehoge" > hoge.txt
git add -A
git commit -m "hogeeee"

echo "fugafuga" > fuga.txt
git add -A
git commit -m "fugaaaa"