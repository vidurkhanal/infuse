package main

import (
	"github.com/vidurkhanal/infuse/cmd"
	"github.com/vidurkhanal/infuse/config"
)

func main() {
	_ = config.Providers()
	cmd.NewApp().StartHTTPServer()
}
