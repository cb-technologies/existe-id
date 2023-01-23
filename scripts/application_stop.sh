#!/bin/bash


# Stopping exist server that is already running
sudo echo "`date` [Stop Exist Server]: Stopping the exist server" >> /var/log/exist_id_server/exist_server.log
sudo systemctl stop exist_id.service