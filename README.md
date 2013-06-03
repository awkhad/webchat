# web chat use websocket

# 需要的依赖

	mysql, redis
	golang, revel
	beedb mysql的orm

# how to run 

	# install revel:
	go get github.com/robfig/revel
	go build -o bin/revel github.com/robfig/revel/cmd

	# goalng package needs 
	go get github.com/astaxie/beedb
	go get github.com/go-sql-driver/mysql
    go get github.com/hoise/redis

	

# ubuntu 下安装依赖包

	sudo apt-get install mysql-server

# mysql数据库相关的操作使用rails 的migration

需要安装ruby, activerecord gem
	
	rake db:create
	rake db:migrate

手动创建数据库

	 CREATE DATABASE `webchat_dev` /*!40100 DEFAULT CHARACTER SET utf8 */;
	 USE webchat_dev;
	 source db/db.sql;




## 使用redis存储数据相关结构

* save message to list

	key: room:roomkey   
    format: type|username|text|time

* room 最近用户存储在集合中

    room:test:users

* user 信息 Hash

   key 是: user:id
   value 是: name, avatar, id

    
