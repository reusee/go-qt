PACKAGE_PATH = ${GOPATH}/src/smoke_info

all:
	g++ extract_info.cpp -o extract_info -lsmokeqtcore -lsmokeqtgui -lsmokebase --std=c++11
	./extract_info
	gofmt -w info.go
	mkdir ${PACKAGE_PATH} -p
	cp -f *.go ${PACKAGE_PATH}
	cd ${PACKAGE_PATH}; go install
