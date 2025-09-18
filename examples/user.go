package examples

//go:generate ../bin/gobuilder -type=User -file=user.go -output=user_builder.go
type User struct {
	ID    int            `required:"true" default:"1"`
	Name  string         `omitempty:"true"`
	Tags  []string       `omitempty:"true"`
	Data  map[string]int `omitempty:"true"`
	Ref   *int           `omitempty:"true"`
	Email string
}
