package httpsrvctx

import (
	"strings"

	"github.com/docker/dctx"
)

// WithVars makes vars available on the returned context. Variables are
// available at keys with the prefix "vars.". For example, if looking for the
// variable "name", it can be accessed as "vars.name". Typically, this is used
// when making un-restricted keys in the context, such as with gorilla/mux.
// Only strings keys and values are allowed.
func WithVars(ctx dctx.Context, vars map[string]string) dctx.Context {
	return &muxVarsContext{
		Context: ctx,
		vars:    vars,
	}
}

type muxVarsContext struct {
	dctx.Context
	vars map[string]string
}

func (ctx *muxVarsContext) Value(key interface{}) interface{} {
	if keyStr, ok := key.(string); ok {
		if keyStr == "vars" {
			return ctx.vars
		}

		if strings.HasPrefix(keyStr, "vars.") {
			keyStr = strings.TrimPrefix(keyStr, "vars.")
		}

		if v, ok := ctx.vars[keyStr]; ok {
			return v
		}
	}

	return ctx.Context.Value(key)
}
