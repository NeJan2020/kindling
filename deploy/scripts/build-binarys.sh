cd ../../
mkdir -p probe/build
cd probe/build
cmake -DBUILD_DRIVER=OFF ..
make
libKindlingPath="./src/libkindling.so"
if [ ! -f "$libKindlingPath" ]; then
  echo "compiler libkindling failed! exit!"

else
  cp -rf ./src/libkindling.so ../../collector/docker/libso
  cp -rf ./src/libkindling.so /usr/lib64/
  cd ../../collector/
  go build -o docker/kindling-collector
fi