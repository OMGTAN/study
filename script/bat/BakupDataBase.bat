:start

rem "����Զ������IP���û��������롢�ļ�·��"
set ftphost=10.108.17.27
set username=tan
set password=123456
set filepath=/home/backup/database/3306
rem "���ñ���·��"
set Ldir=D:/��Ŀ/2020/MEB����/���ݿⱸ��
set timevar=%time:~0,2%
set tmpScript=sftp.text
if /i %timevar% LSS 10 (
	set timevar=0%time:~1,1%
)
set filename=db_%date:~0,4%-%date:~5,2%-%date:~8,2%-%timevar%.sql



echo cd %filepath% >>tmpScript

echo lcd %Ldir% >>tmpScript

echo get %filename% >>tmpScript

echo bye >>tmpScript



psftp -P 22 %username%@%ftphost% -pw %password% -b tmpScript

del tmpScript