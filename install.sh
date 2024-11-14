#!/bin/bash

# Update package list
sudo apt-get update -y

# Download and install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install Docker Compose as a plugin (recommended way)
sudo apt-get install -y docker-compose-plugin

# Add the vagrant user to the Docker group to allow running Docker without sudo
sudo usermod -aG docker vagrant

# Clean up
rm get-docker.sh

# Verify installations
docker --version || echo "Docker command not found."
docker compose version || echo "Docker Compose command not found."

echo "Provisioning complete."
