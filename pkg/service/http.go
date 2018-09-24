package service

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	serviceName = "example"
)

var (
	address string
	path string
)

func Run(_address, _path string) error {
	address = _address
	path = _path

	a := strings.Split(address, ":")
	port, err := strconv.Atoi(a[1])
	if err != nil {
		return err
	}

	s, err := NewService(serviceName, port)
	if err != nil {
		return err
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.Handle(path, promhttp.Handler())
	http.Handle("/", prometheus.InstrumentHandler(serviceName, s))
	fmt.Printf("get the service metrics on %s%s\n", address, path)
	return http.ListenAndServe(address, nil)
}