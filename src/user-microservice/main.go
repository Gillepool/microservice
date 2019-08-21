package main

import (
	"github.com/gin-gonic/gin"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	return err
}

func main() {
	m := Main{}

	if m.initServer != nil {
		return
	}

}
