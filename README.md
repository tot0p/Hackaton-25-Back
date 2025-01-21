# Hackaton-25-Back

## Setup

1. Install Go or Docker
2. Clone the repository
3. Create a `.env` file in the root of the project with the following content:
```
PORT=8080 
DB_FILE_PATH=./db.db
URL_OF_API_CO2=http://localhost:8000
```
4. Clone the repository of the API CO2 [here](https://github.com/Luxchar/api_hackaton) and configure the URL in the `.env` file


## How to run the project

1. Clone the repository
2. Run `go mod download` to download all the dependencies
3. Run `go run main.go` to start the server

## How to run DockerImage

1. Run `docker build -t hackaton-25-back .` to build the image
2. Run `docker run -p 8080:8080 hackaton-25-back` to start the container
3. Open your browser and go to `http://localhost:8080/ping` to see the server running

## API Doc

[API Doc](Docs/ApiDoc.md)


## Authors

- [Thomas Lemaitre](https://github.com/tot0p)