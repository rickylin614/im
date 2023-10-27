package resp

type RouteCacheGet struct{}

type RouteCacheGetList struct {
	Page PageResponse    `json:"page,omitempty"`
	Data []RouteCacheGet `json:"data"`
}
