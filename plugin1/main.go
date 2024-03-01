package main

import (
	"log"
	iface "plugin_test/interface"
)

type WorkerA struct {
}

func (w *WorkerA) Do() error {
	log.Println("WorkerA doing")
	return nil
}

func GetWorker() iface.IWorker {
	return &WorkerA{}
}
