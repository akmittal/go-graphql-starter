package person

type PersonResolver struct {
	Person
}

func (p PersonResolver) FirstName() string {
	return p.Person.FirstName
}
func (p PersonResolver) LastName() string {
	return p.Person.LastName
}

func (p PersonResolver) Username() string {
	return p.Person.Username
}
func (p PersonResolver) EmailID() string {
	return p.Person.EmailID
}
