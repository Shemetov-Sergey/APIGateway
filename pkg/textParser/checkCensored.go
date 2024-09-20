package textParser

import (
	"strings"

	"github.com/Shemetov-Sergey/APIGateway/pkg/models"
)

type Censored struct {
	inChan        chan models.CreateCommentRequestBody
	outChan       chan models.CreateCommentRequestBody
	errChan       chan error
	censoredWords []string
}

func New(in, out chan models.CreateCommentRequestBody, errChan chan error, censoredWords []string) *Censored {
	return &Censored{
		inChan:        in,
		outChan:       out,
		errChan:       errChan,
		censoredWords: censoredWords,
	}
}

func (c *Censored) CheckCensored(text string) bool {
	words := strings.Split(text, " ")
	for _, word := range words {
		for _, cw := range c.censoredWords {
			if word == cw {
				return true
			}
		}
	}

	return false
}

func (c *Censored) Run() {
	go func() {
		for {
			select {
			case cm := <-c.inChan:
				censored := c.CheckCensored(cm.Text)
				cm.Censored = censored
				c.outChan <- cm
			}
		}
	}()
}
