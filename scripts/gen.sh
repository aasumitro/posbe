#!/bin/bash

FILE=.env
if test -f "$FILE"; then
  set="abcdefghijklmonpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  rand=""
  for i in $(seq 1 64); do
      char=${set:$RANDOM % ${#set}:1}
      rand+=$char
  done
  echo $rand

#  TODO UPDATE .ENV JWT_SECRET_KEY
else
  echo "==========================================================="
  echo "|  $FILE (environment) file does not exist.                |"
  echo "|  Please Crete new .env file from .env.example.          |"
  echo "|  by running this script: //:~$ cp .env.example .env     |"
  echo "==========================================================="
  exit 0
fi