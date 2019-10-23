package main

import (
	"github.com/valyala/fasthttp"
)

func init(){
	myLog("Loading",conf.VERSION)
}

func main()  {
	handler := requestHandler
	checkErr(fasthttp.ListenAndServe(":"+conf.PORT, handler))
}
