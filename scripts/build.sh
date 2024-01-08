#!/bin/bash

current_directory=$(pwd)
echo "Current working directory: $current_directory"

# Set the path to the deployment folder
deployment_folder="/deployment"

# Set the name of the Dockerfile
dockerfile_name="Dockerfile"

# Set the name of the Docker image
docker_image_name="api"

# Copy the Dockerfile to the current directory
cp "$deployment_folder/$dockerfile_name" $current_directory

# Build the Docker image
docker build -t "$docker_image_name" .

# Clean up by removing the copied Dockerfile
rm "$dockerfile_name"
