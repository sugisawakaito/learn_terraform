#!/usr/bin/env bash
# You should access .env from makefile instead of here
source .env

if [ $1 == "" ]; then
  echo "ファイルを指定してください。"
  exit 1
fi

echo $1
exit