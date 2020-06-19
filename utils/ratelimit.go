package utils

import (
	"beego-app/controllers"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"net"
	"strconv"
)

type rateLimiter struct {
	generalLimiter *limiter.Limiter
}

func InitLimiter() (*rateLimiter, error) {
	r := &rateLimiter{}

	// 10 requests per second
	rate, err := limiter.NewRateFromFormatted("10-S")
	if err != nil {
		return nil, errors.New("Create rate instance failed.")
	}
	r.generalLimiter = limiter.New(memory.NewStore(), rate)

	return r, nil
}

func RateLimit(r *rateLimiter, ctx *context.Context) {
	var (
		limiterCtx limiter.Context
		ip         net.IP
		err        error
		req        = ctx.Request
	)

	ip = r.generalLimiter.GetIP(req)
	limiterCtx, err = r.generalLimiter.Get(req.Context(), ip.String())
	if err != nil {
		panic(err)
	}

	h := ctx.ResponseWriter.Header()
	h.Add("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
	h.Add("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
	h.Add("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

	if limiterCtx.Reached {
		ctx.Output.SetStatus(429)
		resBytes, err := json.Marshal(controllers.OutResponse(429, "Request too frequent", nil))
		if err != nil {
			fmt.Println(err)
		}
		ctx.Output.Body(resBytes)
		return
	}

}
