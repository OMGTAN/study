@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=
@echo //  ���̶����á�
@echo //  ����IP���������룬����
@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=

%1 mshta vbscript:CreateObject("Shell.Application").ShellExecute("cmd.exe","/c %~s0 ::","","runas",1)(window.close)&&exit
cd /d "%~dp0"

::netsh interface ip set address name="WLAN" source=static addr=192.168.0.230 mask=255.255.255.0 gateway=192.168.0.1 

::netsh interface ip set address name="WLAN" source=static addr=192.168.0.230 mask=255.255.255.0 
netsh interface ip set address name="��̫��" source=static addr=172.30.126.105 mask=255.255.255.0 gateway=172.30.126.1

@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=
@echo //  ����DNS
@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=
::netsh interface ip set dns "WLAN" static 114.114.114.114 primary 
::netsh interface ip add dns "WLAN" 8.8.8.8 
@echo //  �̶�IP������ɣ�
@pause