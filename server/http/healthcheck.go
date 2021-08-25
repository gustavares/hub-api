package http

import (
	"github.com/valyala/fasthttp"
)

func HealthcheckHandler() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)

		ctx.SetBodyString(`{"message": "OK"}`)
	})
}

func ReadinessHandler() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)

		ctx.SetBodyString(`{"message": "OK"}`)
	})
}
