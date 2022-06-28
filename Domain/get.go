package domainGet

import (
	"encoding/json"
	commonGet "github.com/duskNNNN/autodns/Common"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/viper"
)

// 获取域名列表
func ListGet() []string {
	base_url := "https://dnsapi.cn/Domain.List"
	payload := url.Values{
		"type":        {"all"},
		"login_token": {viper.GetString("dnspod.id") + "," + viper.GetString("dnspod.token")},
		"format":      {"json"},
		"lang":        {"cn"},
	}
	response, _ := http.PostForm(base_url, payload)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr.Error())
		os.Exit(-1)
	}
	var list commonGet.DomainList
	json.Unmarshal(body, &list)
	var domainList []string
	for _, v := range list.Domains {
		domainList = append(domainList, v.Name)
	}
	return domainList
}
