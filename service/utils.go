package service

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

const RequestTimeout = 60

func GetHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(RequestTimeout) * time.Second,
	}
	return client
}

func GetCurTimeStr() string {
	timeUnix := time.Now().Unix()
	formatTimeStr := time.Unix(timeUnix, 0).Format("2023-01-01 01:00:00")
	return formatTimeStr
}

func GetSign() string {
	tsn := viper.GetString("tsn")
	curTime := time.Now().Unix()
	joinStr := fmt.Sprintf("%s|%d", tsn, curTime)
	sEnc := base64.StdEncoding.EncodeToString([]byte(joinStr))
	return sEnc
}
