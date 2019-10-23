package main

import (
	"github.com/valyala/fasthttp"
)

func init(){
	myLog("Loading",conf.VERSION)
}

func main()  {
	handler := requestHandler
	if conf.COMPRESS {
		fasthttp.CompressHandler(handler)
	} else {
		checkErr(fasthttp.ListenAndServe(":"+conf.PORT, handler))
	}
}
