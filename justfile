set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

list:
    just --list

# sets up dev envoirment
setup:
    just set-flu
    just set-go

# sets up go dev, adds a web dir and builds flutter web
[working-directory: 'core']
set-go:
    mkdir web
    just ui
    go mod tidy

# sets up flutter dev
[working-directory: 'frontend']
set-flu:
    flutter doctor
    flutter pub get

dkbd:
    docker build . -t ras334/multipacman:dev

dkr:
    just dkbd
    docker run --rm -v ./appdata:/app/appdata -p 11200:11200 ras334/multipacman:dev 

dkp:
    just dkbd
    docker login
    docker push  ras334/multipacman:dev

dklt:
    docker build . -t ras334/multipacman:latest

dpl:
    just dklt
    docker login
    docker push  ras334/multipacman:latest

prune:
    docker image prune -f

[working-directory: 'frontend']
ui:
    flutter build web
    rm -r ../core/web/
    cp -r build/web ../core/web
