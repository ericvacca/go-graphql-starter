package schema

import (
	//"../conf"
	//"../model"
	"../service"
	"golang.org/x/net/context"
)

var userQuery = `
	user(email: String!): User
`

func (r *Resolver) User(ctx context.Context,args struct {
	Email string
}) (*userResolver, error) {
	user, err := service.UserService.FindByEmail(args.Email)
	if err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
