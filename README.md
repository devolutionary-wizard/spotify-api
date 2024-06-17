# Spotify API Explorer

This repository is dedicated to exploring the Spotify API, allowing users to interact with and retrieve data from Spotify's vast library of music and podcasts. The project is built in Go and utilizes the `godotenv` package for environment variable management, ensuring secure API access.

## Getting Started

To get started with this project, follow these steps:

### Prerequisites

- Go (1.15 or later)
- A Spotify Developer account and an API key

# Environment Initialization in Go

This document explains the Go code snippet responsible for initializing environment variables from a `.env` file using the `godotenv` package.

## Code Explanation

The code snippet is part of a Go package named `utils`. It includes the necessary imports and defines a function `Init()` for initializing environment variables.

### Imports

- `log`: Used for logging errors.
- `os`: Used to interact with the operating system, like fetching the current working directory.
- `path/filepath`: Used for manipulating file paths.
- `github.com/joho/godotenv`: A Go package for loading environment variables from a `.env` file.

### Global Variable

- `envPath`: A string variable to store the path to the directory containing the `.env` file.

### The `Init` Function

- The function starts by retrieving the current working directory using `os.Getwd()` and stores the path in `envPath`.
- It then attempts to load the `.env` file located in `envPath` using `godotenv.Load(filepath.Join(envPath, ".env"))`.
- If there is an error loading the `.env` file, the function logs a fatal error message, which includes the error returned by `godotenv.Load`.

This setup is particularly useful in applications that require configuration through environment variables, providing a convenient way to manage sensitive or environment-specific settings without hard-coding them into the source code.
