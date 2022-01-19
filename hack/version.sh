#!/bin/bash

version=""

function getVersion {
  if [ -f "../VERSION" ]; then
    version=$(cat ../VERSION)
  fi

  if [ -z "$version" ]; then
      if [ -d ".git" ]; then
        version=$(git symbolic-ref HEAD | cut -b 12-)-$(git rev-parse HEAD)
      else
        version="0"
      fi
  fi
}

getVersion
