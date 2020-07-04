
## View and Query Commands
### Show Help
```
./daoctl -h
```

### View Proposals
```
./daoctl get proposals
./daoctl get proposal 0
```

Example output: 
```
$ ./daoctl get proposals

NETWORK: Connecting to the DAO Test Network

 Proposals  --scope= proposal

 #   Proposer   Recipient    Title      GBAV    GBAR   Created Date   Ballot
--- ---------- ----------- ---------- -------- ------ -------------- --------
 0    zzz.gba     zzz.gba   Test 1     10,000     --    2020 Jul 03       14
 2    zzz.gba     zzz.gba   Test 2         10      8    2020 Jul 03       1b
--- ---------- ----------- ---------- -------- ------ -------------- --------
                            Subtotal   10,010      8
```
Currently open proposals use the default scope. For passed or failed proposals, use ```--scope passedprops``` or ```--scope failedprops```.

### View Details of a Ballot
See ```Ballot``` column above to get the ballot ID.
```
./daoctl get ballot 14
```

### Query Blockchain History
```
./daoctl query
./daoctl query --contract dao.gba --action enroll 
```
### View Members
```
./daoctl get members
./daoctl get member doris.gba  
```
Member history and voting history will soon be included on the ```get member``` command.

## Wallet/Signing Functionality
On the Telos test network, you can use one of our testing accounts. They share the following key pair: 
```
Private key: 5JA44kUjURbkwMMR5RMaNQrvayJ6jfvztocEEhfsDj1ECi4JL5m
Public key: EOS5ePx2Up6jaFcgDNTbm84cGQjZ7yP12rUwsEiqaqwMxZN61124p
```
#### Testnet DAO Member Accounts (please use for expirementation)
- alice.gba
- bob.gba
- jackson.gba
- doris.gba 

### Step 1: Create a Configuration
The ```daoctl``` configuration is yaml. The default behavior is to connect to the test net for now. Save the below template as a file named "daoctl.yaml" in the root directory (or passed with ```--config```). 
``` yaml
# endpoint for the blockchain
EosioEndpoint: https://test.telos.kitchen

# endpoint for Hyperion, which feeds history records for the 'query' command
HyperionEndpoint: https://testnet.telosusa.io/v2

# display option - assets can either be viewed as asset objects (no commas, with symbol) or just floats (commas, no symbol)
AssetsAsFloat: true

# account name holding the DAO contract
DAOContract: dao.gba

# your account name (you will also need a vault file)
DAOUser: doris.gba  # pretend user on testnet

# reward token
RewardTokenContract: token.gba
RewardTokenSymbol: GBAR

# telos decide parameters
BallotPrefix: gba........
TelosDecideContract: telos.decide
VoteTokenSymbol: GBAV
```

### Step 2: Create a Vault
```
./daoctl vault create add --import
```
Paste in your private key, and secure it with a password. Your private key is encrypted with this password, and you will enter it each time you sign a transaction. (You can use blank <Enter> or just a short PIN for low security.)

See ```eosc``` docs for more info: https://github.com/eoscanada/eosc#eosc-vault-management

Now you are ready to sign transactions with ```daoctl```.

## Create a Proposal
Paste the following into a file named "proposal.json".
```javascript
{
    "scope": "proposal",
    "names": [
        {
            "key": "owner",
            "value": "doris.gba"
        },
        {
            "key": "recipient",
            "value": "doris.gba"
        }
    ],
    "strings": [
        {
            "key": "title",
            "value": "Test Proposal"
        },
        {
            "key": "description",
            "value": "GBA Proposal Test"
        }
    ],
    "assets": [
        {
            "key": "vote_token_amount",
            "value": "5.00 GBAV"
        },
        {
            "key": "reward_token_amount",
            "value": "6.50 GBAR"
        }
    ],
    "ints": [],
    "time_points": [],
    "trxs": []
}
```
See https://raw.githubusercontent.com/GovtBlockchainAssoc/daoctl/master/samples/proposal1.json

Create the proposal on chain
```
./daoctl create -f proposal.json
```
When successful, re-run ```./daoctl get proposals``` to see your new proposal.

#### Vote on a Proposal
Currently, on the testnet, the voting period is set to only 60 seconds.  This is to allow quick testing, so vote fast! The default voting period on the main net is 1 week.
``` bash
./daoctl vote 14 pass
```

### Close a Proposal
``` 
./daoctl close 14
```
If the proposal passes, the token quantities in the proposal are paid. 

## Current daoctl Help
For most up-to-date help, run ```./daoctl -h```.
```
Decentralized Autonomous Organization (DAO) control CLI.

Example use:
	daoctl get proposals --scope passedprops
	daoctl get proposal 34

Usage:
  daoctl [command]

Available Commands:
  backup      creates a local backup of current DAO data
  close       close a proposal
  create      create an object based on the JSON file (use this to create new proposals)
  get         retrieve and display objects from the DAO on-chain smart contract
  help        Help about any command
  query       Query action history on the DAO for specific users
  vault       The vault is a secure key store (wallet). Your key is stored encrypted by the passphrase.
  vote        vote pass or fail on a ballot

Flags:
  -f, --assets-as-floats    Format assets objects as floats (helpful for CSV export)
      --config string       config file (default is ./daoctl.yaml)
      --debug               Enables verbose debug messages
      --expiration int      Set time before transaction expires, in seconds. Defaults to 30 seconds. (default 30)
  -h, --help                help for daoctl
      --vault-file string   Wallet file that contains encrypted key material (default "./eosc-vault.json")

Use "daoctl [command] --help" for more information about a command.
```