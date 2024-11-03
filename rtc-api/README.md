# RTC-API

RTC-API is a lightweight Go API server built with Fiber that provides commenting functionality for the Right to Comment system. It serves as the backend API for managing comments associated with YouTube videos.

## Project Structure

```
RTC-API/
│   go.mod            # Go module file
│   go.sum            # Go module checksum
│   main.go           # Main application file
│
└───database/         # Database operations
        database.go   # Database initialization and queries
```

## Features

- RESTful API endpoints for comment management
- SQLite database integration
- Timestamp formatting for comments

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (1.16 or later)
- SQLite3

## Installation

1. Clone the repository:
```bash
git clone https://github.com/TanishkBansode/rtc.git
cd rtc/rtc-api
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the API Server

Start the server with:
```bash
go run main.go
```

The server will start on `http://localhost:3000`

## API Endpoints

### Root Endpoint, just to test if it works
```
GET /
Response: "Hello, World!"
```

### Get Comments
```
GET /comments/:videoId
Response: JSON array of comments
```

Example response:
```json
[
  {
    "text": "Comment text here",
    "createdAt": "2 Jan 2024"
  }
]
```

### Add Comment
```
POST /comments/:videoId
Body: form-data
  - comment: "Your comment text"
Response: Updated list of comments for the video
```

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber) - Fast HTTP framework
- SQLite3 - Database engine

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- 200: Successful operation
- 500: Internal server error with error message

Example error response:
```json
{
  "error": "Failed to load comments"
}
```

## Development

The API server is configured to:
- Run on port 3000
- Use SQLite database stored in "../data"
- Format dates in "2 Jan 2006" format
- Return JSON responses

## Testing

To test the API endpoints:

1. Start the server
2. Use curl or Postman(or Thunderclient!) to make requests:

```bash
# Get comments
curl http://localhost:3000/comments/VIDEO_ID

# Add comment
curl -X POST -F "comment=Your comment" http://localhost:3000/comments/VIDEO_ID
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
