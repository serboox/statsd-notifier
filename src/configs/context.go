package configs

import (
	statsd "github.com/etsy/statsd/examples/go"
	"github.com/serboox/statsd-notifier/src/stats"
)

// Context main application structure
type Context struct {
	Config *Config
	StatsD stats.ClientInterface
}

// NewContext initialize and return new context exemplar
func NewContext(cnf *Config) *Context {
	ctx := new(Context)
	ctx.Config = cnf
	ctx.initStatsDClient()

	return ctx
}

// NewContextMock creates a new instance of the structure
func NewContextMock() *Context {
	ctx := new(Context)
	ctx.Config = NewConfigMock()
	ctx.initStatsDClient()

	return ctx
}

func (ctx *Context) initStatsDClient() {
	if ctx.Config.StatsD.Mocked {
		ctx.StatsD = stats.NewClientMocked()
		return
	}

	client := statsd.New(ctx.Config.StatsD.Host, int(ctx.Config.StatsD.Port))
	ctx.StatsD = stats.NewClient(client)
}
