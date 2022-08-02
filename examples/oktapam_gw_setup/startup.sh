#!/bin/bash
echo "Retrieve information about new packages"
sudo apt-get update
sudo apt-get install -y curl

echo "Trust the repository signing key"
curl -fsSL https://dist.scaleft.com/pki/scaleft_deb_key.asc | gpg --dearmor | sudo tee /usr/share/keyrings/scaleft-archive-keyring.gpg > /dev/null

echo "Add the ASA repos to the repolist"
printf "deb [arch=amd64 signed-by=/usr/share/keyrings/scaleft-archive-keyring.gpg] http://pkg.scaleft.com/deb focal main\ndeb [arch=amd64 signed-by=/usr/share/keyrings/scaleft-archive-keyring.gpg] http://pkg.scaleft.com/deb linux main" | sudo tee /etc/apt/sources.list.d/scaleft.list > /dev/null

echo "Retrieve information about new packages"
sudo apt-get update

echo "Install gateway"
sudo apt-get install scaleft-gateway
echo ${gwtoken} > /var/lib/sft-gatewayd/setup.token
sudo service sft-gatewayd restart
