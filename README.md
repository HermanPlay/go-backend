# Backend API with Gin Framework

Welcome to the repository for my backend API implemented using the Gin framework. This project serves as the foundation for handling HTTP requests and powering the backend of your application. The API is implemented using clean architecture (at least I tried). 

Structure description:
- `cmd` entry point for the application
- `internal`
    - `api/http` folder containing REST API implementation
        - `constant` as name implies, constants  required for proper functioning
        - `middleware` implementations for the api middleware, such as authorization
        - `routes` definitions of the route handlers
        - `server` definition of the main router
        - `util` useful functions, such as token generation and parsing
    - `config` application config related stuff
    - `database` database interfaces related stuff
- `pkg` implementations of the business logic
    - `domain`
        - `dao` models (entities) which are used to communicate with database
        - `dto` models used to retrieve from and give information to the routers
    - `repository` implementation of database communication depending on usecase
    - `service` core business logic depending on use case 

## Installation

Follow these steps to set up and run the backend API on your local machine.

### Prerequisites

Make sure you have the following installed:

- Go (Golang): [Download and Install Go](https://golang.org/dl/)
- Docker: [Download and Install Docker](https://www.docker.com/products/docker-desktop/)

### Installation

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/HermanPlay/go-backend.git
    ```

2. Change to the project directory:

    ```bash
    cd go-backend
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

### Usage

Run the following command to start the server:

    ```bash
    docker compose up --build -d
    ```

The server will be running at `http://localhost:8080`. You can customise the port in the `.env` file.

## API Endpoints

This backend API provides the following group of endpoints:

Authentication:
- `POST /api/auth/register`: Register a user
- `POST /api/auth/login`: Login using credentials and get Bearer Token

User (Requires Bearer Token):
- `GET /api/user`: Retrieve a list of users.
- `GET /api/user/:userid`: Retrieve a specific user by ID.
- `POST /api/user`: Create a new user.
- `PUT /api/user/:userid`: Update a user by ID.
- `DELETE /api/user/:id`: Delete a user by ID.

Make sure to check the `internal/api/http/routes` folder for detailed information on each endpoint and its corresponding handlers.

## Configuration

The configuration for the server is stored in the `.env` file. You can modify the configuration parameters such as the server port, database connection details, etc., according to your needs. Example can be found in `.env.example` file

## Contributing

Feel free to contribute to the project by opening issues or submitting pull requests. Your feedback and contributions are highly appreciated.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
