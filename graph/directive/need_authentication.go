package directive

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/99designs/gqlgen/graphql"
)

type NeedAuthenticationDirective func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

func GenerateNeedAuthenticationDirective() NeedAuthenticationDirective {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		fmt.Println("-----context: ", ctx)
		if rand.Int()%2 == 0 {
			return nil, fmt.Errorf("access denied")
		}

		return next(ctx)
	}
}
