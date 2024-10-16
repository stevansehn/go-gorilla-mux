# go-gorilla-mux
A Go web server using Gorilla Mux to handle dynamic routes for books. Features include:

- Retrieving all books
- Getting a specific book by title
- Fetching books by author
- Handling dynamic routes for book titles and page numbers
- Demonstrating Gorilla Mux features such as path prefixes, subrouters, hostname restrictions, scheme limitations, and HTTP method constraints

## Features

- Path Prefixes & Subrouters
- Hostnames & Subdomains
- Schemes
- Methods

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/stevansehn/go-gorilla-mux.git
   ```

2. Navigate to the project directory:
   ```bash
   cd go-gorilla-mux
   ```

3. Install the Gorilla Mux dependency:
   ```bash
   go get -u github.com/gorilla/mux
   ```

4. Run the server:
   ```bash
   go run main.go
   ```

5. Access the API endpoints in your browser or using a tool like curl.

## API Usage

- **GET /books**
  - Description: Retrieve a list of all books.
  - Response: JSON array of book objects.

- **GET /books/{title}**
  - Description: Get a specific book by title.
  - Response: JSON object representing the book.

- **GET /books/author/{author}**
  - Description: Fetch books by author.
  - Response: JSON array of book objects.

- **GET /books/{title}/page/{page}**
  - Description: Handle dynamic routes for book titles and page numbers.
  - Response: JSON object with book title and page number.

## License

[MIT License](LICENSE)
