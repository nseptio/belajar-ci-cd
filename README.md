# Indonesia Student Survey - Backend API

This repository contains the backend API for the **Indonesia Student Survey** by Ditjen Diktiristek - Kemdikbudristek. The project is designed to collect feedback from final-year university students about their experiences.

## Table of Contents

- [Indonesia Student Survey - Backend API](#indonesia-student-survey---backend-api)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Environment Configuration](#environment-configuration)
    - [Running the Application](#running-the-application)
  - [Makefile Commands](#makefile-commands)
  - [Docker Support](#docker-support)

## Overview

The Indonesia Student Survey collects feedback from university students to help improve the quality of higher education. This backend API handles the core functionalities such as user management, survey submission, and reporting. The application is built using Go and follows clean architecture principles.

## Getting Started

### Prerequisites

- **Go** (version 1.23 or higher)
- **Docker** (for containerized deployments)
- **MongoDB** (as the database)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/indonesia-student-survey.git
   cd indonesia-student-survey
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

### Environment Configuration

The app uses `air.conf` for live reloading and `.env` for environment variables. Create a `.env` file in the root directory:

```
DB_HOST=localhost
DB_PORT=27017
DB_NAME=admin
DB_USERNAME=admin
DB_PASSWORD=admin
SERVER_PORT=8080
```

### Running the Application

You can run the application locally in development mode with live reloading using:

```bash
make watch
```

For a regular run:

```bash
make run
```

or using regular Go commands:

```bash
go run .
```

## Makefile Commands

The `Makefile` provides several useful commands:

- `run`: Run the web service.
- `watch`: Run the web service with live-reload (using [Air](https://github.com/cosmtrek/air)).
- `migrate-up`: Apply database migrations.
- `migrate-down`: Rollback the last migration.
- `migrate-reset`: Rollback all migrations.
- `migrate-to`: Go to a specific version of the database migration.
- `goose`: Run Goose CLI commands manually.

## Docker Support

This project includes a `Dockerfile` to build and run the application in a Docker container.

To build and run the application with Docker:

1. Build the Docker image:

   ```bash
   docker build -t indonesia-student-survey .
   ```

2. Run the container:

   ```bash
   docker run -p 8080:8080 --env-file .env indonesia-student-survey
   ```

Make sure to provide the correct environment variables for the database connection.