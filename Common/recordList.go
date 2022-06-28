package commonGet

type RecordList struct {
	Status     Status     `json:"status"`
	Domain     Domain     `json:"domain"`
	RecordInfo RecordInfo `json:"info"`
	Records    []Record   `json:"records"`
}

type RecordInfo struct {
	Sub_domains  string   `json:"sub_domains"`
	Record_total string   `json:"record_total"`
	Records_num  string   `json:"records_num"`
	Records      []Record `json:"records"`
}

type Record struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Line           string `json:"line"`
	Line_id        string `json:"line_id"`
	Type           string `json:"type"`
	TTL            string `json:"ttl"`
	Value          string `json:"value"`
	Weight         string `json:"weight"`
	Mx             string `json:"mx"`
	Enabled        string `json:"enabled"`
	Status         string `json:"status"`
	Monitor_status string `json:"monitor_status"`
	Remark         string `json:"remark"`
	Updated_on     string `json:"updated_on"`
	Use_aqb        string `json:"use_aqb"`
}
