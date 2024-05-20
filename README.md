# GreenLight

GreenLight is a robust JSON API for retrieving and managing information about movies. This project was built following the guidance provided in Alex Edwards' book "Let's Go Further" on building REST APIs in Golang.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- Retrieve a list of movies with details such as title, description, release year, and genre
- Create, update, and delete movie records
- Search and filter movies based on various criteria
- Pagination support for efficient data retrieval
- Authentication and authorization for secure access to the API

## Technologies Used

- **Go**: The programming language used for the backend development
- **PostgreSQL**: The relational database management system used to store movie data
- **Docker**: Used to containerize the application and the database
- **GitHub Actions**: Utilized for continuous integration and deployment (CI/CD) workflows
- **Gin Web Framework**: A fast and lightweight web framework used for building the REST API
- **GORM**: An Object-Relational Mapping (ORM) library used for interacting with the database
- **JWT**: JSON Web Tokens used for authentication and authorization

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Docker](https://www.docker.com/get-started)
- [PostgreSQL](https://www.postgresql.org/download/) (optional, if not using Docker)

### Installation

1. Clone the repository:

```sh
git clone https://github.com/your-username/GreenLight.git
```

2. Change into the project directory:

```sh
cd GreenLight
```

3. Build the Docker images:

```sh
docker-compose build
```

### Running the Application

1. Start the Docker containers:

```sh
docker-compose up -d
```

2. The API will be available at `http://localhost:8000`.

## API Endpoints

The GreenLight API provides the following endpoints:

| Method | Endpoint     | Description               |
| ------ | ------------ | ------------------------- |
| GET    | /movies      | Retrieve a list of movies |
| POST   | /movies      | Create a new movie        |
| GET    | /movies/{id} | Retrieve a specific movie |
| PUT    | /movies/{id} | Update a movie            |
| DELETE | /movies/{id} | Delete a movie            |
