# Blogify

A blogging platform API to share what is in my mind

## Table of content

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/hyprhex/blogify.git
```

2. Navigate to the project

```bash
cd blogify
```

3. Install dependencies

```bash

# Handle routing
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware

# Needed to connect with the database
go get github.com/lib/pq

# To validate user inputs
go get github.com/go-playground/validator/v10

```

**Install CLI tools**

- To migrate file I use : [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- To get env variable : [direnv](https://direnv.net/)
- To reload server on change: [air](https://github.com/air-verse/air)

**Install docker**

- Install docker and docker-compose [Here](https://docs.docker.com/compose/install/)

## Usage

1. Build docker compose

```bash
sudo docker compose up --build
```

2. Run the server

```bash
air
```

## Features

- Straightforward Bloggin platform
- No authentication
- Easy to test

## Contributing

Contributions are welcome! just submit your pull requests and I will handle it.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments

- Inspiration: Thanks to [roadmap](https://roadmap.sh/projects/blogging-platform-api)
- Tools: Project built using Go, PostgreSQL
