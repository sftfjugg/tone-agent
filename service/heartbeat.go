package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"tone-agent/entity"
)

func SyncHeartbeat() error {
	url := fmt.Sprintf("%s/%s", viper.Get("proxy"), entity.HeartbeatAPI)
	tsn := viper.GetString("tsn")
	sign := GetSign()
	data := map[string]string{
		"tsn":    tsn,
		"sign":   sign,
		"arch":   ExecCommand("arch"),
		"kernel": ExecCommand("uname -r"),
		//"distro": ExecCommand("cat /etc/os-release | grep -i id="),
	}
	bytesData, _ := json.Marshal(data)
	client := GetHttpClient()
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bytesData))
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Heartbeat requests error:%s", err.Error())
		return err
	}
	//results, _ := ioutil.ReadAll(resp.Body)
	//log.Printf("Heartbeat request code:%d | res:%s", resp.StatusCode, results)
	return nil
}
