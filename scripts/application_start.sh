#!/bin/bash

# Giving permission to the file inside /home/ec2-user/existe_id_server
sudo chmod 777 /home/ec2-user/existe_id_server

#start our app in the background
sudo systemctl restart existe_id.service