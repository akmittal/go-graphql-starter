package schema

import (
	"context"
	"go-graphql-starter/db"
	"go-graphql-starter/person"
)

type ViewerResolver struct {
	viewer person.Person
}

func (resolver *ViewerResolver) User(ctx context.Context, args struct {
	ID string
}) (person.PersonResolver, error) {
	result, err := person.GetByUsername(args.ID, db.GetConnection())
	return person.PersonResolver{result}, err
}
