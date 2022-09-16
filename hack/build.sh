#!/bin/bash

if [ "$1" == "all" ]; then
  ./release.sh activity && \
  ./release.sh account && \
  ./release.sh public && \
  ./release.sh actor && \
  ./release.sh article && \
  ./release.sh auth && \
  ./release.sh channel && \
  ./release.sh device && \
  ./release.sh fs && \
  ./release.sh message && \
  ./release.sh saved && \
  ./release.sh gw
fi