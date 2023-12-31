<img src="tweedledo-logo.png" width="150" height="145">

# tweedledo-backend

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/Postgres-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)

This is To-Do List application backend built in Go programming language (Golang) using the principles of clean architecture, SOLID, and hexagonal architecture. It offers a simple, robust and maintainable foundation for managing tasks and to-do lists.

## Key Technologies and Frameworks

- **GORM (Golang Object-Relational Mapping):** GORM is used to simplify database interactions and ensure efficient communication between the application and the database. It allows for seamless integration with various database systems.

- **GIN (Web Framework):** The application leverages the Gin web framework to handle HTTP requests, routing, and middleware, ensuring fast and scalable API development.

- **GOMOCK (Mocking Framework):** GOMOCK is employed for unit testing. It enables the creation of mock objects and behaviors, facilitating isolated testing of the application components.

- **Docker:** Docker is used to containerize the application, making it easy to manage dependencies and ensure consistency across different environments.

- **PostgreSQL:** The application relies on a PostgreSQL as local database running in a container, providing a reliable and scalable data storage solution.

- **In-Memory Database:** For integration testing, an in-memory SQLite database is utilized. This allows for thorough testing without the need for a persistent database connection.

## Architecture

The application is designed following principles of clean and hexagonal architectures. This architectures emphasizes the separation of concerns, enabling the core business logic to be decoupled from external components, such as databases and web frameworks. The application is structured to support various access methods, making it generic and adaptable to different ways of calling, including API endpoints.

## Get Started

To get started with the To-Do List application backend, follow these steps:

1. Clone the repository to your local machine.

2. Ensure you have Golang and Docker installed.

3. Build and run the application using Docker Compose. 
    - ```docker compose up -d --build```

4. Access the application's API endpoints from **localhost:8080** to interact with the To-Do List.

## Usage

Use the API endpoints to manage your tasks and to-do lists effectively. The Hexagonal Architecture ensures flexibility in how you interact with the application, making it suitable for a variety of use cases.

## Example API Endpoints
- `GET /tasklist`: Retrieve all list of tasks.
- `GET /tasklist/{id}`: Retrieve a list of tasks by ID.
- `POST /tasklist`: Create a new task list.
- `PUT /tasklist/{id}`: Update an tasklist.
- `DELETE /tasklist/{id}`: Delete a tasklist from ID.
- `GET /task/{id}`: Retrieve task by ID.
- `POST /task`: Create a new task.
- `PUT /task/{id}`: Update an existing task.
- `DELETE /task/{id}`: Delete a task from ID.

**OBS: Tasks and Task Lists IDs are UUIDs**

## Testing

The application is covered by both unit tests and integration tests. Unit tests aim to cover every line of code to ensure a good level of quality, while integration tests check that the application functions correctly as a whole.

Run unit and integration tests with coverage:
- ```./run-tests.bat``` (for windows)
- ```./run-tests.sh``` (for linux)

## Contributions

We welcome contributions to this To-Do List application backend. If you're interested in enhancing or extending its functionality, feel free to create pull requests or open issues on the repository.

Enjoy using this flexible To-Do List backend built with Golang and Hexagonal Architecture!
