package resp

type {{ .FileName }}Get struct{}

type {{ .FileName }}GetList struct {
	Page PageResponse `json:"page,omitempty"`
	Data []{{ .FileName }}Get `json:"data"`
}
