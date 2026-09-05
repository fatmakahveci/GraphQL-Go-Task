# GraphQL Full-Stack Task

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![React](https://img.shields.io/badge/React-18-149ECA?logo=react&logoColor=white)](https://react.dev/)
[![GraphQL](https://img.shields.io/badge/GraphQL-gqlgen-E10098?logo=graphql&logoColor=white)](https://graphql.org/)
[![Last commit](https://img.shields.io/github/last-commit/fatmakahveci/GraphQL-Go-Task)](https://github.com/fatmakahveci/GraphQL-Go-Task/commits/main)
[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](LICENSE.md)

A compact full-stack exercise combining a Go GraphQL API with a React client. It demonstrates schema-driven resolvers, GraphQL AST inspection, enum handling, API integration, and interaction tests.

## Features

- Serves a typed GraphQL endpoint with `gqlgen`
- Resolves a collection of Star Wars characters through the `heroes` query
- Inspects schema object types and enum values with `graphql-go-tools`
- Exposes schema introspection for development tooling
- Fetches and renders API results in a React interface
- Tests backend resolvers, schema parsing, data fetching, and UI interaction

## Architecture

```text
.
├── graph/               GraphQL schema, generated code, models, and resolvers
├── schemaparser/        GraphQL AST traversal helpers
├── frontend/            Create React App client and Jest tests
├── server.go            HTTP server on port 8085
└── server_test.go       API-level backend tests
```

The backend accepts GraphQL POST requests at `http://localhost:8085/query`. The frontend runs at `http://localhost:3000`, which is the origin allowed by the server's CORS configuration.

## Requirements

- Go 1.25 or newer
- Node.js 20 or newer
- npm

## Run Locally

Start the API:

```bash
git clone https://github.com/fatmakahveci/GraphQL-Go-Task.git
cd GraphQL-Go-Task
go mod download
go run server.go
```

In a second terminal, start the frontend:

```bash
cd GraphQL-Go-Task/frontend
npm ci
npm start
```

Example request:

```bash
curl --request POST http://localhost:8085/query \
  --header 'Content-Type: application/json' \
  --data '{"query":"{ heroes { name } }"}'
```

## Testing

Pull requests run separate backend and frontend quality jobs. Backend HTTP tests
use an isolated ephemeral server instead of racing a background process on a fixed
port. CI verifies modules, runs `go vet`, race-enabled tests, and builds both stacks.

```bash
# Repository root: backend
go test ./...

# frontend/: client
npm test -- --watchAll=false
npm run build
```

## GraphQL Surface

- `heroes` returns character names and character-specific fields.
- `types` returns GraphQL object type names without interface definitions.
- `Episode` and `Side` exercise enum parsing.

See [`graph/schema.graphqls`](graph/schema.graphqls) for the authoritative schema.

## Contributing

Read the [contributing guide](.github/CONTRIBUTING.md) and add or update tests with every behavioral change.

## Project Resources

- [Changelog](CHANGELOG.md)
- [Security policy](.github/SECURITY.md)
- [License](LICENSE.md)
