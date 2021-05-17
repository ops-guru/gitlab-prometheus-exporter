package http

import (
	"log"
	"net/http"

	"github.com/andreip-og/gitlab-exporter/exporter"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	Handler  http.Handler
	exporter exporter.Exporter
}

func NewServer(exporter exporter.Exporter) *Server {
	r := http.NewServeMux()

	// Register Metrics from each of the endpoints
	// This invokes the Collect method through the prometheus client libraries.
	prometheus.MustRegister(&exporter)

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
		                <head><title>Gitlab Exporter</title></head>
		                <body>
		                   <h1>Gitlab Prometheus Metrics Exporter</h1>
		                   <p><a href='` + "/metrics" + `'>Metrics</a></p>
		                   </body>
		                </html>
		              `))
	})

	return &Server{Handler: r, exporter: exporter}
}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(":8081", s.Handler))
}
