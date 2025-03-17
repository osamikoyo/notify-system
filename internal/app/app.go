package app

import (
	"fmt"
	"time"

	"github.com/koyo-os/notify-system/pkg/logger"
)

func App() error {
	counter := 0

	var timerChan chan struct{}

	go func ()  {
		counter++

		time.Sleep(time.Second)
	}()

	fmt.Println(`
 __    _  _______  _______  ___   _______  __   __ 
|  |  | ||       ||       ||   | |       ||  | |  |
|   |_| ||   _   ||_     _||   | |    ___||  |_|  |
|       ||  | |  |  |   |  |   | |   |___ |       |
|  _    ||  |_|  |  |   |  |   | |    ___||_     _|
| | |   ||       |  |   |  |   | |   |      |   |  
|_|  |__||_______|  |___|  |___| |___|      |___|  
`)

	logger := logger.Init()

	logger.Info("start notify...")


}