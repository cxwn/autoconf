#!/bin/bash

docker run -d --name oasr-autoconf \
	--restart always \
	-p 20101:20101 \
	-e JTR_REDIS_SERVER='192.168.98.211' \
	-e JTR_REDIS_PORT='30010' \
	-e JTR_REDISPASSWORD='' \
	-e JTR_ISCLUSTER='false' \
	-e JTR_VOICETEMPPATH='/root/tempvoice' \
	-e JTJ_URL='jdbc:mysql://192.168.98.211:3306/cmos_oasr_db?&characterEncoding=UTF8' \
	-e JTJ_USERNAME='cmos_oasr' \
	-e JTJ_PASSWORD='W7fMnb4i+NQbKRMMPD//ia7I4oiaHu6nLjKlQ/3pudEJ1Brgq1/wEnTsKFhus3hrC9hAlVeqRyZFNfD2nrHZ2w==' \
	-e OFC_MYSQLSERVERIP='192.168.98.211' \
	-e OFC_MYSQLSERVERPORT='3306' \
	-e OFC_MYSQLDATABASE='cmos_oasr_db' \
	-e OFC_MYSQLUSER='cmos_oasr' \
	-e OFC_MYSQLPASSWORD='W7fMnb4i+NQbKRMMPD//ia7I4oiaHu6nLjKlQ/3pudEJ1Brgq1/wEnTsKFhus3hrC9hAlVeqRyZFNfD2nrHZ2w==' \
	-e OFC_REDISCLUSTER='false' \
	-e OFC_REDISIP='192.168.98.211' \
	-e OFC_REDISPORT='30010' \
	-e OFC_REDISADD='192.168.98.211:30010' \
	-e OFC_REDISPASS='' \
	-e OFC_TASKCTRLSHELL='/Offline_System_Client/ctrlTask.sh' \
	-e JTL_LOG4J_APPENDER_MYRFILE1_FILE='/root/api/jsonrpc.log' \
    -e OCS_THREADS='2' \
	oasr:1.0.0
