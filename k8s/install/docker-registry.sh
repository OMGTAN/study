# 本地镜像服务
sudo docker run -d -p 5000:5000 --restart=always --name registry -v /home/suaee/registry:/var/lib/registry registry:2