#!/bin/bash

openssl genrsa -out private.key 1024
openssl rsa -in private.key -pubout -out public.key