#!/usr/bin/env bash

set -x

ROOT_DIR=/Users/ccjian/go/src/protobuf-demo


echo "请选择";
workdir=$(cd $(dirname $0); pwd);
# 选定项目
select env in "server"
do
    case $env in
        "server")
            cd $ROOT_DIR/self/deploy
            sudo docker-compose -f server.yaml up -d
            break
            ;;
        *)
            echo "输入错误，请重新输入"
    esac
done


