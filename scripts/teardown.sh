#!/bin/bash

DIR=${PWD##*/}

docker rmi $(docker images | grep $DIR | grep -v client | awk '{ print $1 }')