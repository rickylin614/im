package errs

import (
	"errors"
	"strconv"
	"strings"
	"sync"
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
	m     sync.Map
}

func (g *GroupError) Add(code, msg string) ErrorCode {
	e := ErrorCode{
		code:    g.Name + "-" + code,
		message: msg,
	}
	g.count++
	if _, ok := g.m.LoadOrStore(code, e); ok {
		panic(errors.New("group sub error code duplicate definition, code:" + code))
	}
	return e
}

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
	g.m.Range(func(key, value interface{}) bool {
		k := key.(string)
		result = append(result, k)
		return true
	})
	return result
}
