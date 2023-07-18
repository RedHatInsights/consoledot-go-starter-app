package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	route = "/metrics"
	port  = ":9000"
)

var (
	requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "starter_app_requests_total",
		Help: "The total number of processed events",
	})
	errorsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "starter_app_errors_total",
		Help: "The total number of errors",
	})
)

func Serve() {
	http.Handle(route, promhttp.Handler())
	http.ListenAndServe(port, nil)
}

func IncrementRequests() {
	requestsTotal.Inc()
}

func IncrementErrors() {
	errorsTotal.Inc()
}
