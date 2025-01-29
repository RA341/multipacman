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

dkp:
    just dkbd
    docker login
    docker push  ras334/multipacman:dev

prune:
    docker image prune -f

[working-directory: 'frontend']
ui:
    flutter build web
    rm -r ../core/web/
    cp -r build/web ../core/web
