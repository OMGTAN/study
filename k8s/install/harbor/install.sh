安装：
https://zhuanlan.zhihu.com/p/638376379
证书：
https://zhuanlan.zhihu.com/p/234918875

官网

docker login localhost:8020 -u admin
Harbor12345

docker login -u admin -p Harbor12345  10.21.18.91:8020

docker login -u admin -p Harbor12345  harbor91

"insecure-registries": ["harbor91"],

"tlscacert": "D:/Users/certs/ca.crt",
  "tlscert": "D:/Users/certs/harbor91.cert",
  "tlskey": "D:/Users/certs/harbor91.key"

