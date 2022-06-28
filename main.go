package main

import (
	"github.com/duskNNNN/autodns/Conf"
	"github.com/duskNNNN/autodns/Domain"
	"github.com/duskNNNN/autodns/IP"
	"github.com/duskNNNN/autodns/Record"
	"time"
)

func main() {
	// 获取ipv6地址
	ipv6 := ipGet.GetLocalIPv6()
	// // 获取配置文件中的腾讯云token
	confGet.ConfigInit()
	// // 收集获取域名列表
	domainList := domainGet.ListGet()
	// // 获取域名中的记录
	record_ids := recordGet.RecordList(domainList)
	// // 修改记录
	job_id := recordGet.RecordChange(record_ids, ipv6)
	// // 休眠10秒
	time.Sleep(10 * time.Second)
	// // 查询是否修改成功
	recordGet.JobDetail(job_id)
}
