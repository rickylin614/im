package resp

type APIResponse[T any] struct {
	Code    string `json:"code"` // 回傳代碼
	Message string `json:"msg"`  // 訊息
	Data    T      `json:"data"` // 資料
}

type PageResponse struct {
	Index     int   `gorm:"-" json:"index"` // 頁碼
	Size      int   `gorm:"-" json:"size"`  // 筆數
	TotalPage int   `gorm:"-" json:"pages"` // 總頁數
	Total     int64 `gorm:"-" json:"total"` // 總筆數
}
