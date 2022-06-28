package commonGet

type DomainList struct {
	Status  Status   `json:"status"`
	Info    Info     `json:"info"`
	Domains []Domain `json:"domains"`
}

type Status struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Created_at string `json:"created_at"`
}

type Info struct {
	Domain_total    int `json:"domain_total"`
	All_total       int `json:"all_total"`
	Mine_total      int `json:"mine_total"`
	Share_total     int `json:"share_total"`
	Vip_total       int `json:"vip_total"`
	Ismark_total    int `json:"ismark_total"`
	Pause_total     int `json:"pause_total"`
	Error_total     int `json:"error_total"`
	Lock_total      int `json:"lock_total"`
	Spam_total      int `json:"spam_total"`
	Vip_expire      int `json:"vip_expire"`
	Share_out_total int `json:"share_out_total"`
}

type Domain struct {
	ID                int      `json:"id"`
	Status            string   `json:"status" `
	Grade             string   `json:"grade"`
	Group_id          string   `json:"group_id"`
	Searchengine_push string   `json:"searchengine_push"`
	Is_mark           string   `json:"is_mark"`
	TTL               string   `json:"ttl"`
	Cname_speedup     string   `json:"cname_speedup"`
	Remark            string   `json:"remark"`
	Created_on        string   `json:"created_on"`
	Updated_on        string   `json:"updated_on"`
	Punycode          string   `json:"punycode"`
	Ext_status        string   `json:"ext_status"`
	Name              string   `json:"name"`
	Grade_title       string   `json:"grade_title"`
	Grade_ns          []string `json:"grade_ns"`
	Is_vip            string   `json:"is_vip"`
	Owner             string   `json:"owner"`
	Records           string   `json:"records"`
	Vip_start_at      string   `json:"vip_start_at"`
	Vip_end_at        string   `json:"vip_end_at"`
	Vip_auto_renew    string   `json:"vip_auto_renew"`
}
