#!/bin/bash

cd ..
rm -f autoconf
git pull
go build autoconf.go
docker inspect oasr >/dev/null 2>&1
# shellcheck disable=SC2181
[ $? -eq 0 ] && docker rm -f oasr
bash test/oasr-docker-run.sh
# docker cp autoconf oasr:/usr/local/bin
docker logs oasr
docker exec -it oasr bash
