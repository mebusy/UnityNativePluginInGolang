# UnityNativePluginInGolang

writing unity native plugin by golang

env: MacOSX, cross compile to build plugin for other platform.


src/
    golang source code
dist/
    binary library code
OSXBundle
    OSX bundle project
    lib search path: `$PROJECT_DIR/../dist/OSX`
Win
    $ brew install mingw-w64



Platform | Plugins
--- | --- 
OSX  | xcode -> OSX bundle

