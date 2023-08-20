package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jlaffaye/ftp"
)

type FTPServer struct {
	Host     string
	Path     string
	Username string
	Password string
}

var File string

func ftpConnection() string {
	FTP := &FTPServer{Host: os.Getenv("ftp-host"), Path: os.Getenv("ftp-path"), Username: os.Getenv("ftp-user"), Password: os.Getenv("ftp-pass")}
	c, err := ftp.Dial(FTP.Host, ftp.DialWithTimeout(30*time.Second))
	if err != nil {
		fmt.Printf("ftp error: %s", err)
	}

	err = c.Login(FTP.Username, FTP.Password)
	if err != nil {
		fmt.Printf("ftp credentials error: %s", err)
	}

	c.ChangeDir(FTP.Path)
	xmlFile, err := c.Retr(File)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	buffer, _ := io.ReadAll(xmlFile)
	return string(buffer)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.ForwardedByClientIP = true

	g.GET("/:file", func(c *gin.Context) {
		c.Header("Host", "FarmingSimulator-ExtendedAPI/Golang")
		c.Header("Content-Type", "application/xml")
		File = c.Param("file")
		c.Status(http.StatusOK)
		c.Writer.Write([]byte(ftpConnection()))
	})

	g.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Host", "FarmingSimulator-ExtendedAPI/Golang")
		c.Header("Content-Type", "image/x-icon")
		c.Status(http.StatusOK)
	})

	_ = g.Run(":8095")
}
