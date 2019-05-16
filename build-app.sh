#!/usr/bin/env bash

rm -rf dist/zoom-status.app

pushd icons || exit 1
bash generate-icons.sh
bash generate-menu-icon.sh
popd || exit 1

go build github.com/cmmarslender/zoom-status

mkdir -p dist/zoom-status.app/Contents/
mkdir -p dist/zoom-status.app/Contents/MacOS
mkdir -p dist/zoom-status.app/Contents/Resources

cp Info.plist dist/zoom-status.app/Contents/
cp zoom-status dist/zoom-status.app/Contents/MacOS/
cp icons/icon.icns dist/zoom-status.app/Contents/Resources/

