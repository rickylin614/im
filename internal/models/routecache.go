package models

type RouteCache struct {
	ID string `gorm:"column:id"`
}

func (*RouteCache) TableName() string {
	return "route_cache"
}
