
1. From Project Root Directory

# Initialize Modules
# In project root (where go.mod is)
```shell
go mod init example.com/myproject
```

# Ensure module dependencies are resolved
```shell
go mod tidy
```

# Run the application directly
```shell
go run app.go
```



2. If You Want to Build an Executable

# Build for current OS
```shell
go build -o myapp app.go
```

# Run the compiled binary
```shell
./myapp
```