package req

import "time"

type RouteCacheGet struct {
	RouteCacheKey string
}

type RouteCacheSet struct {
	RouteCacheKey  string
	RouteCacheData string
	TTL            time.Duration
}
