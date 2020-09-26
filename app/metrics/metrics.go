package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RequestsToIndex = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "requests_to_root",
		Help: "Total number of requests to root",
	},
)
