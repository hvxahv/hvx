package graph

import (
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/spf13/viper"
	"hvxahv/internal/accounts"
	"net/http"
)


func Graphql() {
	schema := accounts.Acct()
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
	
	http.Handle("/graphql", h)

	port := fmt.Sprintf(":%s", viper.GetString("graphql_port"))
	http.ListenAndServe(port, nil)
}