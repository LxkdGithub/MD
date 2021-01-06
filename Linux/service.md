首先是service和systemd一级

/etc/init.d 是sysVinit服务的启动方式,对于一些古老的系统或者服务 使用这个.
service 也是sysVinit, 比/etc/init.d先进一点,底层还是调用/etc/init.d
systemctl 是systemD命令的主要方式, 尽管一些老的系统或者命令不支持systemctl, 但是systemctl最后会逐渐的替代其他的命令方式的, 能用这个就优先用这个,是最时尚/方便的


init.d原始命令
    /etc/init.d/xxxxx start/stop/...等
    具体参数可以有什么在/etc/init.d/xxxx目录都有，这是真正的bash启动脚本
rc.d原始的自启动有
    update-rc.d xxxxx defaults
    update-rc.d remove xxxxx
    之类的
(service就是这两个的调用罢了)


[完整参考链接](https://blog.yuntianti.com/posts/initd-rcd-systemd/)
systemd最新推荐
systemctl的config文件存在 /lib/systemd/system/xxxxxx.serive
systemd就是init.d和rc.d的封装，底层的改变还是反映在了rc.d目录

怎么把自己的service加入systemd管理
    1.首先先把服务加入/etc/init.d/里边
    2.然后查找或者编写service文件放在/lib/systemd/system下每个service都有
        [Unit] [service] [install]三部分
        和1对应吧
    3.enable disable反映在rc.d的自启动里


Tips(常识):
- 当使用service能启动服务，但停止或查看status时提示找不到服务。多半是服务配置文件中修改了pid文件路径。 而/etc/init.d/script脚本依赖pid来管理服务进程状态，而且脚本中的pid路径是写死的。一旦/etc/init.d/script中使用的pid路径与实际pid文件路径不符时，就无法做后续的所有管理工作了。一旦启动， 就脱离service控制了。
- /etc/init.d/下的脚本需要有可执行权限 chmod +x。
- 大多init.d脚本会依赖与/etc/default/目录中的配置文件，遇到问题时，不妨也排查一下 这个目录中的配置文件是否正确。一般使用的配置文件都在/etc/xxxx/xxxxx.conf或者/etc/xxxx.conf中


