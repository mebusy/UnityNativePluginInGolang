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
    - `$ wget https://dl.google.com/android/repository/android-ndk-r20-darwin-x86_64.zip` 
    - NDK: android-ndk-r20



Platform | Plugins
--- | --- 
OSX  | xcode -> OSX bundle

