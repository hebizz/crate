#!/bin/sh

cd RT_Thread_SDK
echo $1
echo $2
export RTT_ROOT=/home/admin/crate/RT_Thread
export BSP_ROOT=/home/admin/crate/RT_Thread/bsp/stm32/stm32f429-jx-dk1g
scons --app=sdk/$1
cp sdk/$1/sdk/$2 ../
cd ..
chmod 755 $2


