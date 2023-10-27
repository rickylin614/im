package sql

import (
	"context"
	"embed"
	"fmt"

	"go.uber.org/dig"

	"im/internal/pkg/logger"
)

//go:embed **/*.sql
var embedFiles embed.FS

type fileSystem interface {
	ReadFile(name string) ([]byte, error)
}

type Files struct {
	in         digIn
	embedFiles fileSystem
}

type digIn struct {
	dig.In

	Logger logger.Logger
}

func NewSqlEmbedFile(in digIn) *Files {
	return &Files{in: in, embedFiles: embedFiles}
}

// Load 讀取檔案SQL, 找不到檔案則回傳空值
func (s Files) Load(filename string) string {
	f, err := s.embedFiles.ReadFile(filename)
	if err != nil || len(f) == 0 {
		s.in.Logger.Error(context.Background(), fmt.Errorf("sql file:%s not found", filename))
		return ""
	}
	return string(f)
}
