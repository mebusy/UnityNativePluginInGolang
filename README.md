# UnityNativePluginInGolang

writing unity native plugin by golang

env: MacOSX, cross compile to build plugin for other platform.


- src/
    - golang source code
- dist/
    - binary library code
- OSXBundle
    - OSX bundle project
    - lib search path: `$PROJECT_DIR/../dist/OSX`
- Win
    - `$ brew install mingw-w64`
- Android
    - [c++ runtime](https://developer.android.com/ndk/guides/cpp-support.html)
    - `$ wget https://dl.google.com/android/repository/android-ndk-r20-darwin-x86_64.zip` 
    - NDK: android-ndk-r20
    - in UNITY:  `Assets/Plugins/Android/libs/<ARM_ARCH>/`

```bash
$NDK/build/tools/make-standalone-toolchain.sh --arch=arm --abis=armeabi-v7a --platform=android-16  --install-dir=ndkR10eCgoToolchain

# --toolchain=arm-linux-androideabi-4.9 
```

```bash
$ arm-linux-androideabi-gcc -print-multi-lib
.;@Wl,--no-warn-mismatch
armv7-a;@march=armv7-a@Wl,--no-warn-mismatch
thumb;@mthumb@Wl,--no-warn-mismatch
armv7-a/thumb;@march=armv7-a@mthumb@Wl,--no-warn-mismatch
armv7-a/hard;@march=armv7-a@mfloat-abi=hard@Wl,--no-warn-mismatch
armv7-a/thumb/hard;@march=armv7-a@mthumb@mfloat-abi=hard@Wl,--no-warn-mismatch

# -mfpu=  e.g. vfpv3, neon(此设置会强制使用 VFPv3-D32)
```



Platform | Plugins
--- | --- 
OSX  | xcode -> OSX bundle

