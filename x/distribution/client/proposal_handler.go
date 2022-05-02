package client

import (
	"github.com/eligion/cosmos-sdk/x/distribution/client/cli"
	"github.com/eligion/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/eligion/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
