# ToneAgent

[![](https://img.shields.io/badge/made%20by-openanolis-blue.svg?style=flat-square)](https://protocol.ai)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> ToneAgent是T-One运行用例的物理机器代理

ToneAgent实现了多台机器和T-One之间的通讯, 保证T-One中用例的分布式执行, 提高执行效率。

通过ToneAgent, T-One可以向多台物理机/虚拟机分配待执行状态的用例。当用例完成后, T-One会将不同机器的执行结果进行汇总, 并将结果展示在T-One平台上。

## 安装说明

### 前置条件

[Go版本不低于1.15](https://go.dev/dl/)

### 安装

- 克隆当前仓库
```shell script
git clone git@gitee.com:anolis/tone-agent.git
```

- 编译ToneAgent二进制

```shell script
cd tone-agent
go build 
# or
go install ./...
```

> 如果你是用的是unix类的机器,可以直接执行当前目录的脚本`start.sh`

```shell script
chmod +x start.sh
./start.sh
```

## 使用说明

### 配置文件

执行 `ToneAgent` 依赖配置文件 `config.yaml`。

> `ToneAgent` 二进制会默认在 *当前目录 `.`* 和 *用户目录 `$HOME/toneagent`* 查找 配置文件 `config.yaml`

`config.yaml` 样例:

```yaml
beego:  # beego相关配置文件
  AppName: toneagent
  RunMode: dev
  StaticDir: down1
  DirectoryIndex: true
  HttpAddr: 0.0.0.0
  CopyRequestBody: true
  HttpPort: 8479

result: # ToneAgent相关结果文件目录, 默认放在当前目录下
  ResultFileDir: results
  WaitingSyncResultDir: sync_results
  TmpScriptFileDir: scripts
  LogFileDir: logs
  LogFileName: toneagent.log

mode: active # ToneAgent模式
tsn: tone20210101-001 # ToneAgent标识
proxy: https://tone-agent.openanolis.cn # T-One代理地址
```

> 可以通过 `export TONE_AGENT_PATH = $PATH` 来修改 `ToneAgent` 的结果文件地址


### 主动模式

### 被动模式

## 贡献

欢迎加入。
 
创建一个[issue](https://gitee.com/anolis/tone-agent/issues)!

## License

Mulan