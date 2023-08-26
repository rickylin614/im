package po

type {{ .FileName }} struct {
}

func (*{{ .FileName }}) TableName() string {
	return "{{ .FileNameSnakeCase }}"
}

type Query{{ .FileName }}Cond struct {
	Page
}
