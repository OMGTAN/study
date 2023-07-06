# 安装docker
sudo yum install yum-utils -y
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

sudo yum -y install docker-ce
# 安装完成后启动，并加入开机启动
sudo systemctl start docker && sudo systemctl enable docker
# 启动完成后查看
sudo docker info



sudo vi /etc/docker/daemon.json

{ "bip": "10.255.1.1/24" }

sudo systemctl restart docker.service 