package models

type {{ .FileName }} struct {
}

func (*{{ .FileName }}) TableName() string {
	return "{{ .FileNameSnakeCase }}"
}