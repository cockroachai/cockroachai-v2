package utils

import "net/http"

// HeaderModify modifies the header of the request.
func HeaderModify(headers *http.Header) {
	// 移除一些错误的转发头
	headers.Del("X-Forwarded-For")
	headers.Del("X-Forwarded-Host")
	headers.Del("X-Forwarded-Proto")
	headers.Del("X-Forwarded-Server")
	headers.Del("X-Real-Ip")
	headers.Del("X-Forwarded-Port")
	headers.Del("X-Forwarded-Uri")
	headers.Del("X-Forwarded-Path")
	headers.Del("X-Forwarded-Method")
	headers.Del("X-Forwarded-Protocol")
	headers.Del("X-Forwarded-Scheme")

	// 移除一些CF的头
	headers.Del("Cf-Connecting-Ip")
	headers.Del("Cf-Ipcountry")
	headers.Del("Cf-Ray")
	headers.Del("Cf-Visitor")
	headers.Del("Cf-Request-Id")
	headers.Del("Cf-Worker")
	headers.Del("Cf-Access-Client-Id")
	headers.Del("Cf-Access-Client-Device-Type")
	headers.Del("Cf-Access-Client-Device-Model")
	headers.Del("Cf-Access-Client-Device-Name")
	headers.Del("Cf-Access-Client-Device-Brand")

}
