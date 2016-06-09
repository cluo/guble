package metrics

import (
	log "github.com/Sirupsen/logrus"

	"expvar"
	"fmt"
	"io"
	"net/http"
	"os"
)

var Enabled = len(os.Getenv("GUBLE_METRICS")) > 0

type IntVar interface {
	Add(int64)
	Set(int64)
}

type emptyInt struct{}

// Dummy functions on EmptyInt
func (v *emptyInt) Add(delta int64) {}

func (v *emptyInt) Set(value int64) {}

func NewInt(name string) IntVar {
	if Enabled {
		return expvar.NewInt(name)
	}
	return &emptyInt{}
}

func HttpHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	writeMetrics(rw)
}

func writeMetrics(w io.Writer) {
	fmt.Fprintf(w, "{\n")
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}

func LogOnDebugLevel() {
	if log.GetLevel() == log.DebugLevel {
		fields := log.Fields{}
		expvar.Do(func(kv expvar.KeyValue) {
			fields[kv.Key] = kv.Value
		})
		log.WithFields(fields).Debug("metrics: current values")
	}
}
