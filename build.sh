SERVER_DEV="http://localhost:3000"
SERVER_PROD="https://boxes.taiven.com"

rm -rf build

mkdir -p build/chrome-dev
cp -r src/shared/* build/chrome-dev
cp -r src/chrome/* build/chrome-dev
sed -i "" "s~{{SERVER}}~$SERVER_DEV~g" build/chrome-dev/api.js
sed -i "" "s~{{SERVER}}~$SERVER_DEV~g" build/chrome-dev/manifest.json

mkdir -p build/chrome-prod
cp -r src/shared/* build/chrome-prod
cp -r src/chrome/* build/chrome-prod
sed -i "" "s~{{SERVER}}~$SERVER_PROD~g" build/chrome-prod/api.js
sed -i "" "s~{{SERVER}}~$SERVER_PROD~g" build/chrome-prod/manifest.json

mkdir -p build/firefox-dev
cp -r src/shared/* build/firefox-dev
cp -r src/firefox/* build/firefox-dev
sed -i "" "s~{{SERVER}}~$SERVER_DEV~g" build/firefox-dev/api.js
sed -i "" "s~{{SERVER}}~$SERVER_DEV~g" build/firefox-dev/manifest.json

mkdir -p build/firefox-prod
cp -r src/shared/* build/firefox-prod
cp -r src/firefox/* build/firefox-prod
sed -i "" "s~{{SERVER}}~$SERVER_PROD~g" build/firefox-prod/api.js
sed -i "" "s~{{SERVER}}~$SERVER_PROD~g" build/firefox-prod/manifest.json