package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/OlivierArgentieri/go_killprocess/controllers"
	"github.com/judwhite/go-svc"
)

type program struct {
	LogFile *os.File
	svr     *controllers.Server
	ctx     context.Context
}

func (p *program) Context() context.Context {
	return p.ctx
}

func main() {
	ctx := context.Background()

	prg := program{
		svr: &controllers.Server{},
		ctx: ctx,
	}

	if err := svc.Run(&prg); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0])) todo
	dir := "C:/"

	logPath := filepath.Join(dir, "gokillprocess.log")
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	p.LogFile = f
	log.SetOutput(f)

	return nil
}

func (p *program) Start() error {
	log.Printf("Starting...\n")
	go p.svr.Run(":5119")
	return nil
}

func (p *program) Stop() error {
	log.Printf("Stopping... \n")
	go p.svr.Stop(":5119")
	return nil
}
