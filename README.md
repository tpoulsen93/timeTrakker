# timeTrakker
## Requirements
Docker 

Linux

## Dev Container
After opening vscode, when prompted to run in a dev container... do so :)
The dev container already has nodejs, npm, typescript, go, and all the other packages we need to run our 3 applications.

The server and web app can be run from within the dev container.
The mobile app needs to be run from the host machine
(unless you can figure out how to make the docker container ip address available on your local network).

## Applications

### web
`pnpm install`

`pnpm start`

### mobile
`yarn install`

`yarn start`

### server
`go run main.go`
