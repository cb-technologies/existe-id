#!/bin/bash

#THIS SCRIPT IS USED TO LOAD SOME NEW CONFIGURATION IF NEEDED WHEN THERE HAS BEEN A NEW PUSH TO THE MASTER CODE BEFORE RESTARTING THE SERVER


# THIS PART HANDLES THE SSL CERTIFICATE FOR THE HTTPS SERVER

#A new ssl certificate is needed for the exist server because currently we are using a self signed certificate
#for our https server. And everytime, we stop and restart the ec2 instance, the full domain name or common name
#of the server certificate changes. So, we need to configure a new ssl certificate for the exist server that will run on 
#the ec2 that has a different Public DNS Name
# FOr more information on this, visit https://mem.ai/p/AWHuCVHBYeLZMITNCSBU

#MAKE SURE TO CHANGE THIS true value TO false WHEN THERE HAS BEEN NO CHANGE TO THE PUBLIC DNS IP  OF THE EC2 INSTANCE
#USUALLY THE CHANGE HAPPENS WHEN THE EC2 INSTANCE IS STOP THEN RESTARTED
if false ; then

  sudo echo "`date` [Configure SSL]: There has been an update to Public DNS(which means an Update to the self_signed_ssl_certificate.ssl.txt file) file or It has been created for the first time" >> /var/log/exist_id_server/exist_server.log

  #First removing the old SSL files
  #Logging that we are running the command
  sudo echo "`date` [Configure SSL]: Running the command 'sudo rm -rf /etc/httpd/ssl/*.pem'" >> /var/log/exist_id_server/exist_server.log
  sudo rm -rf /etc/httpd/ssl/*.pem

  #NOw creating a new ssl key
  #Logging the creation of the new ssl key"
  sudo echo "`date` [Configure SSL]: Creating the new ssl key" >> /var/log/exist_id_server/exist_server.log
  sudo openssl genpkey -algorithm RSA -out /etc/httpd/ssl/key.pem

  #Now generating the certificate signing request
  #Logging the creation of the certificate signing request
  sudo echo "`date` [Configure SSL]: Creating the certificate signing request" >> /var/log/exist_id_server/exist_server.log
  sudo openssl req -new -key /etc/httpd/ssl/key.pem -out /etc/httpd/ssl/csr.pem -subj "$(cat /home/ec2-user/exist_id_server/useraccount/config/self_signed_ssl_certificate.txt)"

  #Now generating the certificate
  #Logging the creation of the certificate
  sudo echo "`date` [Configure SSL]: Creating the certificate" >> /var/log/exist_id_server/exist_server.log
  sudo openssl x509 -req -days 365 -in /etc/httpd/ssl/csr.pem -signkey /etc/httpd/ssl/key.pem -out /etc/httpd/ssl/cert.pem 

else
  sudo echo "`date` [Configure SSL]: There has been no update to the Public DNS file" >> /var/log/exist_id_server/exist_server.log
fi

# THIS PART RESTART THE APACHE SERVER SO THAT THE NEW SSL CERTIFICATE CAN BE USED BY THE EXIST SERVER
sudo echo "`date` [Configure SSL]: Restarting the apache server" >> /var/log/exist_id_server/exist_server.log
sudo systemctl restart httpd


# THIS PART HANDLES REBUILDING THE SERVER BINARY CODE AFTER THE NEW PUSH TO THE MASTER CODE
sudo echo "`date` [Rebuilding Server Binary Code]: Rebuilding the server binary code" >> /var/log/exist_id_server/exist_server.log
go build -o /home/ec2-user/exist_id_server/useraccount/cmd/existserver /home/ec2-user/exist_id_server/useraccount/cmd/main.go

# THIS PART HANDLES RELOADING THE SYSTEMD CONFIGURATION FILES BECAUSE THERE HAS BEEN A CHANGE TO THE SERVER BINARY CODE
# THIS WILL RELOAD THE CONFIGURATION OF THE EXIST SERVER SERVICE
sudo echo "`date` [Reload Systemd Configuration]: Reloading the systemd configuration" >> /var/log/exist_id_server/exist_server.log
sudo systemctl daemon-reload
