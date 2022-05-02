package rest

import (
	"github.com/gorilla/mux"

	"github.com/eligion/cosmos-sdk/client"
	"github.com/eligion/cosmos-sdk/client/rest"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
}
