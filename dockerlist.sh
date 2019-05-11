#!/usr/bin/env bash

echo -n "Username: "
read username
echo -n "Password: "
read -s password
echo ""

curl http://$username:$password@127.0.0.1:5050/v2/_catalog
