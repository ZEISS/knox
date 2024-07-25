#!/bin/bash
# This script is executed after the creation of a new project.

go install github.com/air-verse/air@latest

bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -b . ory
sudo mv ./ory /usr/local/bin/
