#!/bin/bash

APP=${PWD##*/}

logs() {
  CONTAINERIDS=$(sudo docker ps -aq --filter ancestor="$APP")
  sudo docker logs -f "${CONTAINERIDS}"
}

start() {
  CONTAINERIDS=$(sudo docker ps -aq --filter ancestor="$APP")
  for CONTAINERID in ${CONTAINERIDS}
  do
    sudo docker stop "${CONTAINERID}"
    sudo docker rm "${CONTAINERID}"
  done

  sudo docker build -t "$APP" .
  NEWCONTAINERID=$(sudo docker run -d --net='host' --env BEEGO_RUNMODE=prod "$APP")
  sudo docker logs -f "${NEWCONTAINERID}"
}

update() {
  git pull
  start
}

stop() {
  CONTAINERIDS=$(sudo docker ps -aq --filter ancestor="$APP")
  sudo docker stop "${CONTAINERIDS}"
}

if [ "$1" == "update" ]
then
  update
elif [ "$1" == "start" ]
then
  start
elif [ "$1" == "logs" ]
then
  logs
  elif [ "$1" == "stop" ]
then
  stop
else
  update
fi