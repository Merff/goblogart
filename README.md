## This is an application for learn golang development

### Used:
- GIN for routing
- godotenv for configuration
- gorm/postgres for database
- testify for tests

### Functionality:
- user can register
- can login via Authorization cookie (jwt token)
- can logout
- can create a post linked to the current user
- can do other CRUD operations with posts

### Database migrations:
- run by `go run migrations/migrations.go`
- for tests `GO_ENV=test go run migrations/migrations.go`

### Tests:
run by `GO_ENV=test go test ./test/controllers -v`
