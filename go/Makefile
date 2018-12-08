LINUX=linux
WINDOWS=windows
DARWIN=darwin
AMD64=amd64
X86=386

.PHONY: build
all: build

build:
	go build -o magvk .

build_win:
	mkdir -p build
	GOARCH=${X86} GOOS=${WINDOWS} go build -o build/magvk_${WINDOWS}_${X86}.exe .
	GOARCH=${AMD64} GOOS=${WINDOWS} go build -o build/magvk_${WINDOWS}_${AMD64}.exe .

build_linux:
	mkdir -p build
	GOARCH=${X86} GOOS=${LINUX} go build -o build/magvk_${LINUX}_${X86} .
	GOARCH=${AMD64} GOOS=${LINUX} go build -o build/magvk_${LINUX}_${AMD64} .

build_mac:
	mkdir -p build
	GOARCH=${X86} GOOS=${DARWIN} go build -o build/magvk_${DARWIN}_${X86} .
	GOARCH=${AMD64} GOOS=${DARWIN} go build -o build/magvk_${DARWIN}_${AMD64} .
