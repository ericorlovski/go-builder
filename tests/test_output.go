
package 

type UserBuilder struct {
	iD int
	name string
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) WithID(v int) *UserBuilder {
	b.iD = v
	return b
}

func (b *UserBuilder) WithName(v string) *UserBuilder {
	b.name = v
	return b
}

func (b *UserBuilder) Build() User {
	return User{
		ID: b.iD,
		Name: b.name,
	}
}
