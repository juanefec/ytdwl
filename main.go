package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rylio/ytdl"
	"net/http"
)

func main() {
	fmt.Println("Starting app...")
	r := gin.Default()
	r.Static("/index", "./public")
	api := r.Group("/api")
	{
		api.GET("/getURL/:ID", 
func(c *gin.Context) {
			id := c.Query("ID")
			url := getAudioURL(id)
			c.String(http.StatusOK,
				url)
		})
	}
	r.Run()

}

func getAudioURL(id string) string {
	vinfo, err :=

		ytdl.GetVideoInfoFromID(id)
	if err != nil {
		panic(err)
	}
	url, err :=

		vinfo.GetDownloadURL(vinfo.Formats.Best(ytdl.FormatAudioEncodingKey)[0])
	if err != nil {
		panic(err)
	}
	return url.String()
}

type test struct {
	a string
	b string
}
