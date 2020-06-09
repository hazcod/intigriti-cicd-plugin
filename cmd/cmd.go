package main

import (
	"flag"
	"github.com/hazcod/intigriti-cicd-plugin/checker"
	"github.com/hazcod/intigriti-cicd-plugin/config"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	defaultConfName = "isa.yaml"
	defaultLogLevel = log.InfoLevel
	exitCodeOK  = 0
	exitCodeNOK = 2
)

func main() {
	confPath := flag.String("conf", defaultConfName, "The path to the configuration file.")
	logLevelStr := flag.String("loglevel", defaultLogLevel.String(), "The log level from debug (5) to error (1).")
	flag.Parse()

	if _, err := os.Stat(*confPath); os.IsNotExist(err) {
		log.Fatalf("configuration file not found: %s", *confPath)
	}

	logLevel, err := log.ParseLevel(*logLevelStr)
	if err != nil {
		log.Fatalf("invalid log level: %v", err)
	}
	log.SetLevel(logLevel)

	config, err := config.ParseConfig(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	exit, err := checker.RunChecker(config)
	if err != nil {
		log.Fatalf("could not verify submissions: %v", err)
	}

	if ! exit {
		log.Info("no outstanding submissions found")
		os.Exit(exitCodeOK)
	}

	log.Infof("open submissions found, returning %d", exitCodeNOK)
	os.Exit(exitCodeNOK)
}