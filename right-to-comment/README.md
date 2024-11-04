# Right to Comment

Right to Comment is a Go web application that allows users to search YouTube videos and add comments on our platform. It provides a simple interface for video search, embedding, and commenting functionality.

[![Watch the Demo](https://img.youtube.com/vi/f3N9m0Suhf8/maxresdefault.jpg)](https://youtu.be/f3N9m0Suhf8)

## Features

- YouTube video search using the YouTube Data API
- Video embedding and playback
- Commenting system

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (1.16 or later)
- SQLite3
- A YouTube Data API key

## Project Structure

```
RIGHT-TO-COMMENT/
│   .env                # Environment variables configuration
│   go.mod             # Go module file
│   go.sum             # Go module checksum
│   main.go            # Main application file
│   README.md          # Project documentation
│
├───database/          # Database operations
│       database.go    # Database initialization and queries
│
├───static/           # Static assets
│       logo.png      # Application logo
│
└───templates/        # HTML templates
        embed.html    # Video embedding template
        index.html    # Homepage template
        results.html  # Search results template
```

## Setup

1. Clone the repository:
```bash
git clone https://github.com/YourUsername/rtc.git
cd rtc/right-to-comment
```

2. Create a `.env` file in the root directory with your YouTube API key:
```
YOUTUBE_API_KEY=YOUTUBE_API_KEY
```

3. Install dependencies:
```bash
go mod tidy
```

## Running the Application

1. Start the server:
```bash
go run main.go
```

2. Open your web browser and navigate to:
```
http://localhost:8080
```

## Dependencies

- [Gin Web Framework](https://github.com/gin-gonic/gin) - Web framework
- [godotenv](https://github.com/joho/godotenv) - Environment variable loading
- [YouTube Data API v3](https://developers.google.com/youtube/v3) - YouTube video search and metadata

## How It Works:

### Video Search
- Uses YouTube Data API to search for videos
- Returns top 10 results with title, channel, and duration
- Formats video duration in human-readable format (H:MM:SS)

### Video Embedding
- Embeds YouTube videos using the official iframe player
- Provides a clean interface for video playback

### Commenting System
- Local SQLite database for storing comments
- Comment display with timestamps
- Real-time comment updates

## API Endpoints

- `GET /` - Homepage with search interface
- `POST /search` - Handle video search
- `GET /embed/:id` - Embed video by ID
- `GET /comments/:videoId` - Get comments for a video
- `POST /comments/:videoId` - Add a new comment to a video

## Environment Variables

| Variable | Description |
|----------|-------------|
| YOUTUBE_API_KEY | Your YouTube Data API key |

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
