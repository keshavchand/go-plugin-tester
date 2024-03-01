plugin1/plugin.so: plugin1/main.go
	go build -buildmode=plugin -o plugin1/plugin.so plugin1/main.go

plugin2/plugin.so: plugin2/main.go
	go build -buildmode=plugin -o plugin2/plugin.so plugin2/main.go

plugin_test: main.go
	go build -o plugin_test

.PHONY: run clean

run: plugin_test plugin1/plugin.so plugin2/plugin.so
	./plugin_test -plugin plugin1/plugin.so 
	./plugin_test -plugin plugin2/plugin.so

clean: 
	rm -f plugin_test plugin1/plugin.so plugin2/plugin.so
