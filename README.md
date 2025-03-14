## Infinite Scroll Demo

This is an application built with LeapKit to demonstrate infinite scrolling using Go, HTMX, and Tailwind CSS.

### Getting Started

#### Setup tools

To get started, ensure you have Go 1.24 installed. Once Go is installed, download the required dependencies:

```sh
go mod download && go tool tailo download
```

#### Running the application

To run the application in development mode, use the following command:

```sh
go tool dev
```

Once the application is running, you can access it at [http://localhost:3000](http://localhost:3000).

### Infinite Scroll Feature

This demo implements infinite scrolling using HTMX and Go. The server dynamically loads additional content when the user reaches the end of the page.

#### How It Works
- The frontend uses HTMX to make asynchronous requests when scrolling down.
- The backend in Go handles paginated requests and serves HTML snippets dynamically.
- Tailwind CSS is used for styling and layout.