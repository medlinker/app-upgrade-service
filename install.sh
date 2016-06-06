#!/bin/bash

currentDir=`pwd`

echo "create tmp dir ..."
cd ../
mkdir -p tmp/src
mkdir -p tmp/bin
mkdir -p tmp/pkg

cd tmp

chDir=`pwd`

echo "set gopath for tmp..."
export GOPATH="$chDir"

cp -r $currentDir/vendor/* ./src/
cp -R $currentDir ./src/

cd src/app-upgrade-service

echo 'compile...'
go build -o $currentDir/bin/appUpgradeService -ldflags '-s -w'

rm -rf $chDir

exit 0
