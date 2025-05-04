#!/bin/bash

# ベースディレクトリを指定
BASE_DIR="scripts"

# findコマンドで .go ファイルを検索
find "$BASE_DIR" -maxdepth 1 -type f -name '*.go' -not -name '.*' | while read -r FILE; do
    # ディレクトリとファイル名を分割
    DIR=$(dirname "$FILE")
    BASENAME=$(basename "$FILE" .go)

    # 新しいディレクトリのパスを作成
    NEW_DIR="$DIR/$BASENAME"

    # ディレクトリが存在しない場合は作成
    if [ ! -d "$NEW_DIR" ]; then
        mkdir -p "$NEW_DIR"
        echo "ディレクトリ $NEW_DIR を作成しました。"
    fi

    # ファイルを新しいディレクトリに移動
    mv "$FILE" "$NEW_DIR/$BASENAME.go"
    echo "ファイル $FILE を $NEW_DIR/$BASENAME.go に移動しました。"
done
