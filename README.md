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
- Go

### Steps to Configure and Use

1. **Clone the repository**
   git clone https://github.com/souru98/Assignment_Scalable_User.git
2. **Build the Docker image**
   docker build -t lms-userservice .
3. **Run the Docker container**
  docker run -e PORT=8081 -p 8081:8081 lms-userservice


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

   
