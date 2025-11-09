package statics

import (
	"os"

	"github.com/gin-gonic/gin"
)

const (

)

func Register(group *gin.RouterGroup) {
	os.MkdirAll("./uploads", 0755)
	group.Static("/uploads", "./uploads")
}