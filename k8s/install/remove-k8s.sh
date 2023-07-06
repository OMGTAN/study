sudo kubeadm reset -f
sudo modprobe -r ipip
sudo lsmod
sudo rm -rf ~/.kube/
sudo rm -rf /etc/kubernetes/
sudo rm -rf /etc/systemd/system/kubelet.service.d
sudo rm -rf /etc/systemd/system/kubelet.service
sudo rm -rf /usr/bin/kube*
sudo rm -rf /etc/cni
sudo rm -rf /opt/cni
sudo rm -rf /var/lib/etcd
sudo rm -rf /var/etcd
sudo yum -y remove kubeadm* kubectl* kubelet* docker*
sudo reboot