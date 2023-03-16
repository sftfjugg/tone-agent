package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
	"tone-agent/entity"
)

func MainController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mode":  viper.Get("mode"),
		"tsn":   viper.Get("tsn"),
		"proxy": viper.Get("proxy"),
	})
}

func SetAgentConfig(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]string)
	json.Unmarshal(body, &data)
	mode := strings.Trim(strings.Trim(data["mode"], " "), "/")
	tsn := strings.Trim(strings.Trim(data["tsn"], " "), "/")
	proxy := strings.Trim(strings.Trim(data["proxy"], " "), "/")

	var configViper = viper.New()
	configViper.AddConfigPath(entity.ConfigDir)
	configViper.SetConfigName(entity.APPName)
	configViper.SetConfigType("yaml")
	configViper.Set("tsn", tsn)
	configViper.Set("mode", mode)
	configViper.Set("proxy", proxy)
	err := configViper.WriteConfig()
	msg := entity.SuccessOk
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": gin.H{
			"mode":  viper.Get("mode"),
			"tsn":   viper.Get("tsn"),
			"proxy": viper.Get("proxy"),
		},
	})
}
