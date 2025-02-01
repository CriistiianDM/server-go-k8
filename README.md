# Kubernetes + Google Cloud

This project provides a setup for deploying a Go-based web server using Kubernetes on Google Cloud. It includes instructions for building and pushing the Docker image to Docker Hub, as well as running the development environment with docker-compose.

# Installation & Local Execution
```
    docker-compose up --build
```

* docker-compose up → Starts the containers defined in docker-compose.yml.

* --build → Forces a rebuild of the image before running the containers.

# Deploying the Image to Docker Hub

## 1. Build the Docker Image
```
    docker build -t [your-image-name] [path-to-Dockerfile]
```

## 2. Tag the Image for Docker Hub
```
    docker tag server [yourUserName]/server:1.0
```

## 3. Push the Image to Docker Hub
```
    docker tag push [yourUserName]/server:1.0
```

# Deploying on Kubernetes in Google Cloud
Navigate to the k8 folder and apply the YAML files:

```
    kubectl apply -f golang_server.yaml
    kubectl apply -f ingress.yaml
```