@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=
@echo //  ���Զ���ȡ��
@echo //  ����IP���������룬����
@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=

%1 mshta vbscript:CreateObject("Shell.Application").ShellExecute("cmd.exe","/c %~s0 ::","","runas",1)(window.close)&&exit
cd /d "%~dp0"


netsh interface ip set address name="WLAN" source=dhcp 
netsh interface ip set address name="��̫��" source=dhcp 

@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=
@echo //  ����DNS
@echo //-=-=-=-=-=-=-=-=-=-=-=-=-=-=
netsh interface ip set dns name = "WLAN" static addr = none   
netsh interface ip set dns name = "��̫��" static addr = none 
@echo //  �Զ���ȡIP������ɣ�
@pause