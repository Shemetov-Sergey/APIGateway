package main

import (
	"log"
	"strings"

	"github.com/Shemetov-Sergey/APIGateway/pkg/auth"
	"github.com/Shemetov-Sergey/APIGateway/pkg/comment"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/Shemetov-Sergey/APIGateway/pkg/gonews"
	"github.com/Shemetov-Sergey/APIGateway/pkg/models"
	"github.com/Shemetov-Sergey/APIGateway/pkg/textParser"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	in := make(chan models.CreateCommentRequestBody)
	out := make(chan models.CreateCommentRequestBody)
	errChan := make(chan error)
	censoredWords := strings.Split(c.Censored, ",")

	cd := textParser.New(in, out, errChan, censoredWords)
	cd.Run()

	authSvc := *auth.RegisterRoutes(r, &c)
	gonews.RegisterRoutes(r, &c, &authSvc)
	comment.RegisterRoutes(r, &c, &authSvc, in, out)

	err = r.Run(c.Port)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
