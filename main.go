package main

import (
	"errors"
	"flag"
	"log"
	"plugin"
	iface "plugin_test/interface"
)

func main() {
	var pluginPath string
	flag.StringVar(&pluginPath, "plugin", "", "path to the plugin")
	flag.Parse()

	worker, err := GetWorkerFromPlugin(pluginPath)
	if err != nil {
		log.Fatal(err)
	}
	worker.Do()
}

func GetWorkerFromPlugin(path string) (iface.IWorker, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, (err)
	}

	fn, err := plug.Lookup("GetWorker")
	if err != nil {
		return nil, (err)
	}

	var worker iface.IWorker

	workfactory, ok := fn.(func() iface.IWorker)
	if ok {
		worker = workfactory()
	} else {
		return nil, errors.New("unexpected type from module symbol")
	}

	return worker, nil
}
