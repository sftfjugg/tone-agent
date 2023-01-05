# Toneagent Client Pack

生成 run 安装文件,适配 deb 和 rpm 的 x86_64 与 arm 的双平台.

## 使用

```bash
chmod +x gen_run.sh
bash gen_run.sh
```

## 文件

```sh
├── filelist.txt
├── gen_run.sh
├── install.sh
├── README.md
└── toneagent.run
```

- filelist 存放全平台的 Toneagent Client 下载 url
- gen_run.sh 用于生成 run 文件
- install.sh run 文件生成时的安装脚本
- README.md
- toneagent.run 生成的 run 文件

## 客户端文件 filelist

目前已有客户端文件如下

- debian-x86_64-1.0.3  https://anolis-service-pub.oss-cn-zhangjiakou.aliyuncs.com/biz-resource/tone/rpms/toneagent-1.0.3-x86_64.deb
- debian-aarch64-1.0.3  https://anolis-service-pub.oss-cn-zhangjiakou.aliyuncs.com/biz-resource/tone/rpms/toneagent-1.0.3-aarch64.deb
- linux-x86_64-1.0.3  https://anolis-service-pub.oss-cn-zhangjiakou.aliyuncs.com/biz-resource/tone/rpms/toneagent-1.0.3-1.an8.x86_64.rpm
- linux-aarch64-1.0.3  https://anolis-service-pub.oss-cn-zhangjiakou.aliyuncs.com/biz-resource/tone/rpms/toneagent-1.0.3-1.an8.aarch64.rpm

## 注意事项

install.sh

- 当需要更换 客户端文件 时,需要在 `install.sh` 更新文件名.
- **请勿删除 install.sh 最后一行的空白行** 这是正确解包 run 文件必要的空白行