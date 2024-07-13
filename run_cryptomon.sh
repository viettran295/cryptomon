#!/bin/bash 

sudo apt update -y 
sudo apt install conky-all 
make all 
conky -d