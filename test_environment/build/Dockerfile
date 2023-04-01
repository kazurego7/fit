# *************************************************************
# テストリポジトリのビルドステージ
# テストで利用する git リポジトリをスクリプトからビルドする
# *************************************************************
FROM golang:1.20-bullseye AS test_repo_build_stage

# ツールのインストール
RUN apt-get update && apt-get install -y \
    parallel \
    git \
    # クリーンアップ
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# git の初期設定
RUN git config --global user.name "kazurego7" && \
    git config --global user.email "kazurego7@gmail.com"

# テスト用gitリポジトリのビルドスクリプトを配置する
RUN mkdir /usr/src/build_script
RUN mkdir /usr/src/repo
COPY build_script /usr/src/build_script
COPY build_executor.sh /usr/src/

# ビルド実行スクリプトを起動する
WORKDIR /usr/src/repo
RUN ls -1 /usr/src/build_script \
    | parallel -j 100% /usr/src/build_executor.sh /usr/src/build_script/{}

ENTRYPOINT [ "bash" ]


# *************************************************************
# テスト実行ステージ
# golang によるテストの実行を行う
# *************************************************************
FROM golang:1.20-bullseye AS test_run_stage

# ツールのインストール
RUN apt-get update && apt-get install -y \
    git \
    # クリーンアップ
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# git の初期設定
RUN git config --global user.name "kazurego7" && \
    git config --global user.email "kazurego7@gmail.com"
    
# fit プロジェクトをコピー
RUN mkdir /usr/src/fit
COPY ./fit /usr/src/fit
WORKDIR /usr/src/fit

# fit のインストール
RUN go install

# テスト用のリポジトリを配置
COPY --from=fit-test-repo-build /usr/src/repo /usr/src/repo

# テスト実行
ENTRYPOINT ["bash"]