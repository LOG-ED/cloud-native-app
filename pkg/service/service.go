package service

import (
	"net/http"
	"strings"
	"html/template"

	prometheus "github.com/prometheus/client_golang/prometheus"
	consul "github.com/hashicorp/consul/api"
)

type Service struct {
	Name        	string
	Port			int
	ConsulAgent		*consul.Agent
	ConsulHealth	*consul.Health
	Metrics     	Metrics
}

func NewService(name string, port int) (*Service, error) {
	s := new(Service)
	s.Name = name
	s.Port = port

	s.registerMetrics()
	s.registerConsul()

	return s, nil
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		us := v * 1000000 // make microseconds
		s.Metrics.Durations.Observe(us)
	}))
	defer timer.ObserveDuration()
	
	key := strings.Trim(r.URL.Path, "/")
	if key != "" {
		s.Metrics.Requests.WithLabelValues("fail").Inc()
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	num := s.countConsulServices()

	t := template.Must(template.ParseFiles("tmpl/homepage.html"))
	t.Execute(w, num)

	s.Metrics.Requests.WithLabelValues("success").Inc()
}