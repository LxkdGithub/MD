添加库类似如下(这里是一种，还有一种是ppa库)
    sudo add-apt-repository "deb [arch=amd64] https://xxxx.com/xxxx"
如果出错公钥失效则要重新获取,使用如下命令
    sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys '目标key' 即可
最后更新即可
    sudo apt-get update

## 删除库
### [完整连接](https://linux.cn/article-11115-1.html)
1./etc/sources.list删除链接
2.如果添加ppa 到/etc/sources.list.d删除对应文件
3.删除key
    sudo  apt-key del xxxx xxxx 
    添加是apt-key add xxxx即可

## 最后
add-apt-repository时ppa和deb的区别
其实没太大区别，一个写入sources.list,
一个写入sources.list.d的文件中
甚至可以用deb的方式添加ppa
