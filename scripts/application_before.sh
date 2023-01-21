#!/bin/bash

#create our working directory if it doesn't exist
DIR="/home/ec2-user/existe_id_server"
if [ -d "$DIR" ]; then
  echo "Directory $DIR exists."
else
  echo "Directory $DIR does not exist. Creating it now."
  mkdir $DIR
fi