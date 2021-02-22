package caddymin

import (
	"net/http"
	"net/http/httptest"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(Mini{})
	httpcaddyfile.RegisterHandlerDirective("mini", parseCaddyfileHandlerDirective)
}

// Mini _
type Mini struct {
	logger *zap.Logger
}

func (m Mini) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	rr := httptest.NewRecorder()

	next.ServeHTTP(rr, r)

	// m.logger.Error(rr.Header().Get("Content-Type"))
	if rr.Header().Get("Content-Type") == "text/css; charset=utf-8" {
		mi := minify.New()
		mi.AddFunc("text/css", css.Minify)
		b, _ := mi.Bytes("text/css", rr.Body.Bytes())
		w.Header().Add("Content-Type", "text/css; charset=utf-8")
		w.Write(b)
		return nil
	}

	w.Write(rr.Body.Bytes())

	return nil
}

// CaddyModule returns the Caddy module information.
func (Mini) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.mini",
		New: func() caddy.Module { return new(Mini) },
	}
}

// Provision mini
func (m *Mini) Provision(ctx caddy.Context) error {
	m.logger = ctx.Logger(m)

	return nil
}

func parseCaddyfileHandlerDirective(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := &Mini{}

	return m, nil
}
