Go URL Shortener

A URL shortener service built with Go and React, allowing users to create shortened URLs that redirect to the original URL. The backend is powered by Go with an SQLite database for persistent storage, while the frontend is built using React.
Features

    Shorten long URLs into easily shareable links.
    Persist URLs in an SQLite database.
    Frontend built with React to interact with the Go backend.
    Supports Docker for easy deployment.

Prerequisites

Before you begin, ensure you have the following installed:

    Go (v1.22 or higher)
    Node.js (v18 or higher)
    Docker (optional, for Docker-based deployment)
    Git (optional, for cloning the repository)

Installation

  Clone the Repository:

  bash

    git clone https://github.com/VGRobert/go-url-shortener.git
    cd go-url-shortener

Backend Setup:

  Install Go dependencies:

  bash

    go mod download

Initialize the SQLite database:

bash

    go run main.go

Frontend Setup:

Navigate to the frontend directory:

  bash

    cd url-shortener-frontend

Install Node.js dependencies:

  bash

    npm install

Start the React development server:

  bash

    npm start

Usage
Running the Backend

To start the Go backend, navigate to the project root and run:

  bash

    go run main.go

The backend will start on http://localhost:8080.
Running the Frontend

To start the React frontend, navigate to the url-shortener-frontend directory and run:

  bash

    npm start

The frontend will start on http://localhost:3000.

API Endpoints
POST /shorten

    Description: Shortens a given URL.
    Request Body: JSON { "url": "https://example.com" }
    Response: JSON { "short_url": "abc123" }

GET /{shortURL}

    Description: Redirects to the original URL based on the shortened URL.
    Response: Redirects to the original URL.

Running with Docker
Build and Run the Application with Docker Compose

  Build and Start the Containers:

  bash

    docker-compose up --build

  Access the Application:
        Frontend: http://localhost:3000
        Backend API: http://localhost:8080

Stopping the Containers

To stop the Docker containers, press Ctrl + C or run:

  bash

    docker-compose down
