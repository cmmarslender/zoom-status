#!/usr/bin/env bash

mkdir -p out.iconset

SIZE=512
sips -z $SIZE $SIZE icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}.png

SIZE=256
SIZE2x=$(($SIZE * 2))
sips -z $SIZE $SIZE icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}.png
sips -z $SIZE2x $SIZE2x icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}@2x.png

SIZE=128
SIZE2x=$(($SIZE * 2))
sips -z $SIZE $SIZE icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}.png
sips -z $SIZE2x $SIZE2x icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}@2x.png

SIZE=64
SIZE2x=$(($SIZE * 2))
sips -z $SIZE $SIZE icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}.png
sips -z $SIZE2x $SIZE2x icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}@2x.png

SIZE=32
SIZE2x=$(($SIZE * 2))
sips -z $SIZE $SIZE icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}.png
sips -z $SIZE2x $SIZE2x icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}@2x.png

SIZE=16
SIZE2x=$(($SIZE * 2))
sips -z $SIZE $SIZE icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}.png
sips -z $SIZE2x $SIZE2x icon-original.png --out out.iconset/icon_${SIZE}x${SIZE}@2x.png

iconutil -c icns -o icon.icns out.iconset
