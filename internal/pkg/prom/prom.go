package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

func NewPromManager() *Manager {
	registry := prometheus.DefaultRegisterer
	onlineGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "online_user_num",
		Help: "The number of online user num",
	})
	registry.MustRegister(onlineGauge)
	return &Manager{
		OnlineUserGauge: onlineGauge,
		Registry:        prometheus.DefaultRegisterer,
		Gather:          prometheus.DefaultGatherer,
	}
}

type Manager struct {
	OnlineUserGauge prometheus.Gauge
	Registry        prometheus.Registerer
	Gather          prometheus.Gatherer
}
