# Chuck Norris Joke Microservice 🤠

Welcome to **Chuck Norris Joke Microservice** - this microservice fetches random Chuck Norris jokes and delivers them fast (like Chuck Norris).

## 💻 How It Works

This microservice provides an API with two endpoints:

- **`GET /status`**: Check if the service is up. It returns a humble "OK" response.
- **`GET /chuckNorrisJoke`**: Get a random Chuck Norris joke. Because why not?

The joke data is fetched from the [Chuck Norris Jokes API](https://api.chucknorris.io/).

## 🏗 Project Structure

```bash
.
├── api.go              # API server with HTTP handlers
├── client              # HTTP client
│   └── client.go
├── dto                 # Data transfer objects (DTOs) for Chuck Norris joke
│   └── joke.go
├── logging.go          # Service decorator for logging
├── main.go             # Entry point for the application
├── service.go          # Core service that fetches Chuck Norris jokes
├── tests               # Unit tests to make sure Chuck Norris approves the code
│   └── joke_test.go
├── go.mod              # Go module file
├── Makefile            # Automates building and running the service
└── README.md           # The document you're reading
```

## 🛠️ How to Run It

Make sure you have [Go](https://golang.org/) installed on your machine. Then, you can run this service in a few simple steps.

```bash
git clone https://github.com/M-Komorek/joke_norris.git
cd joke_norris
make build
./bin/joke_norris
```

The service will start on port `4200` by default, but you can change this in the `main.go`.

## 🧪 Running Tests

```bash
go test ./tests/...
```

## 🚀 API Endpoints

Here are the two main endpoints:

### 1. Check Service Status

```bash
GET /status
```

Response:

```json
OK
```

### 2. Get a Random Chuck Norris Joke

```bash
GET /chuckNorrisJoke
```

Sample Response:

```json
{
    "id": "123abc",
    "value": "Chuck Norris doesn’t read books. He stares them down until he gets the information he wants."
}
```

## 🧙 Configuration

If you want to change the source URL for Chuck Norris jokes or modify other parameters, you can do so in the `service.go` file. Right now, the jokes are pulled from the [Chuck Norris Jokes API](https://api.chucknorris.io/), but you can modify that if you want to.

