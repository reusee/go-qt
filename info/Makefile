all:
	g++ extract_info.cpp -o extract_info -lsmokeqtcore -lsmokeqtgui -lsmokebase --std=c++11
	./extract_info
	gofmt -w info.go
	go build
