Quick Admin Guide
=================

Source directory:
	
	cd /root/go/src/github.com/STARTeurope/startuplive.in/

Deploy updates:

	cd ~
	./deploy.bash
	
Change github username used for deploys:

	vim /root/go/src/github.com/STARTeurope/startuplive.in/.git/config
	Change:
	url = https://ungerik@github.com/STARTeurope/startuplive.in.git
	To:
	url = https://YOUR_USERNAME@github.com/STARTeurope/startuplive.in.git

Repair and restart MongoDB:

	cd /root/go/src/github.com/STARTeurope/startuplive.in/
	./stopmongodb.bash
	./repairmongodb.bash
	./startmongodb.bash

View MongoDB log file:

	tail -n 20 -f /var/log/mongodb/mongodb.log

View server log file:

	tail -n 20 -f /root/go/src/github.com/STARTeurope/startuplive.in/nohup.out
	tail -n 60 -f /root/go/src/github.com/STARTeurope/startuplive.in/log.txt

Backup MongoDB:

	cd /root/go/src/github.com/STARTeurope/startuplive.in/
	./stopmongodb.bash
	zip -r9T mongodb.zip /var/lib/mongodb/*
	./startmongodb.bash

Update Go:

	cd /opt/go/src
	hg pull
	hg update release
	./all.bash

Get source:

	go get -u github.com/STARTeurope/startuplive.in



Ubuntu 10.04 LTS 64bit Setup Guide
==================================

Update System:

	apt-get update
	apt-get upgrade

Set hostname:
	
	echo "startuplive.in" > /etc/hostname
	hostname -F /etc/hostname
	echo "178.79.174.140 startuplive.in startuplive.in" >> /etc/hosts
	/etc/init.d/networking restart

Install Go Dependencies:

	apt-get install bison gawk gcc libc6-dev make
	apt-get install python-setuptools python-dev build-essential
	easy_install mercurial
	apt-get install bzr
	apt-get install git-core
	
Create SSH Key for GitHub:

	ssh-keygen -t rsa -C "erik.unger@starteurope.at"
	cat .ssh/id_rsa.pub
	
Install Go:

	cd /opt
	hg clone -u release https://go.googlecode.com/hg/ go
	cd go/src
	./all.bash

Set bin path:

	cd ~
	echo "export PATH=/opt/go/bin:$PATH" >> .profile
	source .profile

Get startuplive.in source (first clone needs username):

	git clone https://ungerik@github.com/STARTeurope/startuplive.in.git /root/go/src/github.com/STARTeurope/startuplive.in
	go get github.com/STARTeurope/startuplive.in

Install MongoDB:

	echo "deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen" >> /etc/apt/sources.list
	apt-key adv --keyserver keyserver.ubuntu.com --recv 7F0CEB10
	apt-get update
	apt-get install mongodb-10gen
	stop mongodb
	rm /etc/mongodb.conf
	ln -s /root/go/src/github.com/STARTeurope/startuplive.in/mongodb.conf /etc/mongodb.conf
	start mongodb

Configure MongoDB:

	mongo
	use admin
	db.addUser("admin", "-eWin3TE5c _3eJs")
	db.auth("admin", "-eWin3TE5c _3eJs")
	use startuplive
	db.addUser("startuplive", "D 4ej-_7ohZ8UAXj2cPP")
	exit


Deploy latest version and restart server:

	cd /root/go/src/github.com/STARTeurope/startuplive.in/
	./deploy.bash

Dev Environment (local)
==================================

If there are any errors with launchpad.net/mgo then switch to launchpad.net/mgo directory and run
	
	go get -u
