USE ``go get github.com/bytecodealliance/wasmtime-go``\
Needs to be version 1.0.0

Install msys64
https://www.msys2.org/ \
Make sure the path is C:\msys64

Open up Msys2 terminal.

Install the required gcc using pacman in the Msys2 terminal.
```
pacman -S mingw-w64-x86_64-gcc
```

Guides below show example with either Powershell or Git Bash

### Powershell
Open up Powershell\
Now add msys64 to the path
```
$Env:PATH += C:\msys64\mingw64\bin
```

Make sure gcc in enabled ``gcc --version``

Enable CGO
```
$Env:CGO_ENABLED = 1
```
Confirm ``CGO_ENABLED=1`` with ``$Env:CGO_ENABLED``

Now you can run ``go build -o main.exe`` at main.go

### Git Bash
Open up Git Bash \
Now add msys64 to the path
```
export PATH=$PATH:/c/msys64/mingw64/bin
```

Make sure gcc in enabled ``gcc --version``

Enable CGO
```
export CGO_ENABLED=1
```
Confirm ``CGO_ENABLED=1`` with ``go env | grep CGO_ENABLED``

Now you can run ``go build -o main.exe`` at main.go
