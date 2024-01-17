package response

type ExampleGet struct {
	Id          string `json:"id"`          // 數據ID
	Name        string `json:"name"`        // 範例名
	Description string `json:"description"` // 描述
}

type ExampleGetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []ExampleGet `json:"data"`
}
