package main

import "os"
import "log"
import "fmt"
import "path"
import "flag"
import "net/http"
import "crypto/sha1"
import "github.com/gin-gonic/gin"

var Service = ":10011"
var DataPath = "/tmp/incs"

func Init() {
	flag.StringVar(&Service, "service", ":10011", "the service address")
	flag.StringVar(&DataPath, "data-path", "/tmp/incs", "the path to storage chunks")
	flag.Parse()
	if _, err := os.Stat(DataPath); os.IsNotExist(err) {
		err = os.Mkdir(DataPath, 0755)
		if nil != err {
			log.Fatalln(err)
		}
	}
}

func Sha1(name string) string {
	sha := sha1.New()
	sha.Write([]byte(name))
	bs := sha.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func main() {
	Init()
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.PUT("/chunks", func(ctx *gin.Context) {
			f, _ := ctx.FormFile("file")
			name := Sha1(ctx.Query("name"))
			err := ctx.SaveUploadedFile(f, path.Join(DataPath, name))
			if nil != err {
				ctx.JSON(http.StatusOK, gin.H{
					"status":  true,
					"message": fmt.Sprintf("fail, %s", err.Error())})
			}
			ctx.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "ok"})
		})
		v1.GET("/chunks", func(ctx *gin.Context) {
			name := Sha1(ctx.Query("name"))
			ctx.File(path.Join(DataPath, name))
		})
	}
	router.Run(Service)
}
