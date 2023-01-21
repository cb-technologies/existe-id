#!/bin/bash

#start our app in the background
go build -o /home/ec2-user/exist_id_server/useraccount/cmd/existserver /home/ec2-user/exist_id_server/useraccount/cmd/main.go

#Running the server
sudo nohup /home/ec2-user/exist_id_server/useraccount/cmd/existserver >> /home/ec2-user/exist_id_server/useraccount/existserver.log 2>&1 &