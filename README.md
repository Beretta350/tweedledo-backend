# tweedledo-backend

This To-Do List application backend is built in the Go programming language (Golang) and adheres to the principles of Hexagonal Architecture, SOLID, and Clean Architecture. It offers a robust and maintainable foundation for managing tasks and to-do lists.

## Key Technologies and Frameworks

- **GORM (Golang Object-Relational Mapping):** GORM is used to simplify database interactions and ensure efficient communication between the application and the database. It allows for seamless integration with various database systems.

- **GIN (Web Framework):** The application leverages the Gin web framework to handle HTTP requests, routing, and middleware, ensuring fast and scalable API development.

- **GOMOCK (Mocking Framework):** GOMOCK is employed for unit testing. It enables the creation of mock objects and behaviors, facilitating isolated testing of the application components.

- **Docker:** Docker is used to containerize the application, making it easy to manage dependencies and ensure consistency across different environments.

- **PostgreSQL:** The application relies on a PostgreSQL as local database running in a container, providing a reliable and scalable data storage solution.

- **In-Memory Database:** For integration testing, an in-memory SQLite database is utilized. This allows for thorough testing without the need for a persistent database connection.

## Architecture

The application is designed following the principles of Hexagonal Architecture. This architecture emphasizes the separation of concerns, enabling the core business logic to be decoupled from external components, such as databases and web frameworks. The application is structured to support various access methods, making it generic and adaptable to different ways of calling, including API endpoints.

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
- `GET /tasklist/{id}`: Retrieve a list of tasks.
- `GET /task/{id}`: Retrieve a specific task by ID.
- `POST /tasklist`: Create a new task list.
- `POST /task`: Create a new task.
- `PUT /tasks/{id}`: Update an existing task.
- `DELETE /tasks/{id}`: Delete a task.

**OBS: Tasks and Task Lists IDs are UUIDs**

## Testing

The application is covered by both unit tests and integration tests. Unit tests aim to cover every line of code to ensure a good level of quality, while integration tests check that the application functions correctly as a whole.

Run unit tests and generate coverage:
    - ```./run-tests.bat``` (for windows)
    - ```./run-tests.sh``` (for linux)

**OBS: Integration tests must be run manually**

## Contributions

We welcome contributions to this To-Do List application backend. If you're interested in enhancing or extending its functionality, feel free to create pull requests or open issues on the repository.

Enjoy using this flexible To-Do List backend built with Golang and Hexagonal Architecture!

<img src="tweedledo-logo.png" width="100" height="100">
