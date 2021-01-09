#!/bin/bash

GOPROXY=https://goproxy.io
cd src && go build -i -v -o finance-spider
#cd console && npm install && npm run build
#cd ../theme && npm install && npm run build

echo 'build done'
