package service

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Requests  *prometheus.CounterVec
	Durations prometheus.Summary
}

func (s *Service) registerMetrics() {
	s.Metrics.Requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: s.Name,
			Name: "requests_total",
			Help: "Number of requests processed by status",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(s.Metrics.Requests)

	s.Metrics.Durations = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: s.Name,
			Name:       "request_durations",
			Help:       "Requests latencies in microseconds",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		})
	prometheus.MustRegister(s.Metrics.Durations)
}