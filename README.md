<<<<<<< HEAD
# AMI Client

## Overview
The `ami-client` project is a Go-based application. This repository contains the source code for the `amin.go` file, along with the `go.mod` file for dependency management.

## Prerequisites
To run this project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or later)

## Project Structure
```
ami-client/
├── amin.go       # Main Go source code file
├── go.mod        # Go module file for dependency management
├── logs.txt      # Log file for application output
```

## Getting Started

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ami-client
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run amin.go
   ```

## Logs
The application generates logs in the `logs.txt` file. You can check this file for detailed information about the application's execution.

## Contributing
If you'd like to contribute to this project, feel free to fork the repository and submit a pull request. Please ensure your code follows the Go coding standards and includes appropriate tests.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
=======
# Go AMI Event Listener

A Go-based client for the **Asterisk Manager Interface (AMI)**. This tool connects to your Asterisk PBX, authenticates, and streams real-time events to both your console and a persistent log file.


## Prerequisites

* **Go**: Version 1.18 or higher.
* **Asterisk**: An active server with AMI enabled in `manager.conf`.
* **Permissions**: Ensure your AMI user has `read` privileges for the events you want to capture.

## Installation

1. **Clone the repository**:
   ```bash
   git clone [https://github.com/yourusername/your-repo-name.git](https://github.com/yourusername/your-repo-name.git)
   cd your-repo-name
>>>>>>> bf0e6d74fa0dc5d02fb92b3c1a96b329368675b9
