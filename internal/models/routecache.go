package models

import (
	"log/slog"
	"time"

	"github.com/goccy/go-json"
)

type RouteCacheGet struct {
	RouteCacheKey string
}

type RouteCacheSet struct {
	RouteCacheKey  string
	RouteCacheData *RouteCache
	TTL            time.Duration
}

func NewRouteCache(status int, body []byte) *RouteCache {
	return &RouteCache{
		Status: status,
		Body:   body,
	}
}

type RouteCache struct {
	Body   []byte
	Status int
}

func (r *RouteCache) Bytes() []byte {
	jsonData, _ := json.Marshal(r)
	return jsonData
}

func (r *RouteCache) String() string {
	jsonData, _ := json.Marshal(r)
	return string(jsonData)
}

func (r *RouteCache) Set(b []byte) *RouteCache {
	err := json.Unmarshal(b, r)
	if err != nil {
		slog.Error("RouteCache Unmarshal", err)
	}
	return r
}
