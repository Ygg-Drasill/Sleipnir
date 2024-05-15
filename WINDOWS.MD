USE ``go get github.com/bytecodealliance/wasmtime-go``\
Needs to be version 1.0.0

Install msys64
https://www.msys2.org/ \
Make sure the path is C:\msys64


Enable CGO (Try 'set' first, if that doesn't work use 'export')
```
set CGO_ENABLED=1
export CGO_ENABLED=1
```
Confirm ``CGO_ENABLED=1`` with ``go env | grep CGO_ENABLED``

Now add msys64 to the path
``export PATH=$PATH:/c/msys64/mingw64/bin``

Now you can run ``go build -o main.exe`` at main.go
