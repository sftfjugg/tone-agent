#!/bin/bash

# set -x

DIR_File="./file"

toneagent_file_gz="toneagent.tar.gz"
toneagent_file="toneagent.run"

if [ -d $DIR_File ]; then
    echo $DIR_File "exists"
    rm -rf $DIR_File
fi
# 下载全平台安装包
mkdir $DIR_File
wget -i filelist.txt -P $DIR_File

tar -zcvf $toneagent_file_gz ./file/*

cat install.sh $toneagent_file_gz > $toneagent_file
chmod +x toneagent.run

# 清理临时文件
rm -rf $DIR_File
rm $toneagent_file_gz

if [ -f $toneagent_file ]; then
    echo "$toneagent_file gen success"
fi


