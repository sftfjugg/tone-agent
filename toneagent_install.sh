#!/bin/bash

wget https://anolis-service-pub.oss-cn-zhangjiakou.aliyuncs.com/biz-resource/tone/rpms/toneagent-1.1.0-1.an8.x86_64.rpm

rpm -e toneagent

yum -y localinstall toneagent-1.1.0-1.an8.x86_64.rpm


ip=`curl http://ifconfig.io`

echo "↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓"
echo "配置方法：浏览器打开 ->  http://$ip:8479"
echo "↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓"
