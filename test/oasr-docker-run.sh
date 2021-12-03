#!/bin/bash

docker run -d --name oasr \
	--restart always \
	-p 20101:20101 \
	-e JTR_REDIS_QUEUE='voicetask' \
	-e JTR_REDIS_SERVER='172.25.209.248' \
	-e JTR_REDIS_PORT='30010' \
	-e JTR_REDISPASSWORD='' \
	-e JTR_ISCLUSTER='false' \
	-e JTR_VOICETEMPPATH='/root/tempvoice' \
	-e JTJ_URL='jdbc:mysql://172.25.34.222:3306/cmos_oasr_db?&characterEncoding=UTF8' \
	-e JTJ_USERNAME='cmos_oasr' \
	-e JTJ_PASSWORD='XQ0htWPRgxlXAynJjexYZ/16YInwatW3MJSNjiZAAbg=' \
	-e OFC_MYSQLSERVERIP='172.25.34.222' \
	-e OFC_MYSQLSERVERPORT='3306' \
	-e OFC_MYSQLDATABASE='cmos_oasr_db' \
	-e OFC_MYSQLUSER='cmos_oasr' \
	-e OFC_MYSQLPASSWORD='XQ0htWPRgxlXAynJjexYZ/16YInwatW3MJSNjiZAAbg=' \
	-e OFC_REDISCLUSTER='false' \
	-e OFC_REDISIP='172.25.209.248' \
	-e OFC_REDISPORT='30010' \
	-e OFC_REDISADD='172.25.209.248:30010' \
	-e OFC_REDISPASS='' \
	-e OFC_REDISTASKLIST='voicetask' \
	-e OFC_REDISNOTIFYLIST='zyzx-rqy-version-1.0.0' \
	-e OFC_TASKCTRLSHELL='/Offline_System_Client/ctrlTask.sh' \
	-e JTL_LOG4J_APPENDER_MYRFILE1_FILE='/root/api/jsonrpc.log' \
	oasr:1.0.0
