package tracing

import (
	"fmt"
	"net/http"
	"strconv"

	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"github.com/openzipkin/zipkin-go/model"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const endpointURL = "http://localhost:9411/api/v2/spans"

type Middler struct {
	Tracer      *zipkin.Tracer
	EndpointURL string
	Port        int
}

func (m *Middler) NewTracer(service string) (*zipkin.Tracer, error) {
	// reporter is responsible for sending traces to zipkin server
	zipEndpoint := fmt.Sprintf("%s:%s", m.EndpointURL, strconv.Itoa(m.Port))
	reporter := reporterhttp.NewReporter(zipEndpoint)

	// local service endpoint
	localEndpoint := &model.Endpoint{ServiceName: service, Port: uint16(m.Port)}

	// Which traces to be sampled. In this case 100% (ie 1.00) of traces will be recorded.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}

	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, err
	}

	return t, err
}

func (m *Middler) Instrument(span string, next http.Handler) http.Handler {
	return zipkinhttp.NewServerMiddleware(
		m.Tracer,
		zipkinhttp.SpanName(span),
	)(next)
}

func (m *Middler) InstrumentHTTPClient(client *http.Client) error {
	var err error

	client.Transport, err = zipkinhttp.NewTransport(
		m.Tracer,
		zipkinhttp.TransportTrace(true),
	)
	if err != nil {
		return err
	}
	return nil
}

// Zipkin middleware for RPC to be implemented in near future (23/7/2018)
