package examples

//go:generate ../bin/gobuilder -type=User -file=user.go -output=user_builder.go
type User struct {
	ID    int
	Name  string
	Email string
}
