result_dir=$HOME/ToneAgent/results
sync_dir=$HOME/ToneAgent/sync_results
scripts_dir=$HOME/ToneAgent/scripts
logs_dir=$HOME/ToneAgent/logs
config_dir=$HOME/ToneAgent

mkdir -p $result_dir $sync_dir $scripts_dir $logs_dir
go mod tidy
go build
cp ./tone-agent $HOME/ToneAgent
cd $HOME/ToneAgent
echo "beego:
  AppName: toneagent
  RunMode: dev
  StaticDir: down1
  DirectoryIndex: true
  HttpAddr: 0.0.0.0
  CopyRequestBody: true
  HttpPort: 8479

result:
  ResultFileDir: results
  WaitingSyncResultDir: sync_results
  TmpScriptFileDir: scripts
  LogFileDir: logs
  LogFileName: toneagent.log


mode: active
tsn: tone20210101-001
proxy: https://tone-agent.openanolis.cn" > config.yaml

echo "ToneAgent start ..."
nohup ./tone-agent >> toneagent.log 2>&1 &
# ./tone-agent