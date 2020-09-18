# LearningNotes
用于记录以后的学习笔记,


下面介绍如何使用本项目

1>.在GitHub创建项目
	略

2>.在Linux本地生成密钥对，并将公钥信息复制到GitHub中.
    [root@hadoop101.yinzhengjie.com ~]# ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa
    [root@hadoop101.yinzhengjie.com ~]# 
    [root@hadoop101.yinzhengjie.com ~]# cat ~/.ssh/id_rsa.pub 


3>.在Linux安装git工具，并将GitHub项目克隆到本地.
    [root@hadoop101.yinzhengjie.com ~]# yum -y install git
    [root@hadoop101.yinzhengjie.com ~]# 
    [root@hadoop101.yinzhengjie.com ~]# git clone git@github.com:yinzhengjie/LearningNotes.git

4>.将本地的代码移动到克隆的目录项目中
	略

5>.将linux本地文件推到GitHub上

[root@hadoop101.yinzhengjie.com ~/LearningNotes]# git config --global user.name "Jason Yin"
[root@hadoop101.yinzhengjie.com ~/LearningNotes]# 
[root@hadoop101.yinzhengjie.com ~/LearningNotes]# git config --global user.email "y1053419035@qq.com"
[root@hadoop101.yinzhengjie.com ~/LearningNotes]# 


