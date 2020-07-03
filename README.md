

### View Proposals
```
./daoctl get proposals
```
#### Vote on a Proposal
Voting requires signing a transaction, so the vault-file parameters is required.
```
./daoctl --config daoctl-test.yaml --vault-file ~/dev/gba/gba-vault-test.json vote 14 pass
```
### View Details of a Ballot
```
./daoctl get ballot d4
```

With much credit and appreciation to ```eosc``` at https://github.com/eoscanada/eosc
