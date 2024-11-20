# Go + Gin Todo App

## Setup

Make sure to install dependencies:

```bash
go install
```

## Development Server

Start the development server on `http://localhost:8080`:

```bash
go run main.go
```

## Production

Build the application for production:

```bash
docker build -t <image-name> .
docker run -d --name <container-name> -p 127.0.0.1:8080:8080 <image-name>
```
