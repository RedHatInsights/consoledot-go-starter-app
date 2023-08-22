package metrics

import (
	"net/http"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "starter_app_metric_prefix_requests_total",
		Help: "The total number of processed events",
	})
	errorsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "starter_app_metric_prefix_errors_total",
		Help: "The total number of errors",
	})
)

func Serve(conf *config.Config) {
	http.Handle(conf.AppConfig.MetricsPath, promhttp.Handler())
	http.ListenAndServe(conf.GetMetricsPort(), nil)
}

func IncrementRequests() {
	requestsTotal.Inc()
}

func IncrementErrors() {
	errorsTotal.Inc()
}
