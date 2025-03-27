package main

import (
	"flag"
	"github.com/WattsonG42/GwinMon/internal/config"
	"github.com/WattsonG42/GwinMon/internal/logger"
	"github.com/WattsonG42/GwinMon/internal/service"
	"log"
	"time"
)

func main() {

	verboseFlag := flag.Bool("v", false, "Verbose output (CLI logging)")
	logFlag := flag.Bool("log", false, "Log output (file logging)")
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if *verboseFlag {
		cfg.Verbose = true
	}

	if cfg.Interval <= 0 {
		cfg.Interval = 60
	}

	logger.Init(cfg.Verbose, *logFlag)
	defer logger.Close()

	logger.Info("GwinMon is starting...")

	for _, svc := range cfg.Services {

		s := svc
		go service.MonitorService(s.Name, s.ExpectedStatus, cfg.Interval, func(message string) {
			logger.Info(message)
		}) // Fix panic handling
	}

	for {
		time.Sleep(1 * time.Hour)
	}

}
