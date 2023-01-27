#!/usr/bin/env bash
# You should access .env from makefile instead of here
source .env

readonly TARGET_DIR=src
readonly PLI_KEY_PAIR=$KEY_PAIR_NAME.pem
readonly PUB_KEY_PAIR=$KEY_PAIR_NAME.pub

if [ ! -d ${TARGET_DIR} ]; then
  echo "srcディレクトリが作成されていませんでした。"
  echo "srcディレクトリを作成します。"
  mkdir ${TARGET_DIR}
fi

cd src

if [ -f "${KEY_PAIR_NAME}" ]; then
  echo "${KEY_PAIR_NAME}が既に存在していました。"
  echo "${KEY_PAIR_NAME}を${PLI_KEY_PAIR}に変更します。"
  mv "$KEY_PAIR_NAME" "$KEY_PAIR_NAME".pem
fi

if [ -f "${PLI_KEY_PAIR}" ]; then
  echo "既にプライベートのキーペアが作成されています。"
  exit
fi

if [ -f "${PUB_KEY_PAIR}" ]; then
  echo "既にパブリックのキーペアが作成されています。"
  exit
fi

ssh-keygen -t rsa -b 2048 -f "$KEY_PAIR_NAME"
mv "$KEY_PAIR_NAME" "$KEY_PAIR_NAME".pem
exit
