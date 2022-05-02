package rest

import (
	"github.com/gorilla/mux"

	"github.com/eligion/cosmos-sdk/client"
	"github.com/eligion/cosmos-sdk/client/rest"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
