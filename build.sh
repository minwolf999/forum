#!/bin/bash
# Build Docker image
sudo docker build -t forum .
# Run Docker container
sudo docker run -p 8080:8080 forum
# Remove Docker image
sudo docker rmi -f forum