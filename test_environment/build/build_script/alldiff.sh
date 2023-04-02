#!/bin/bash 

# Description
#
# git status --short におけるワークツリーとインデックスの状態、以下の3パターンを検証する
# _X
# X_
# XX
#
# git status は、以下のパターン
# _ 変更されていない
# M 修正済み
# T ファイルの種類が変更された (通常ファイル、シンボリックリンク、サブモジュール)
# A 追加
# D 削除(_D, D_ の2パターンのみ)
# R リネーム(_R, R_ の2パターンのみ)
# U 更新されたが未合体(UU の1パターンのみ)
# ? 追跡されていない(?? の1パターンのみ)
#
# -------------------------------------------------
# X          Y     Meaning
# -------------------------------------------------
#          [AMD]   not updated
# M        [ MTD]  updated in index
# T        [ MTD]  type changed in index
# A        [ MTD]  added to index
# D                deleted from index
# R        [ MTD]  renamed in index
# C        [ MTD]  copied in index
# [MTARC]          index and work tree matches
# [ MTARC]    M    work tree changed since index
# [ MTARC]    T    type changed in work tree since index
# [ MTARC]    D    deleted in work tree
#             R    renamed in work tree
#             C    copied in work tree
# -------------------------------------------------
# D           D    unmerged, both deleted
# A           U    unmerged, added by us
# U           D    unmerged, deleted by them
# U           A    unmerged, added by them
# D           U    unmerged, deleted by us
# A           A    unmerged, both added
# U           U    unmerged, both modified
# -------------------------------------------------
# ?           ?    untracked
# !           !    ignored
# -------------------------------------------------


set -eux # エラーおよび未定義変数があれば停止

## Main
git init

## 事前準備
echo a > _M.txt
echo a > M_.txt
echo a > MM.txt
echo a > _T.txt
echo a > T_.txt
echo a > TT.txt
echo a > _A.txt
echo a > A_.txt
echo a > AM.txt
echo xxx > _D.txt
echo xxx > D_.txt
echo r1 > _R.txt
echo r2 > R_.txt
echo r3 > RM.txt
echo a > UU.txt
# echo a > XX.txt

git add _M.txt
git add M_.txt
git add MM.txt
git add _T.txt
git add T_.txt
git add TT.txt
git add _A.txt
git add A_.txt
git add _D.txt
git add D_.txt
git add _R.txt
git add R_.txt
git add RM.txt
git add UU.txt
# git add XX.txt

git commit --message "準備"

rm _A.txt
rm A_.txt
rm AM.txt

git commit -a --message "準備"

## マージ準備
echo x > UU.txt
git stash push --message "コンフリクト"

## インデックス設定

echo b > M_.txt
echo b > MM.txt

rm T_.txt && ln -s ln.txt T_.txt
rm TT.txt && ln -s ln.txt TT.txt

echo b > A_.txt
echo b > AM.txt

git rm D_.txt

mv R_.txt R_2.txt
mv RM.txt RM2.txt

git add --all

echo c > UU.txt
git commit UU.txt --message "conflict"
git stash apply --quiet || true

## ワークツリー設定

echo c > _M.txt
echo c > MM.txt

rm _T.txt && ln -s ln2.txt _T.txt
rm TT.txt && echo c > TT.txt

echo c > _A.txt && git add --intent-to-add _A.txt
echo c > AM.txt

rm _D.txt

mv _R.txt _R2.txt && git add --intent-to-add _R2.txt
echo c > RM2.txt

echo b > XX.txt

## success
echo "success"