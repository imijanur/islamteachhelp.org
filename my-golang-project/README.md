# My Golang Project

This project is a simple web application built with Go. It serves static content and uses templates for rendering HTML pages. The application is structured to separate concerns, making it easy to maintain and extend.

## Project Structure

```
my-golang-project
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handler.go   # HTTP request handlers
│   ├── templates
│   │   └── base.html     # Base HTML template
│   └── models
│       └── model.go      # Data models
├── static
│   ├── css
│   │   └── styles.css    # CSS styles
│   └── js
│       └── scripts.js     # JavaScript code
├── go.mod                # Module definition
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd my-golang-project
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

4. **Access the application:**
   Open your web browser and navigate to `http://localhost:8080`.

## Usage

- The application serves static files from the `static` directory.
- HTML pages are rendered using the templates located in the `internal/templates` directory.
- HTTP request handling is managed in the `internal/handlers` package.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.