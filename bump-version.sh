#!/usr/bin/env bash

flag=$1

version=$(cat version)

major=$(echo $version | awk -F'.' '{print $1}')
minor=$(echo $version | awk -F'.' '{print $2}')
patch=$(echo $version | awk -F'.' '{print $3}')

((patch++))


# write new version to file, only if write flag is set
if [ "$flag" = "-w" ]; then
    echo $major.$minor.$patch > version
fi

echo $major.$minor.$patch