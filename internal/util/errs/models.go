package errs

import (
	"im/internal/util/stringutil"
	"strconv"
	"strings"
)

type ErrorManager struct {
	groupError []*GroupError
}

func NewErrorManager() *ErrorManager {
	return &ErrorManager{
		groupError: make([]*GroupError, 0),
	}
}

func (d *ErrorManager) Group(group string) *GroupError {
	g := &GroupError{Name: group}
	d.groupError = append(d.groupError, g)
	return g
}

type GroupError struct {
	Name  string
	count int
	m     []ErrorCode
}

// Add 新增錯誤
func (g *GroupError) Add(msg string, stautsCode int) *ErrorCode {
	g.count++
	code := IntToStringThreeChar(g.count)
	e := ErrorCode{
		code:    g.Name + "-" + code,
		message: msg,
	}
	g.m = append(g.m, e)
	return &e
}

// IntToStringThreeChar like fmt.Sprintf(%03d,i) , quick response string
func IntToStringThreeChar(i int) string {
	str := strconv.Itoa(i)
	diff := 3 - len(str)
	if diff > 0 {
		var builder strings.Builder
		builder.Grow(3)
		for j := 0; j < diff; j++ {
			builder.WriteByte('0')
		}
		builder.WriteString(str)
		return builder.String()
	}
	return str[:3]
}

func (g *GroupError) ListCodes() []string {
	result := make([]string, 0)
	for _, v := range g.m {
		result = append(result, v.code)
	}
	return result
}

func (g *GroupError) ListCodeNMsg() []string {
	result := make([]string, 0)
	for _, v := range g.m {
		result = append(result, stringutil.Join(v.code, ":", v.message))
	}
	return result
}
