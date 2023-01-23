#!/bin/bash

# Starting exist id server service
sudo echo "`date` [Start Exist Server]: Starting the exist server" >> /var/log/exist_id_server/exist_server.log
sudo systemctl start exist_id.service