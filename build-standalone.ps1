if (Test-Path -Path "build") {
    Remove-Item -Path "build" -Recurse -Force
}

New-Item -ItemType Directory -Path "build" | Out-Null

Copy-Item -Path "resources\*" -Destination "build\resources\" -Recurse -Force

Set-Location "src"
go build -ldflags "-s -w" -o "../build/GameServer.exe" main.go
