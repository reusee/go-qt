all:
	g++ codegen.cpp -o cpp_codegen -lsmokeqtcore -lsmokeqtgui -lsmokebase --std=c++11
	./cpp_codegen
	gofmt -w info.go
	go build
	./codegen
