package instrumentation

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	defBuckets = []float64{1, 10, 50, 100, 500, 1000}
)

const (
	reqsName    = "requests_total"
	latencyName = "request_duration_milliseconds"
)

type Middler struct {
	requests *prometheus.CounterVec
	latency  *prometheus.HistogramVec
}

// NewInstrMiddler creates new instrumentation middleware.
// name input refers to the microservice that is instrumented. In this case it is
// the fe-service, but the instrumentation code can be extracted
// as an independent package for use with
func NewInstrMiddler(name string, buckets ...float64) *Middler {
	var m Middler

	m.requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        reqsName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"service": name},
		},
		[]string{"code", "method", "path"},
	)
	prometheus.MustRegister(m.requests)

	if len(buckets) == 0 {
		buckets = defBuckets
	}
	m.latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        latencyName,
		Help:        "How long it took to process the request, partitioned by status code, method and HTTP path.",
		ConstLabels: prometheus.Labels{"service": name},
		Buckets:     buckets,
	},
		[]string{"code", "method", "path"},
	)
	prometheus.MustRegister(m.latency)
	return &m
}

func (m *Middler) Instrument(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		var status string
		status = w.Header().Get("error")
		if status == "" {
			status = http.StatusText(200)
		}

		m.requests.WithLabelValues(status, r.Method, r.URL.Path).Inc()
		m.latency.WithLabelValues(status, r.Method, r.URL.Path).Observe(float64(time.Since(start).Nanoseconds()) / 1000000)
	})
}
