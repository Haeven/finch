# Finch

This service manages user interactions for a video platform, including likes, dislikes, and comments. It uses Go for the backend, Bun as the ORM, and Socket.IO for real-time updates.

## Features

- User account management
- Video interaction handling (likes, dislikes, comments)
- Real-time updates using Socket.IO
- PostgreSQL database for data persistence

## Prerequisites

- Docker
- Docker Compose (optional, for easier management)

## Running with Docker

1. Build the Docker image:
   ```
   docker build -t flux-video-interaction .
   ```

2. Run the container:
   ```
   docker run -p 8080:8080 flux-video-interaction
   ```

   The service will be available at `http://localhost:8080`.
