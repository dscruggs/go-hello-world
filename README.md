# Hello World Go Project

This is a basic repo to help with setting up and running a new Go project. Includes package/dependency management, importing an external package, and building and running a simple command-line application. 

I have a few short notes and resources on go in my [devnotes](https://github.com/dscruggs/devnotes) repo [here](https://github.com/dscruggs/devnotes/blob/main/quickref/go.md).

`go mod init <module_name>` to set up a new go module

## Project Structure

```
hello-world-go/
├── go.mod                # Go module file
├── go.sum                # Dependency checksum file
├── internal/
│   └── greeter/
│       └── greeter.go    # Greeter package with HelloWorld logic
├── main.go               # Main entry point
```

### Key Files
1. **`main.go`**:
   - The entry point for the application.
   - Directly calls the `HelloWorld` function from the `greeter` package.

2. **`internal/greeter/greeter.go`**:
   - A simple package for printing a "Hello, World!" message with color.

3. **`go.mod`**:
   - Specifies the module name and dependencies.

## Setup Instructions

1. **Initialize and Download Dependencies**:
   If you cloned the project, dependencies will already be listed in `go.mod`. Run:
   ```bash
   go mod tidy
   ```

3. **Run the Application**:
   ```bash
   go run main.go
   ```

## Running the Application

### Using `go run`
Run the project directly:
```bash
go run main.go
```

### Building an Executable
To build a standalone executable:
```bash
go build -o hello
```

Run the executable (on linux):
```bash
./hello
```



