do {
yarn build ;
mv -f build ui ;
scp -r ui ariel@raspberrypi.local:dev/daikin
scp -r package.json ariel@raspberrypi.local:dev/daikin/ui
scp -r yarn.lock ariel@raspberrypi.local:dev/daikin/ui
} | complete
