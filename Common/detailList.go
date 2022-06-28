package commonGet

// 任务详情
type DetailList struct {
	Status     Status `json:"status"`
	Total      int    `json:"total"`
	Success    int    `json:"success"`
	Fail       int    `json:"fail"`
	Type       string `json:"type"`
	Created_at string `json:"created_at"`
}
