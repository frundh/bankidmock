# BankID Mock

BankID Mock is a development tool designed to mimic the behavior of the actual BankID service by running an HTTP server.
This mock service is intended for use in development environments, allowing developers to interact with a simulated BankID service without the need for actual communication with the production BankID infrastructure.
The API is designed from the BankID [specification](https://www.bankid.com/utvecklare/guider/teknisk-integrationsguide/graenssnittsbeskrivning).

**Note** that this is currently under active development and is very bare bone and additional features may be added in the future.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
- [Configuration](#configuration)
- [TODO](#todo)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Mock BankID Server:** Simulates the behavior of the BankID service.
- **Development Environment:** Ideal for use in development environments.
- **HTTP Server:** Provides a simple HTTP server for easy integration.

## Getting Started

### Prerequisites

Make sure you have the following prerequisites installed on your system:

- [go](https://go.dev/doc/install) (version 1.21.5 or higher)
- [docker](https://docs.docker.com/get-docker/) (for building container images)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/rojanDinc/bankidmock.git
   ```

2. Navigate to the project directory:

   ```bash
   cd bankidmock
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

### Usage

To start the BankID Mock server, run the following command:

```bash
go run cmd/bankidmock/main.go
```

The server will start on a default port (e.g., `http://localhost:8888`). You can configure the port, see the [configuration](#configuration) section.

Now, you can make HTTP requests to the BankID Mock server.

## Configuration

You can configure the BankID Mock by adding environment variables. Adjust the following parameters as needed:

- `PORT`: The port on which the BankID Mock server will listen (default: `8888`).

**TODO**: Add more configuration options
- [ ] `responseDelay`: Simulated response delay in milliseconds (default: `0`).
- [ ] `logRequests`: Log incoming requests to the console (`true` or `false`, default: `true`). Currently there are some request logging but not really useful, mostly for logging errors.

## TODO
- [ ] Add more endpoints.
- [ ] Add support for certificates.

## Contributing

Contributions are welcome! Feel free to open issues or pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the [MIT License](LICENSE.md).

