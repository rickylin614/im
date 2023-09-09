package models

type {{ .FileName }} struct {
	ID string `gorm:"column:id"`
}

func (*{{ .FileName }}) TableName() string {
	return "{{ .FileNameSnakeCase }}"
}