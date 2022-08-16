package comm

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"tone-agent/entity"
)

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
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

func GetConfig() (entity.Config, error){
	var configViper = viper.New()
	configViper.AddConfigPath(".")
	configViper.AddConfigPath("/usr/local/toneagent/conf/toneagent.toneagent.config.yaml")
	configViper.SetConfigName("config")
	configViper.SetConfigType("yaml")
	//读取配置文件内容
	if err := configViper.ReadInConfig(); err != nil {
		log.Println("[GetConfig] get config file error:", err)
	}
	var config entity.Config
	if err := configViper.Unmarshal(&config); err != nil {
		log.Println("[GetConfig] unmarshal config error:", err)
	}
	return config, nil
}

func SetConfig(tsn string, mode string, proxy string) error{
	var configViper = viper.New()
	configViper.AddConfigPath(".")
	configViper.SetConfigName("config")
	configViper.SetConfigType("yaml")
	if tsn != ""{
		configViper.Set("tsn", tsn)
	}
	if mode != ""{
		configViper.Set("mode", mode)
	}
	if proxy != ""{
		configViper.Set("proxy", proxy)
	}
	c := configViper.AllSettings()
	bytess, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	fmt.Println(string(bytess))
	return nil
}

func MD5Encrypt(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
