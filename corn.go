package main

import (
	"github.com/jacklove/go-gin-example/corn"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		corn.CornTag()
	})

	c.Start()

	t1 := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t1.C:
			c.Stop()

		}
	}
}
