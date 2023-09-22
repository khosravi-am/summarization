package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/JesusIslam/tldr"
	"github.com/gin-gonic/gin"
)

var text string

func getIndexPage(c *gin.Context) {
	var title string = "SUMMARIZE YOUR ARTICLE!"
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": title,
	})
}

func getUploadPage(c *gin.Context) {
	fmt.Println("geter")
	intoSentences := 3
	res := strings.Split(text, "\n")
	bag := tldr.New()
	result, err := bag.Summarize(text, intoSentences)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	c.HTML(http.StatusOK, "showArticle.html", gin.H{
		"title": res[0],
		"text":  result,
	})

	text = ""
}

func uploadFile(c *gin.Context) {
	fmt.Println("uploadFile")
	file, header, err := c.Request.FormFile("articleFile")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	buf := make([]byte, header.Size)

	for {
		n, err := file.Read(buf)

		if err != nil && err != io.EOF {
			log.Println("Couldn't write file: ")
			break
		}

		if n == 0 {
			break
		}
		text = text + string(buf[:n])
	}

	fmt.Println(text)
	c.JSON(http.StatusOK, gin.H{"filepaths": "filePaths"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", getIndexPage)
	router.GET("/upload", getUploadPage)
	router.POST("/", uploadFile)
	router.Run("localhost:8080")

}
