package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

type TemplateData struct {
	HostNumber int
}

func main() {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Printf("Hosting on %s\n", ip.String())
	}

	r := gin.New()
	r.GET("/rancher/:hostnum", func(g *gin.Context) {
		hostNum, err := strconv.Atoi(g.Param("hostnum"))
		if err != nil {
			g.String(http.StatusBadRequest, "Invalid host number %s: %v", g.Param("hostnum"), err)
			return
		}

		t, err := template.ParseFiles("rancher.yml")
		if err != nil {
			g.String(http.StatusInternalServerError, "Unable to read template: %v", err)
			return
		}

		if err := t.Execute(g.Writer, TemplateData{HostNumber: hostNum}); err != nil {
			g.String(http.StatusInternalServerError, "Error rendering template: %v", err)
			return
		}
	})
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
