package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	ccacheexporter "github.com/virtualtam/ccache_exporter"
)

const (
	DefaultListenAddr = ":9508"
)

func main() {
	listenAddr := flag.String("listenAddr", DefaultListenAddr, "Listen on this address")
	flag.Parse()

	ccacheCollector := ccacheexporter.NewCcacheCollector()
	prometheus.MustRegister(ccacheCollector)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(
			`<html>
             <head><title>ccache exporter</title></head>
             <body>
             <h1>ccache exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	log.Println("Listening on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
