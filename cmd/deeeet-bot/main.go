package main

import (
	"flag"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

type IncommingMesage struct {
	Token       string `json:"token" form:"token" binding:"required"`
	TeamId      string `json:"team_id" form:"team_id" binding:"required"`
	ChannelId   string `json:"channel_id" form:"channel_id" binding:"required"`
	ChannelName string `json:"channel_name" form:"channel_name" binding:"required"`
	Timestamp   string `json:"timestamp" form:"timestamp" binding:"required"`
	UserId      string `json:"user_id" form:"user_id" binding:"required"`
	UserName    string `json:"user_name" form:"user_name" binding:"required"`
	Text        string `json:"text" form:"text" binding:"required"`
	TriggerWord string `json:"trigger_word" form:"trigger_word"`
}

var (
	re   = regexp.MustCompile(`\bde+t\b`)
	addr = flag.String("addr", defaultAddr(), "server address")
)

func defaultAddr() string {
	if s := os.Getenv("PORT"); s != "" {
		return ":" + s
	}
	return ":8080"
}

func main() {
	flag.Parse()

	gin.DefaultWriter = colorable.NewColorableStderr()
	r := gin.Default()
	r.POST("/v1/slack/inbound", func(c *gin.Context) {
		var msg IncommingMesage
		err := c.Bind(&msg)
		if err != nil {
			c.Error(err)
			return
		}
		if re.MatchString(msg.Text) && !strings.Contains(msg.Text, "deeeet") {
			msg.Text = "deeeet です..."
			c.JSON(200, msg)
		}
	})
	r.Run(*addr)
}
