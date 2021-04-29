#!/bin/sh

cd RT_Thread_SDK
echo $1
echo $2
scons --app=sdk/$1
cp sdk/$1/sdk/$2 ../
cd ..
chmod 755 $2


