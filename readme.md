## Required commands for compile/execution of go code.

# Create module
- go mod init

# Download module
- go get github.com/gin-gonic/gin

# Download all dependency (go.mod)
- go mod download (downloads modules)
- go mod tidy (downloads + removes unused deps)

# Compile project
- go build

# Run without build
- go run main.go / go run .