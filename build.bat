set GOARCH=amd64
set GOOS=windows
set CGO_ENABLED=1

go build -tags osusergo -v -ldflags "-s -w" -o winResize.exe

echo go build -tags osusergo -v -compiler=gccgo -gccgoflags "-s -w" -o ./winResize_gccgo.exe
