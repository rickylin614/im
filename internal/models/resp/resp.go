package resp

type APIResponse[T any] struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Data    T      `json:"data"`
}

type PageResponse struct {
	Index     int   `gorm:"-" json:"index"` // 頁碼
	Size      int   `gorm:"-" json:"size"`  // 筆數
	TotalPage int   `gorm:"-" json:"pages"` // 總頁數
	Total     int64 `gorm:"-" json:"total"` // 總筆數
}
