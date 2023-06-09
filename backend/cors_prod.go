//go:build prod

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func config(router *gin.Engine) {
	useCors(router)
}

func useCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://lastfm-intersector-render-zuys.onrender.com",
		},
	}))
}
