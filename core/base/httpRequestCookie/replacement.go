package httpRequestCookie

import (
	"go-agent/model/request"
	"go-agent/utils"
	"net/http"
)

func Cookie(req *http.Request, name string) (*http.Cookie, error) {
	cookie, err := CookieT(req, name)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(name),
		Reqs:            utils.Collect(cookie, err),
		Source:          true,
		OriginClassName: "http.(*Request)",
		MethodName:      "Cookie",
		ClassName:       "http.(*Request)",
	})
	return cookie, err
}

func CookieT(req *http.Request, name string) (c *http.Cookie, e error) {
	return c, e
}