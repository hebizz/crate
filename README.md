## 启动server 端： 

* go run main.go 

## 导入环境变量,安装依赖： 
* git clone https://gitee.com/JiangxingDev/RT_Thread.git
* git clone https://gitee.com/JiangxingDev/RT_Thread_SDK.git
* export BSP_ROOT=/home/admin/crate/RT_Thread/bsp/stm32/stm32f429-jx-ttu(ttu)
* export BSP_ROOT=/home/admin/crate/RT_Thread/bsp/stm32/stm32f429-jx-net(rt)   
* export RTT_ROOT=/home/admin/crate/RT_Thread
* apt-get install scons gcc-arm-none-eabi
* chmod 755 /dev/ttysWK3（开机自启,rt）
* chmod 755 /dev/ttyS1   (开机自启，ttu)
* 在RT_Thread_SDK的sdk目录下提前放好Sconsript文件


## cli端支持的功能:

* **创建项目:**
 crate create mcu容器名称

* **上传mcu容器c文件:**
 crate upload 文件路径

* **编译并上传mcu容器:**
 crate compile mcu容器名称

* **启动mcu容器:**
 crate dlmodule start mcu容器名称

* **停止mcu容器:**
 crate dlmodule stop mcu容器名称

* **检查mcu容器状态:**
 crate dlmodule check mcu容器名称

* **删除mcu容器:**
 crate dlmodule rm mcu容器名称

* **列出所有mcu容器:**
 crate dlmodule list 

* **设置mcu容器开机模式:**
 crate dlmodule auto mcu容器名称 模式 (这里模式分为0和1, 1为开机启动)

* **获取mcu侧所有支持的指令:**
 crate help

