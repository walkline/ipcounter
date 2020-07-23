package metrics

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/walkline/ipcounter"
)

type Server struct {
	http.Server
}

func NewServer(index ipcounter.IPIndex, logger log.Logger) *Server {
	prometheus.MustRegister(
		prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "unique_ip_addresses",
				Help: "Counts the unique IP addresses from logs endpoint",
			},
			func() float64 {
				l, err := index.Len()
				if err != nil {
					logger.Log("err", "can't get len of index for prometheus counter, err: ", err.Error())
				}

				return float64(l)
			},
		),
	)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return &Server{
		Server: http.Server{
			Handler: mux,
		},
	}
}

func (s *Server) ListenAndServe(addr string) error {
	s.Server.Addr = addr
	return s.Server.ListenAndServe()
}
