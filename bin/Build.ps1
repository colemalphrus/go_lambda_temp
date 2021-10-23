$env:GOOS = "linux"

go build -o main

~\Go\Bin\build-lambda-zip.exe -o main.zip main