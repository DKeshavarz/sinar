package server

import (
	"github.com/gin-gonic/gin"
)

func New()*gin.Engine {
	r := gin.Default()
	setup(r)
	return r
}

func setup(r *gin.Engine) {
	//TODO
}