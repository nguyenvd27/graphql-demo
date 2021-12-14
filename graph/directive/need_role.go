package directive

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	graphqlmodel "github.com/nguyenvd27/graphql-test/graph/model"
)

type NeedRoleDirective func(ctx context.Context, obj interface{}, next graphql.Resolver, role graphqlmodel.Role) (res interface{}, err error)

func GenerateNeedRoleDirective() NeedRoleDirective {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role graphqlmodel.Role) (res interface{}, err error) {
		fmt.Println("-----context: ", ctx)
		fmt.Println("-----role: ", role)

		if role.String() == "ADMIN" {
			return nil, fmt.Errorf("access denied")
		}

		return next(ctx)
	}
}
