package zipkin

import (
	"time"

	"go-kartos-study/pkg/conf/env"
	"go-kartos-study/pkg/net/trace"
	xtime "go-kartos-study/pkg/time"
)

// Config config.
// url should be the endpoint to send the spans to, e.g.
// http://localhost:9411/api/v2/spans
type Config struct {
	Endpoint      string         `dsn:"endpoint"`
	BatchSize     int            `dsn:"query.batch_size,100"`
	Timeout       xtime.Duration `dsn:"query.timeout,200ms"`
	DisableSample bool           `dsn:"query.disable_sample"`
}

// Init init trace report.
func Init(c *Config) {
	if c.BatchSize == 0 {
		c.BatchSize = 100
	}
	if c.Timeout == 0 {
		c.Timeout = xtime.Duration(200 * time.Millisecond)
	}
	trace.SetGlobalTracer(trace.NewTracer(env.AppID, newReport(c), c.DisableSample))
}
