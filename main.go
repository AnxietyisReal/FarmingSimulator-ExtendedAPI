package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jlaffaye/ftp"
)

type FTPServer struct {
	Host     string `json:"host"`
	Path     string `json:"path"`
	File     string `json:"file"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func JSONLoader(file string) FTPServer {
	var ftpServerFile FTPServer
	loader, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer loader.Close()
	parseJson := json.NewDecoder(loader)
	parseJson.Decode(&ftpServerFile)
	return ftpServerFile
}

func ftpConnection() string {
	FTP := JSONLoader("ftp_details.json")
	c, err := ftp.Dial(FTP.Host, ftp.DialWithTimeout(30*time.Second))
	if err != nil {
		fmt.Printf("ftp error: %s", err)
	}

	err = c.Login(FTP.Username, FTP.Password)
	if err != nil {
		fmt.Printf("ftp credentials error: %s", err)
	}

	c.ChangeDir(FTP.Path)
	xmlFile, err := c.Retr(FTP.File)
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

	g.GET("/", func(c *gin.Context) {
		c.Header("Host", "FarmingSimulator-ExtendedAPI/Golang")
		c.Header("Content-Type", "application/xml")
		c.Status(http.StatusOK)
		c.Writer.Write([]byte(ftpConnection()))
	})

	g.Run(":8095")
}
