package recordGet

import (
	"encoding/json"
	"fmt"
	commonGet "github.com/duskNNNN/autodns/Common"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/viper"
)

// 修改记录
func RecordChange(id_lists []string, ipv6 string) string {
	var record_id string
	for _, v := range id_lists {
		record_id += v
		record_id += ","
	}
	fmt.Println(record_id)
	base_url := "https://dnsapi.cn/Batch.Record.Modify"
	payload := url.Values{
		"record_id":   {record_id},
		"change":      {"value"},
		"change_to":   {ipv6},
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
	var job commonGet.JobList
	json.Unmarshal(body, &job)
	return job.Job_id
}
