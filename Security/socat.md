Socat 是 Linux 下的一个多功能的网络工具，名字来由是 「Socket CAT」。其功能与有瑞士军刀之称的 Netcat 类似，可以看做是 Netcat 的加强版。

Socat 的主要特点就是在两个数据流之间建立通道，且支持众多协议和链接方式。如 IP、TCP、 UDP、IPv6、PIPE、EXEC、System、Open、Proxy、Openssl、Socket等。
[官网](http://www.dest-unreach.org/socat/)

### 反弹shell(发送终端)
    客户机: socat - tcp:ip:8888 (ip为服务端ip)
---
    服务端: socat tcp-l:8888 exec:sh,pty,stderr

### 发送文件
    发送端: socat -u open:hello.txt tcp-l:8888
---
    接收端: socat -u tcp:IP:8888 open:hello.txt,create,append

### 执行命令(这里是cat, 发什么数据返回什么)
    客户端: socat - tcp:IP:port
    相当于: nc IP port
---
    服务端: socat tcp-l:port exec:"/bin/cat",fork
    这里加和不加fork的区别是会不会为每一个链接单独开一个进程
    不加fork
    501 27202 27201   0 10:25AM ttys003    0:00.26 -bash
    501 27501 27202   0 10:37AM ttys003    0:00.00 socat -v tcp-l:8181 exec:/bin/cat
    此单进程一直工作
    bash->socat

    加fork
    501 27519 27202   0 10:41AM ttys003    0:00.00 socat -v tcp-l:8182,fork exec:/bin/cat
    501 27527 27519   0 10:44AM ttys003    0:00.00 socat -v tcp-l:8182,fork exec:/bin/cat
    501 27540 27519   0 10:47AM ttys003    0:00.00 socat -v tcp-l:8182,fork exec:/bin/cat
    不同的链接会由有不同的进程
    bash->socat->不同的进程工作