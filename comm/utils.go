package comm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	"net"
	"os"
)

func MD5Encrypt(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func GetMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func GetLocalIP() (ip string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ip
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return
}

func SetConfig(
		tsn string, mode string, proxy string,
		tone_proxy string, testlib_proxy string,
	) error{
	var configViper = viper.New()
	configViper.AddConfigPath(".")
	configViper.AddConfigPath(beego.AppConfig.String("AgentConfigFilePath"))
	configViper.SetConfigName(beego.AppConfig.String("AgentConfigFileName"))
	configViper.SetConfigType("yaml")
	if tsn == ""{
		tsn = viper.GetString("tsn")
	}
	if mode == ""{
		mode = viper.GetString("mode")
	}
	if proxy == ""{
		proxy = viper.GetString("proxy")
	}
	if tone_proxy == ""{
		tone_proxy = viper.GetString("tone-proxy")
	}
	if testlib_proxy == ""{
		testlib_proxy = viper.GetString("testlib-proxy")
	}
	configViper.Set("tsn", tsn)
	configViper.Set("mode", mode)
	configViper.Set("proxy", proxy)
	configViper.Set("tone-proxy", tone_proxy)
	configViper.Set("testlib-proxy", testlib_proxy)
	configViper.WriteConfig()
	return nil
}

func GetLog() string {
	logFilePath := beego.AppConfig.String("LogFileDir")
	logFileName := beego.AppConfig.String("LogFileName")
	file, err := os.Open(fmt.Sprintf("%s/%s", logFilePath, logFileName))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	log := string(buffer)
	logLimit := 100000
	if len(log) < logLimit{
		return log
	}else{
		return log[len(log)-logLimit:]
	}
}