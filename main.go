package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
		api.Static("/video", "./videos")
	}
	r.Run()

}

func getAudioURLOrDefault(id string) string {
	vinfo, err := ytdl.GetVideoInfoFromID(id)
	check(err)
	best := vinfo.Formats.Best(ytdl.FormatAudioEncodingKey) // ytdl.FormatAudioBitrateKey
	if len(best) > 0 {
		log.Println(best[0])
		check(err)
		log.Println("AUDIO BITRATE")
		return saveFile(vinfo, best[0])
	}
	url, err := vinfo.GetDownloadURL(vinfo.Formats[0])
	check(err)
	log.Println("DEFAULT FORMAT")
	return url.String()

}

var VideoPath string = "./videos/"

func saveFile(v *ytdl.VideoInfo, format ytdl.Format) string {
	filename := strings.Replace(v.Title+".mp4", " ", "", -1)
	filepath := VideoPath + filename
	var _, err = os.Stat(filepath)
	if os.IsNotExist(err) {

		var file, _ = os.Create(filepath)
		defer file.Close()
		v.Download(format, file)
	}
	return filename
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

type test struct {
	a string
	b string
}
