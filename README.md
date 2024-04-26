<img width="1303" alt="Screenshot 2024-04-26 at 9 19 48 AM" src="https://github.com/Prince-1501/url-shortner/assets/37762770/a97137c1-c4b9-4266-843b-b98e0edac128">

# URL Shortener

A simple URL shortening service built in Go.

## Overview

This project provides a basic URL-shortening service implemented in Go. It allows users to shorten long URLs into more manageable and shareable links. The service also includes a redirect feature to redirect users from the shortened URL to the original long URL.

## Installation

To use the URL shortener, you need to have Go installed on your system. You can download and install Go from the [official Go website](https://go.dev/).

Clone the repository to your local machine:

```sh
git clone https://github.com/Prince-1501/url-shortner.git
```

## Usage

### Running the Server

Navigate to the project directory and run the following command to start the server:

```sh
go run main.go
```

The server will start listening on port `8080` by default. You can change the port in the `main.go` file if needed.

### Shortening a URL

To shorten a URL, send a POST request to the `/shorten` endpoint with a JSON payload containing the original URL:

```sh
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
```

The response will contain a JSON object with the shortened URL:

```json
{
    "short_url": "https://www.linkedin.com/in/iamprince/"
}
```

### Redirecting to the Original URL

To redirect to the original URL, visit the shortened URL in your browser or send a GET request to the `/redirect/{id}` endpoint, where `{id}` is the shortened URL ID:

```sh
curl http://localhost:8080/redirect/abcdef
```

This will redirect you to the original URL associated with the shortened URL.

## Contributing

Contributions are welcome! Feel free to fork the repository and submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
