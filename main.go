package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rylio/ytdl"
)

func main() {
	fmt.Println("Starting app...")
	r := gin.Default()
	r.Static("/index", "./public")
	api := r.Group("/api")
	{
		api.GET("/getURL",
			func(c *gin.Context) {
				params := c.Request.URL.Query()
				log.Println(params)
				id := params.Get("id")
				log.Println("id:", id)
				url := getAudioURLOrDefault(id)
				c.String(http.StatusOK, url)
			})
	}
	r.Run()

}

func getAudioURLOrDefault(id string) string {
	vinfo, err := ytdl.GetVideoInfoFromID(id)
	log.Println("vinfo: ", vinfo)
	check(err)
	best := vinfo.Formats.Best(ytdl.FormatAudioBitrateKey) // ytdl.FormatAudioBitrateKey
	if len(best) > 0 {
		log.Println(best)
		url, err := vinfo.GetDownloadURL(best[0])
		check(err)
		log.Println("AUDIO BITRATE")
		return url.String()
	}
	url, err := vinfo.GetDownloadURL(vinfo.Formats[0])
	check(err)
	log.Println("DEFAULT FORMAT")
	return url.String()

}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

type test struct {
	a string
	b string
}
