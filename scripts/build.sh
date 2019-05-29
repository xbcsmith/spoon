#!/bin/bash


DATE=$(date '+%Y%m%d%s')

export DOCKER_ID_USER=$DOCKER_ID_USER

docker login

docker build -t $DOCKER_ID_USER/spoon:$DATE -f Dockerfile.build .

docker push $DOCKER_ID_USER/spoon:$DATE

docker tag $DOCKER_ID_USER/spoon:$DATE $DOCKER_ID_USER/spoon:latest

docker push $DOCKER_ID_USER/spoon:latest

