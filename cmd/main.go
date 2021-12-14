package main

import config "learn-golang/pkg/config"

func main() {
	r := config.SetupRoute()
	r.Run(":5000")
}
