# Right to Comment System

A website that lets you search and watch YouTube videos while adding comments(right-to-comment) made with Gin, and a simple API service that handles just the commenting part (rtc-api) made with fiber - both sharing the same database.

## Project Structure

```
RTC/
├───right-to-comment/    # Web Application (Gin)
│                       # Complete website with YouTube integration
└───rtc-api/            # REST API Service (Fiber)
│                       # Standalone comments API
└───data/               # Shared SQLite database
```

Technology Stack
Web Application (right-to-comment)

Framework: Gin Web Framework
Features:

Complete web interface with HTML templates
YouTube video search and embedding
YouTube Data API v3 integration
Comment system integration
Server-side rendered pages


Dependencies:

godotenv for configuration
YouTube API client
HTML templates
Static asset serving



API Service (rtc-api)

Framework: Fiber v2
Features:

RESTful JSON API endpoints
Comment CRUD operations
Lightweight and fast


Endpoints:

GET /comments/:videoId - Retrieve comments
POST /comments/:videoId - Add new comment



Shared Components

Language: Go 1.16+
Database: SQLite3 (shared in root /data directory)
Version Control: Git

Service Communication

Web Application (:8080):

Serves HTML pages
Handles user interactions
Manages YouTube video integration
Stores/retrieves comments from shared database


API Service (:3000):

Provides RESTful JSON endpoints
Handles comment operations through API calls
Accesses same database as web application



Development Setup

Configure the web application:
bashCopycd right-to-comment
# Add your YouTube API key to .env

Start both services:
bashCopy# Terminal 1 - Web Application
cd right-to-comment
go run main.go

# Terminal 2 - API Service
cd rtc-api
go run main.go

Access the services:

Web Application: http://localhost:8080
API Endpoints: http://localhost:3000



API Testing
Test the API service using curl:
bashCopy# Get comments
curl http://localhost:3000/comments/VIDEO_ID

# Add comment
curl -X POST -F "comment=Your comment" http://localhost:3000/comments/VIDEO_ID


License
MIT License
