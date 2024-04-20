# User Service

## Overview
This repository contains the code for the User Service, a part of a microservices architecture application. The User Service is responsible for managing the users in a book rental system.

## Features
- Fetch all users
- Fetch user by ID
- Create new user
- Update existing user
- Authenticate user

## Technologies Used
- Go: An open source programming language that makes it easy to build simple, reliable, and efficient software.
- SQLite: A C library that provides a lightweight disk-based database.

## Getting Started

### Prerequisites
- Docker
- MiniKube
- Go
- SqlLite

### Steps to Configure and Use

1. **Clone the repository**
   ```
   git clone https://github.com/souru98/Assignment_Scalable_User.git
2. **Build the Docker image**
   ```
   docker build -t lms-userservice .
3. **Run the Docker container**
   ```
   docker run -e PORT=8081 -p 8081:8081 lms-userservice

## Deploying to MiniKube

Follow these steps to deploy your application to MiniKube:

1. **Navigate to Project Directory**: 
   Open Windows PowerShell and navigate to your project directory using the `cd` command.

2. ** change MiniKube Driver to Docker**
   ```
   minikube config set driver docker

3. **Start MiniKube**: 
   Start your MiniKube cluster with the command 
   ```
   minikube start`.

4. ** Initialize MiniKube Env**
   ```
   minikube docker-env

5. **Set Docker Environment**: 
   Set up the Docker environment inside MiniKube. Run the following command in PowerShell:
   ```powershell
   minikube -p minikube docker-env --shell powershell | Invoke-Expression
   
6. **Build Docker Image**
   ```powershell
   docker build -t lms-userservice .
   
7. **Create Kubernetes Deployment**
   ```powershell
   kubectl run lms-us-mkk22 --image=lms-userservice --image-pull-policy=Never --port=8080
   
8. **Port Forwarding**
   ```powershell
   kubectl port-forward lms-us-mkk22 8080


## API Endpoints
- `/users`: Fetch all users
- `/users/:id`: Fetch user by ID
- `/user`: Create new user
- `/user/:id`: Update existing user
- `/auth`: Authenticate user

## Contributing
Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## License
This project is licensed under the MIT License - see the LICENSE.md file for details

   
