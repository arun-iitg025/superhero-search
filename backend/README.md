# Superhero Search Engine - Backend

## Overview
This backend provides a search engine for superheroes using data from the SuperHero API and OMDb API. It supports:
- Fuzzy searching
- Advanced filtering by name, powers, and movies
- Pagination for large result sets

## Tech Stack
- **Programming Language**: Golang
- **Database**: MongoDB
- **Cache**: Redis
- **Containerization**: Docker

## API Endpoints
### 1. `/search`
- **Method**: GET
- **Query Parameters**:
  - `query` (string): The search term.
- **Response**: JSON list of superheroes and their data.

### 2. `/update`
- **Method**: POST
- **Description**: Fetches new data from the SuperHero API and updates the local database.

## How to Run

### 1. With Docker
1. Build and start the containers:
   ```bash
   docker-compose up -d
