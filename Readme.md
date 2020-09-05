# Simple game example with Golang, Vue.js, and Docker

`backend` folder contains server app written on Go which provide API and websocket for the game.  
`website` contains frontend written based on Vue.js.

## Requirements to run locally

[Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/).

## Steps to run locally

#### 1. Run containers
```
$ docker-compose build
$ docker-compose up -d
```

#### 2. Copy and add backend host

Copy ID of Docker container `netgame_backend_container` using `$ docker ps`. Get container IP with below command:
```
$ docker inspect -f '{{ .NetworkSettings.Networks.netgame_backend.IPAddress }}' [container ID]
```

Set this IP for host `backend_container`. On linux in file `/etc/hosts`.

#### 4. Open website in the browser

Following same steps copy IP of container `netgame_app`. Open on the browser `[IP]:8080`.

## Instruction to start the game

- Open website in the browser
- Clink "Play" button
- Enter player name and adjust two number
- Click "Submit" button
- Repeat above steps for second player

Backend is limited to 10 repeated games and can be adjusted. Between games pause takes 10 seconds.

## Additional

Ranking can be obtained with request `GET http://backend_container:8081/v1/game/ranking`.
