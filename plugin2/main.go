package main

import (
	"log"
	iface "plugin_test/interface"
)

type WorkerB struct {
}

func (w *WorkerB) Do() error {
	log.Println("WorkerB doing")
	return nil
}

func GetWorker() iface.IWorker {
	return &WorkerB{}
}
