do {
"building"  ;
env GOOS=linux GOARCH=arm GOARM=6 CC=arm-linux-musleabihf-gcc CXX=arm-linux-musleabihf-g++ CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" main.go ;
mv -f main daikin-api ;
scp daikin-api ariel@raspberrypi.local:dev/daikin
} | complete
