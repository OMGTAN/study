route delete 0.0.0.0
route -p add 1.1.1.1 mask 255.255.255.255 192.168.8.1 METRIC 1
route -p add 0.0.0.0 mask 0.0.0.0 10.104..0.1 METRIC 100
route -p add 0.0.0.0 mask 0.0.0.0 192.168.8.1 METRIC 200




@REM NETSH INT IP RESET  （重置IP设置）
@REM NETSH WINSOCK RESET （充值网络设置）
@REM NETSH WINHTTP RESET PROXY  （重置代理设置）
@REM IPCONFIG /FLUSHDNS  （刷新DNS缓存）