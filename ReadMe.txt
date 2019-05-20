
用golang编译成Reflective DLL,这里dllmain.h里直接用的是x64，如果想修改成x86的请参考ReflectiveDLL源码

Build
go build -i -v -o Reflective.dll -buildmode=c-shared -ldflags "-w -s -X main.version=1.1"



参考
https://github.com/NaniteFactory/hookwin10calc
https://github.com/stephenfewer/ReflectiveDLLInjection