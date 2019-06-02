#! /bin/sh

GOOS=linux go build -o $2 "$1"
GOOS=linux go build -ldflags="-s -w" -o $2.-sw "$1"
goupx -f --brute -o $2.upx $2
goupx -f --brute -o $2.-sw.upx $2.-sw

GOOS=linux gotip build -o $2.tip "$1"
GOOS=linux gotip build -ldflags="-s -w" -o $2.tip.-sw "$1"
goupx -f --brute -o $2.tip.upx $2.tip
goupx -f --brute -o $2.tip.-sw.upx $2.tip.-sw

ls -l $2*
