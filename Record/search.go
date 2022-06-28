package recordGet

import (
	"encoding/json"
	commonGet "github.com/duskNNNN/autodns/Common"
	"github.com/duskNNNN/autodns/Mail"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// 根据job_id查看任务详情
func JobDetail(job_id string) {
	base_url := "https://dnsapi.cn/Batch.Detail"
	payload := url.Values{
		"job_id":      {job_id},
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
	var detail commonGet.DetailList
	json.Unmarshal(body, &detail)
	if detail.Total == detail.Success {
		MailSend.MailSend("域名记录全部修改完成")
		log.Println("域名记录全部修改完成")
	} else {
		log.Println("域名记录修改失败")
	}
}
