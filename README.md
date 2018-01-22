# `spec-cli`

This project was bootstrapped using `spec-cli`.

## Usage

Available scripts:

```bash
# Build your project's static assets
spec-cli build

# Run a local server for development
spec-cli dev

# Develop your server locally
spec-cli dev:server

# Run the Jest Test-Runner
spec-cli test

# Start a production server, make sure to run spec-cli build first!
spec-cli start
```

## Folder Structure

You should have a project scaffold with a folder structure similar to the following:

```
.
├── README.md
├── config                   # Config directory for Node.js
│   ├── default.js
│   └── production.js
├── .gitignore
├── public
└── src
    ├── client               # Client-specific code
    │   ├── components
    │   │   └── App.js
    │   ├── index.js
    │   └── styles
    │       └── styles.scss
    └── server               # Server-specific code
        ├── index.js
        └── server.js
```
