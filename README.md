# JNDIScan

![](https://img.shields.io/badge/build-passing-brightgreen)
![](https://img.shields.io/badge/golang-1.17-blue)

## 介绍
这是什么？

该项目一款无须借助`dnslog`且完全无害的`JNDI`反连检测工具，解析`RMI`和`LDAP`协议实现，可用于甲方内网自查

特点：
- 由`Golang`编写，直接运行可执行文件进行检测
- 完全无害，不存在任何`Payload`仅基于`LDAP`和`RMI`协议检测
- 根据检测结果动态生成`HTML`页面
- 检测当前内网和外网的IP打印可用的`Payload`
- 支持在路径中带出参数（方便批量扫描）

![](https://github.com/EmYiQing/JNDIScan/blob/master/img/01.png)

## 使用

使用`${jndi:ldap://ip:端口/}`这样的`Payload`

如果目标存在漏洞，该项目就会收到`ldap/rmi`请求，从而快速定位哪些目标存在漏洞

![](https://github.com/EmYiQing/JNDIScan/blob/master/img/02.png)

命令：`./JNDIScan -p port(默认8001)`

```text
       ___   ______  _________
      / / | / / __ \/  _/ ___/_________ _____
 __  / /  |/ / / / // / \__ \/ ___/ __ `/ __ \
/ /_/ / /|  / /_/ // / ___/ / /__/ /_/ / / / /
\____/_/ |_/_____/___//____/\___/\__,_/_/ /_/
    coded by 4ra1n
[+] [11:55:41] use port: 8001
[+] [11:55:41] start fake reverse server
|-------------------------------------------|
|--Payload: ldap://192.168.1.4:8001---------|
|--Payload: ldap://127.0.0.1:8001-----------|
|--Payload: ldap://your-ip:8001-------------|
|--Payload: rmi://192.168.1.4:8001/xxx------|
|--Payload: rmi://127.0.0.1:8001/xxx--------|
|--Payload: rmi://your-ip:8001/xxx----------|
|-------------------------------------------|

```

只需要将`Payload`设置为输出的这些之一

即可在命令行看到连接结果，当前目录的`result.html`是动态渲染后的`HTML`页面

如果想带出参数，直接加入即可：`ldap://127.0.0.1:8001/PARAMS`

## 其他

其实该项目是我与某位大佬共同开发的，由于一些原因他让我删了他的ID

本来支持动态的web页面，由于一些原因被迫删除，并且不会再添加该功能