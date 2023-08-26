# Go Example Builder

## Command

### 単体実行

```shell
# No build, running program
go run main.go
# Please access http://localhost:8080
```

### ビルドして実行

```shell
# Build main.go, export filename easy-server
go build -o easy-server main.go
# Run app
./easy-server
# Please access http://localhost:8080
```

### Docker でビルド

example `<tagName>` = app:latest

```shell
# Build container image
docker build -t <tagName> .
# Running Container
docker run -p 8080:8080 <tagName>
# Please access http://localhost:8080
```
