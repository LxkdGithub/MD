## 安装docker
### 可以在linux上直接安装，但是不是最新版，所以不用此法(以下为ubuntu)
[完整链接](https://imroc.io/posts/docker/quick-start-for-docker-3/)

    sudo apt-get update

    sudo apt-get -y install apt-transport-https ca-certificates curl software-properties-common

    使用aliyun的库(docker官方库太慢)
    curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

    sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"

    sudo apt-get update 
    sudo apt-get install docker-ce