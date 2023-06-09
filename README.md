# Go Blog API

<img src=".documentation_files/blog-go.png" alt="imagem de uma máquina de escrever com alguns livros ao lado">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

This repository contains the code for a Go API that implements a simple blog with user management, post creation, and commenting functionality.

> **Warning**
> This project is in its early stages, and I'm using it as an opportunity to deepen my knowledge of Go. Therefore, I kindly request that you refrain from using it in a production environment for the time being, as there may be several bugs that could potentially lead to problems.

## Installation

To install and run this project, follow these steps:

1. Ensure you have Go installed (version 1.20 or higher).
2. Clone this repository to your local machine.
3. Open a terminal and navigate to the project's directory.
4. Run the following command to download the project dependencies:

   ```shell
   go mod tidy
   ```
## Dependencies
This project uses the following third-party dependencies:

- github.com/gin-gonic/gin: A web framework for building APIs in Go.
- github.com/golang-jwt/jwt: JSON Web Token implementation for Go.
- golang.org/x/crypto: Go cryptography libraries.
- gorm.io/driver/postgres: PostgreSQL driver for the GORM ORM library.
- gorm.io/driver/sqlite: SQLite driver for the GORM ORM library.
- gorm.io/gorm: A powerful ORM library for Go.

Please refer to the respective documentation for more information on each dependency.

## Usage
To start the Go Blog API, run the following command:

 ```shell
   cd ./cmd/API 
   go run main.go
   ```

By default, the API will be accessible at http://localhost:3333. You can modify the port or other configurations in the main.go file.

## Contributing
Contributions are welcome! If you find any issues or want to enhance the functionality of this project, please submit a pull request. Make sure to follow the established coding guidelines and provide clear commit messages.

## License
This project is licensed under the MIT License. Feel free to use, modify, and distribute the code as per the terms of the license.
