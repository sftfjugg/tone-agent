package core

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/spf13/viper"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var src = rand.NewSource(time.Now().UnixNano())

func GetProxyAPIUrl(api string) string {
	domain := viper.Get("proxy")
	return fmt.Sprintf("%v/%v", domain, api)
}

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func GetHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}

func CheckPid(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Printf("Unable to find the process %d", pid)
		return false
	}

	err = process.Signal(syscall.Signal(0))
	if err != nil {
		log.Printf("Process %d is dead!", pid)
		return false
	} else {
		log.Printf("Process %d is alive!", pid)
		return true
	}
}

func GetCurTimeStr() string {
	timeUnix := time.Now().Unix()
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	return formatTimeStr
}

func GetSign() string {
	tsn := viper.GetString("tsn")
	curTime := time.Now().Unix()
	joinStr := fmt.Sprintf("%s|%d", tsn, curTime)
	sEnc := base64.StdEncoding.EncodeToString([]byte(joinStr))
	return sEnc
}
