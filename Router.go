package main

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
	"time"
)


var (
	filesHandler = fasthttp.FSHandler("/home/daniel/repo/Metis-Link/files", 0)
	activeKeys = map[string]string{}
	cookieKey = "Authenticate"
	storageUrl = os.Getenv("STORAGE")
)

func requestHandler(ctx *fasthttp.RequestCtx) {

	if ctx.IsGet(){
		filesHandler(ctx)
	}
	if ctx.IsPost(){
		link := string(ctx.FormValue("link"))
		if len(link) < 5{
			ctx.Redirect("badboi", 400)
			return
		}
		entry := CreateShare("link", time.Now(), ctx.RemoteAddr().String(), link)
		json, err := entry.toJson()
		if err != nil{
			ctx.Redirect("badboi", 500)
			return
		}

		res, err := http.Post(storageUrl, "application/json", bytes.NewBuffer(json))
		if err != nil{
			ctx.Redirect("badboi", 500)
			return
		}

		if res.StatusCode < 300 && res.StatusCode > 199{
			ctx.Redirect("goodboi", 201)
		}
	}

}

