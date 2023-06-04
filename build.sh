rm -rf build

mkdir -p build/chrome
cp -r src/shared/* build/chrome
cp -r src/chrome/* build/chrome

mkdir -p build/firefox
cp -r src/shared/* build/firefox
cp -r src/firefox/* build/firefox