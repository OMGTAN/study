# 设置主机名
sudo hostnamectl set-hostname k8s-node01 && bash
# hostnamectl set-hostname k8s-node01

#关闭selinux
sudo sed -i 's/^ *SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config

#关闭防火墙
sudo systemctl stop firewalld.service
sudo systemctl disable firewalld

# 关闭交换分区
sudo sed -ri 's/.*swap.*/#&/' /etc/fstab
sudo swapoff -a

# 配置k8s源
sudo cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

#  永久添加模块
sudo vim /etc/modules-load.d/k8s.conf
overlay
br_netfilter

# 手动添加
sudo modprobe br_netfilter
sudo modprobe overlay

# 查看添加效果
lsmod | grep br_netfilter

# 配置 k8s 网络配置
sudo cat > /etc/sysctl.d/k8s.conf <<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
sudo sysctl -p /etc/sysctl.d/k8s.conf
sudo sysctl --system

# 安装ipvs
sudo yum -y install ipset ipvsadm
sudo vim /etc/sysconfig/modules/ipvs.modules

modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack

sudo chmod 755 /etc/sysconfig/modules/ipvs.modules && sudo bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack

# 安装基础包 看情况安装
yum install -y yum-utils device-mapper-persistent-data lvm2 wget net-tools nfs-utils lrzsz gcc gcc-c++ make cmake libxml2-devel openssl-devel curl curl-devel unzip sudo ntp libaio-devel wget vim ncurses-devel autoconf automake zlibdevel python-devel epel-release openssh-server socat ipvsadm conntrack ntpdate telnet ipvsadm

# 安装Containerd
sudo yum install containerd -y
sudo systemctl start containerd && sudo systemctl enable containerd

# 创建Containerd配置文件
sudo mkdir -p /etc/containerd
sudo containerd config default > /etc/containerd/config.toml
#老版containerd
#替换配置文件 
#替换grc.io为阿里源
sudo sed -i "s#k8s.gcr.io#registry.aliyuncs.com/google_containers#g" /etc/containerd/config.toml
#替换containerd驱动为SystemCgroup，与kubelet保持一致
sudo sed -i '/containerd.runtimes.runc.options/a\ \ \ \ \ \ \ \ \ \ \ \ SystemdCgroup = true' /etc/containerd/config.toml
#镜像仓库地址
sudo sed -i "s#https://registry-1.docker.io#https://registry.aliyuncs.com#g" /etc/containerd/config.toml
#重启生效配置
sudo systemctl restart containerd

tar xvf containerd-1.6.12-linux-amd64.tar.gz
sudo systemctl stop containerd
cd bin
sudo cp * /usr/bin/
sudo systemctl start containerd


#新版containerd
#命令根据https://blog.csdn.net/zhh763984017/article/details/126714567 修改，增加registry.k8s.io替换
sudo mkdir -p /etc/containerd && \
sudo containerd config default > /etc/containerd/config.toml && \
sudo sed -i "s#k8s.gcr.io/pause#registry.aliyuncs.com/google_containers/pause#g"  /etc/containerd/config.toml  && \
sudo sed -i 's#SystemdCgroup = false#SystemdCgroup = true#g' /etc/containerd/config.toml  && \
sudo sed -i 's#registry.k8s.io/pause:3.6#registry.aliyuncs.com/k8sxio/pause:3.6#g' /etc/containerd/config.toml  && \
sudo sed -i '/registry.mirrors]/a\ \ \ \ \ \ \ \ [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]' /etc/containerd/config.toml  && \
sudo sed -i '/registry.mirrors."docker.io"]/a\ \ \ \ \ \ \ \ \ \ endpoint = ["http://hub-mirror.c.163.com"]' /etc/containerd/config.toml && \
sudo sed -i '/hub-mirror.c.163.com"]/a\ \ \ \ \ \ \ \ [plugins."io.containerd.grpc.v1.cri".registry.mirrors."k8s.gcr.io"]' /etc/containerd/config.toml  && \
sudo sed -i '/"k8s.gcr.io"]/a\ \ \ \ \ \ \ \ \ \ endpoint = ["http://registry.aliyuncs.com/google_containers"]' /etc/containerd/config.toml && \
sudo echo "===========restart containerd to reload config===========" && \
sudo systemctl restart containerd

# 安装k8s
#查看可安装的kubernetes 的版本
dnf list kubelet --showduplicates | sort -r

#安装
sudo dnf install -y kubelet-1.27.0 kubeadm-1.27.0 kubectl-1.27.0 --disableexcludes=kubernetes

#配置开机自启
sudo systemctl enable --now kubelet

#Kubeadm: kubeadm 是一个工具，用来初始化 k8s 集群的
#kubelet: 安装在集群所有节点上，用于启动 Pod 的
#kubectl: 通过 kubectl 可以部署和管理应用，查看各种资源，创建、删除和更新各种组件

#查看版本
kubeadm version

#指定容器运行时为containerd
sudo crictl config runtime-endpoint /run/containerd/containerd.sock

# 节点加入控制 24小时过期
sudo kubeadm join 10.21.18.91:6443 --token abcdef.0123456789abcdef --discovery-token-ca-cert-hash sha256:f4adbae15af1fe985431fe2ec60307f2e66694058992b4dd5485836ccb80f237 


# 配置kubectl配置文件config
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> ~/.bash_profile
source ~/.bash_profile

export KUBECONFIG=/etc/kubernetes/admin.conf

# 检查节点
kubectl get nodes


# 配置节点 仅node节点
# 仅node节点上执行
mkdir -p $HOME/.kube

# 同步.kube/config文件夹
xsync ~/.kube/config

# 此命令为init之后打印在控制台中
# 仅node节点执行
kubeadm join 192.168.2.40:6443 --token abcdef.0123456789abcdef \
        --discovery-token-ca-cert-hash sha256:09bdee3b5da241080291a2aed41939a801dc5932a2caedf90b1fac8831450acb

#执行后可能存在以下报错，因为存在cri-docker与containerd两个容器运行时
Found multiple CRI endpoints on the host. Please define which one do you wish to use by setting the 'criSocket' field in the kubeadm configuration file: unix:///var/run/containerd/containerd.sock, unix:///var/run/cri-dockerd.sock
To see the stack trace of this error execute with --v=5 or higher
#手动指定容器运行时为containerd，其他情况同理
kubeadm join 192.168.2.40:6443 --token abcdef.0123456789abcdef \
        --discovery-token-ca-cert-hash sha256:09bdee3b5da241080291a2aed41939a801dc5932a2caedf90b1fac8831450acb \
				--cri-socket=unix:///var/run/containerd/containerd.sock

#未记录可执行此命令重新打印
kubeadm token create --print-join-command

#查看节点标签
kubectl get nodes --show-labels

# 设置标签
kubectl label nodes k8s-node01 node-role.kubernetes.io/work=work

# 网络插件 Calico/flannel

#下载后修改IP如下图 192.168 改成k8s配置的子网络
curl https://projectcalico.docs.tigera.io/manifests/calico.yaml -O

# 修改完成后执行
kubectl apply -f calico.yaml

# 下载busybox
docker pull busybox

# 使用busybox测试
kubectl run busybox --image busybox --restart=Never --rm -it busybox -- sh
ping www.baidu.com
nslookup kubernetes.default.svc.cluster.local
exit
