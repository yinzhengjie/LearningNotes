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
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# vim README.md 
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# 
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# git add *
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# 
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# git commit -m "first commit"
    [master 1c67c9f] first commit
     1 file changed, 5 insertions(+)
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# 
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# git push -u origin master
    Counting objects: 8, done.
    Delta compression using up to 2 threads.
    Compressing objects: 100% (4/4), done.
    Writing objects: 100% (6/6), 952 bytes | 0 bytes/s, done.
    Total 6 (delta 1), reused 0 (delta 0)
    remote: Resolving deltas: 100% (1/1), done.
    To git@github.com:yinzhengjie/LearningNotes.git
       08b06ae..1c67c9f  master -> master
    Branch master set up to track remote branch master from origin.
    [root@hadoop101.yinzhengjie.com ~/LearningNotes]# 

