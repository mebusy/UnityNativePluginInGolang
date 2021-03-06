# to list all GOOS/GOARCH:  go tool dist list

PLATFORM=$1

if [ "$PLATFORM" = "" ] ; then
    echo "must specify a platform Darwin|Windows|Android|iOS"
    exit 1
fi



if [ "$PLATFORM" = "Darwin" ] ; then
    echo Building Plugin for $PLATFORM

    # OSX
    DIST="../dist/OSX/"
    mkdir -p $DIST
    GOOS=darwin GOARCH=amd64 go build -v -ldflags "-w" -buildmode=c-archive -o $DIST/libgo.a  libgo/

    echo $DIST
    # build OSX bundle
    xcodebuild -project ../OSxBundle/OSXBundle.xcodeproj -scheme OSXBundle  -configuration Release
fi

if [ "$PLATFORM" = "Windows" ] ; then
    echo Building Plugin for $PLATFORM
    
    # Win64
    DIST="../dist/Win64/"
    echo $DIST
    mkdir -p $DIST
    GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc  CXX=x86_64-w64-mingw32-g++  go build -v -ldflags "-w" -buildmode=c-shared -o $DIST/libgo.dll  libgo/

    cp -f $DIST/libgo.dll ../Unity_PluginTest/Assets/Plugins/
fi

if [ "$PLATFORM" = "Android" ] ; then
    echo Building Plugin for $PLATFORM

    # Android
    TOOLCHAIN="$NDK/../ndkR10eCgoToolchain"
    TARGET="arm-linux-androideabi"
    # TARGET="aarch64-linux-android"
    API=
    CC=$TOOLCHAIN/bin/$TARGET$API-gcc
    CXX=$TOOLCHAIN/bin/$TARGET$API-g++
    
    ARM_ARCH=armeabi-v7a
    DIST="../dist/Android/libs/$ARM_ARCH"
    echo $DIST
    mkdir -p $DIST
    GOOS=android GOARCH=arm  CGO_ENABLED=1 \
        CC=$CC CXX="$CXX" \
        go build -v -ldflags "-w" -buildmode=c-shared -o $DIST/libgo.so  libgo/


    cp -f $DIST/libgo.so ../Unity_PluginTest/Assets/Plugins/Android/libs/$ARM_ARCH/

fi

if [ "$PLATFORM" = "iOS" ] ; then
    echo Building Plugin for $PLATFORM

    CLANG=`xcodebuild -find clang`
    CLANGXX=`xcodebuild -find clang++`
    # echo clang: "$CLANG"
    SYSROOT=`xcrun --sdk iphoneos --show-sdk-path`

    DIST="../dist/iOS/"
    mkdir -p $DIST
    GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 \
    CC=$CLANG CXX=$CLANGXX  \
    CGO_LDFLAGS="$CGO_LDFLAGS -isysroot $SYSROOT -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64 " \
    CGO_CFLAGS="$CGO_CFLAGS   -isysroot $SYSROOT -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64 " \
    go build -v -ldflags "-w" -buildmode=c-archive -o $DIST/libgo.a  libgo/


fi


