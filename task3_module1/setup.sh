#!/bin/bash
echo installing hugo and make...
apt-get update && apt-get install -y make wget 
wget https://github.com/gohugoio/hugo/releases/download/v0.101.0/hugo_0.101.0_Linux-64bit.deb
apt install ./hugo_0.101.0_Linux-64bit.deb
echo building website...
make build
echo done!
