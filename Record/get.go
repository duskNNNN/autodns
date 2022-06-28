package recordGet

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

// 获取域名记录
func RecordList(domainList []string) []string {
	var id_lists []string
	for i, _ := range domainList {
		base_url := "https://dnsapi.cn/Record.List"
		payload := url.Values{
			"domain":      {domainList[i]},
			"record_type": {"AAAA"},
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
		var recordList commonGet.RecordList
		json.Unmarshal(body, &recordList)
		// 只获取id
		for _, rv := range recordList.Records {
			id_lists = append(id_lists, rv.ID)
		}
	}
	return id_lists
}
