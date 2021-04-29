#! /bin/bash
mkdir -p /home/admin/crate
apt-get install supervisor -y
apt-get install scons gcc-arm-none-eabi -y
cp crate.conf /etc/supervisor/conf.d/
mkdir -p /data/logs/supervisor
service supervisor restart
cp crate /usr/local/bin
cp main /home/admin/crate/
cp compile.sh /home/admin/crate/
cd /home/admin/crate
git clone  https://touyunaxiaozi:6651358.hebi@gitee.com/JiangxingDev/RT_Thread.git
git clone  https://touyunaxiaozi:6651358.hebi@gitee.com/JiangxingDev/RT_Thread_SDK.git
echo "export RTT_ROOT=/home/admin/crate/RT_Thread" >> /etc/profile 
echo "export BSP_ROOT=/home/admin/crate/RT_Thread/bsp/stm32/stm32f429-jx-dk1g" >> /etc/profile
source /etc/profile
cp RT_Thread_SDK/SConscript /home/admin/crate/RT_Thread_SDK/sdk/SConscript

rm -rf /home/admin/crate/RT_Thread/.git
rm -rf /home/admin/crate/RT_Thread_SDK/.git

rm -rf /home/admin/crate/RT_Thread/bsp/stm32/stm32f429-jx-*/dlmodule_ctl.c 
rm -rf /home/admin/crate/RT_Thread/bsp/stm32/stm32f429-jx-*/dlmodule_ctl.h
