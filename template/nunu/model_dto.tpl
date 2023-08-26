package dto

type {{ .FileName }} struct {
}

type Query{{ .FileName }}Cond struct {
    PageCond
}

type Create{{ .FileName }}Cond struct {
}

type Update{{ .FileName }}Cond struct {
}

type Delete{{ .FileName }}Cond struct {
}

type {{ .FileName }}Response struct {
	{{ .FileName }} []*{{ .FileName }}   `json:"{{ .FileNameTitleLower }}"`
	Page   *PageResult `json:"page"`
}
