package models

import (
	"context"
	"fmt"
	"strconv"

	"github.com/GovtBlockchainAssoc/daoctl/util"
	eos "github.com/eoscanada/eos-go"
	"github.com/ryanuber/columnize"
	"github.com/spf13/viper"
)

// Proposal ...
type Proposal struct {
	ID                uint64
	Approved          bool
	Recipient         eos.Name
	Proposer          eos.Name
	BallotName        eos.Name
	Title             string
	Description       string
	VoteTokenAmount   eos.Asset
	RewardTokenAmount eos.Asset
	CreatedDate       eos.BlockTimestamp
}

func (r *Proposal) String() string {

	output := []string{
		fmt.Sprintf("Proposal ID|%v", strconv.Itoa(int(r.ID))),
		fmt.Sprintf("Recipient|%v", string(r.Recipient)),
		fmt.Sprintf("Proposer|%v", string(r.Proposer)),
		fmt.Sprintf(viper.GetString("VoteTokenSymbol")+"|%v", util.FormatAsset(&r.VoteTokenAmount, 0)),
		fmt.Sprintf(viper.GetString("RewardTokenSymbol")+"|%v", util.FormatAsset(&r.RewardTokenAmount, 0)),
		fmt.Sprintf("Created Date|%v", r.CreatedDate.Time.Format("2006 Jan 02 15:04:05")),
		fmt.Sprintf("Ballot ID|%v", string(r.BallotName)[11:]),
		fmt.Sprint("\n"),
		fmt.Sprintf("Title: |%v", string(r.Title)),
		fmt.Sprintf("Description|%v", r.Description),
	}
	return columnize.SimpleFormat(output)
}

// NewProposal converts a generic DAO Object to a typed Payout
func NewProposal(daoObj DAOObject) Proposal {
	var a Proposal
	a.ID = daoObj.ID
	a.Recipient = daoObj.Names["recipient"]
	a.Proposer = daoObj.Names["owner"]
	a.Title = daoObj.Strings["title"]
	a.BallotName = daoObj.Names["ballot_id"]
	a.RewardTokenAmount = daoObj.Assets["reward_token_amount"]
	a.VoteTokenAmount = daoObj.Assets["vote_token_amount"]
	a.CreatedDate = daoObj.CreatedDate
	return a
}

// NewProposalByID loads a single role based on its ID number
func NewProposalByID(ctx context.Context, api *eos.API, ID uint64) Proposal {
	daoObj := LoadObject(ctx, api, "proposal", ID)
	return NewProposal(daoObj)
}

// Proposals provides the set of active proposals
func Proposals(ctx context.Context, api *eos.API, scope string) []Proposal {
	objects := LoadObjects(ctx, api, scope)
	var proposals []Proposal
	for index := range objects {
		daoObject := ToDAOObject(objects[index])
		proposal := NewProposal(daoObject)
		proposal.Approved = scopeApprovals(scope)
		proposals = append(proposals, proposal)
	}
	return proposals
}
