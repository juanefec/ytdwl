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
		format := getBestAudio(vinfo.Formats)
		check(err)
		log.Println("AUDIO BITRATE", format)
		return saveFile(vinfo, format)
	}
	url, err := vinfo.GetDownloadURL(vinfo.Formats[0])
	check(err)
	log.Println("DEFAULT FORMAT")
	return url.String()

}

func getBestAudio(fl ytdl.FormatList) ytdl.Format {
	var best ytdl.Format
	for _, f := range fl {
		if isAcceptedAudio(f) {
			best = f
		}
	}
	return best
}
func isAcceptedAudio(f ytdl.Format) bool {
	if f.Resolution == "" && f.VideoEncoding == "" && (f.Extension == "webm" || f.Extension == "mp4") {
		return true
	}
	return false
}

// VideoPath for writing files
var VideoPath = "./videos/"

func saveFile(v *ytdl.VideoInfo, format ytdl.Format) string {
	filename := fmtTitleToFilename(v.Title + "." + format.Extension)
	filepath := VideoPath + filename
	var _, err = os.Stat(filepath)
	if os.IsNotExist(err) {

		var file, _ = os.Create(filepath)
		defer file.Close()
		v.Download(format, file)
	}
	return filename
}

func fmtTitleToFilename(t string) string {
	replacer := strings.NewReplacer(" ", "", "\"", "")
	return replacer.Replace(t)
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
