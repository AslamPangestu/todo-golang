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

<!-- TODO -->
## Routes

```
POST   /api/v1/register          --> todo-be/controllers.(*authController).Register
POST   /api/v1/login             --> todo-be/controllers.(*authController).Login
GET    /api/v1/notes             --> todo-be/controllers.(*noteController).GetNotes
POST   /api/v1/notes             --> todo-be/controllers.(*noteController).CreateNote
PATCH  /api/v1/notes/:id         --> todo-be/controllers.(*noteController).UpdateStatusNote
PUT    /api/v1/notes/:id         --> todo-be/controllers.(*noteController).UpdateNote
DELETE /api/v1/notes/:id         --> todo-be/controllers.(*noteController).DeleteNote
```
