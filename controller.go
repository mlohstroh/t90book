package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func getCurrentID(c *gin.Context) {
	url := "https://www.facebook.com/gaming/T90Official"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to retrieve id",
		})
		return
	}

	// Necessary "browser" headers
	req.Header.Add("authority", "www.facebook.com")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("accept-language", "en")

	res, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to retrieve id",
		})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to retrieve id",
		})
		return
	}

	r, _ := regexp.Compile("videoId\":\"(.*)\",\"isPremiere")

	if !r.Match(body) {
		c.Error(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to retrieve id",
		})
		return
	}
	id := r.FindStringSubmatch(string(body))[1]

	// id = "453129616218866"
	c.JSON(http.StatusOK, gin.H{
		"id":  id,
		"url": fmt.Sprintf("https://www.facebook.com/T90Official/videos/%v/", id),
	})
}
