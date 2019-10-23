package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

//   -----------------------------------------------------------------------------
//    _____                            _     _    _                 _ _
//   |  __ \                          | |   | |  | |               | | |
//   | |__) |___  __ _ _   _  ___  ___| |_  | |__| | __ _ _ __   __| | | ___ _ __
//   |  _  // _ \/ _` | | | |/ _ \/ __| __| |  __  |/ _` | '_ \ / _` | |/ _ \ '__|
//   | | \ \  __/ (_| | |_| |  __/\__ \ |_  | |  | | (_| | | | | (_| | |  __/ |
//   |_|  \_\___|\__, |\__,_|\___||___/\__| |_|  |_|\__,_|_| |_|\__,_|_|\___|_|
//                  | |
//                  |_|
func requestHandler(ctx *fasthttp.RequestCtx) {
	defer func() { r := recover(); if r != nil { log.Print("Error detected:", r) } }()
	switch string(ctx.Path()){
	case "/test":
		fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
		fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
		fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
		fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
		fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
		fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
		fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
		fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
		fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
		fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())
		response, _ := json.MarshalIndent(conf,"","    ")
		fmt.Fprintf(ctx,"Configurations\n [%v]",response)
	case "/login":
		response := login()
		fmt.Fprintf(ctx,"Login Response [%v]",response)
	case "/register":
		fmt.Fprintf(ctx,"Register response [%v]",register())
	case "/upload":
		fmt.Fprintf(ctx,"Last Transaction")
	default:
		ctx.Error("Not a valid command", fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx,"\nstatus -> returns the accumilations of the engine")
		fmt.Fprintf(ctx,"\nconnections -> returns the database connection details")
	}

}

func login() (response bool) {

	response = true
	return
}

func register() (response bool){

	response=true
	return
}

func upload() (response bool){

	response = true
	return
}
