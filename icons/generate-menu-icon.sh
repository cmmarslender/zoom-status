#/bin/sh

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go get github.com/cratonica/2goarray
    if [ $? -ne 0 ]; then
        echo Failure executing go get github.com/cratonica/2goarray
        exit
    fi
fi

#if [ -z "$1" ]; then
#    echo Please specify a PNG file
#    exit
#fi
#
#if [ ! -f "$1" ]; then
#    echo $1 is not a valid file
#    exit
#fi

SIZE=44
sips -z $SIZE $SIZE icon-original.png --out menu_icon.png

OUTPUT=iconunix.go
echo Generating $OUTPUT
echo "//+build linux darwin" > $OUTPUT
echo >> $OUTPUT
cat "menu_icon.png" | $GOPATH/bin/2goarray Data icon >> $OUTPUT
if [ $? -ne 0 ]; then
    rm menu_icon.png
    echo Failure generating $OUTPUT
    exit
fi
rm menu_icon.png
echo Finished