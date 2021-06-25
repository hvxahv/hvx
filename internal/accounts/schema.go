package accounts

import (
	"github.com/graphql-go/graphql"
	"log"
)

func populate() *accounts {
	name := "hvturingga"

	aq := NewQueryAcctByName(name)
	query, err := aq.Query()
	if err != nil {
		return nil
	}

	return query
}

var acctType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "accounts",
		Fields: graphql.Fields{
			"avatar": &graphql.Field{
				Type: graphql.String,
			},
			"bio": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"telegram": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func Acct() graphql.Schema {

	fields := graphql.Fields{
		"acct": &graphql.Field{
			Type: acctType,
			Description: "Get account data.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				acct := populate()
				return acct, nil
			},
		},
	}

	acctQuery := graphql.ObjectConfig{Name: "acctQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(acctQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}

