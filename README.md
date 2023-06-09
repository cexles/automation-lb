# Automation-lb
<div>

<a href="https://go.dev/doc/go1.20">![go-version](https://img.shields.io/badge/Go-1.20-414141?style=for-the-badge&logo=Go)</a>
<a href="https://docs.chain.link/chainlink-automation/introduction">![chainlink-automation](https://img.shields.io/badge/Chainlink-Automation-375BD2?style=for-the-badge&logo=Chainlink)</a>

</div>

<div>
  
<a href="">![license](https://img.shields.io/github/license/cexles/automation-lb)</a>

</div>

## Documentation

<a href="https://docs.cexles.finance/">![Documentation](https://img.shields.io/badge/GitBook-GitBook-3884FF?style=for-the-badge&logo=GitBook)</a>

### Chainlink Automation scaling &amp; loadbalancing 
![chainlink-view.png](docs/chainlink-view.png)

![sample-log.png](docs/sample-log.png)


## Motivation
_Nobody likes out-of-gas._

In [Cexles](https://cexles.finance) we need to scale [Chainlink Automation](https://chain.link/automation) to perform a huge number of on-chain operations in a single block and distribute their execution between multiple Chainlink Nodes without manual user intervention. 
This software will allow us to unlock the full potential of Chainlink Automation. 

## Current stage
We started an initial development of MVP with smart contracts deployed on Goerli Ethereum testnet.
MVP PoC version. Needs refactor & minor bug fixing.

## Desired result
Hassle-free automated scaling for any Chainlink Automation compatible contract. 

## TODO
- [x] Initial research
- [x] Initial project skeleton
- [x] Sample test smart contract
- [ ] Multiple contracts support
- [x] Automated upkeeps creating
- [ ] Upkeeps load metrics
- [x] Automated on-condition upkeeps scaling
- [ ] Upkeeps scailing prediction
- [x] Automated Upkeep's balance top-up
- [ ] API
- [ ] Dashboard
- [ ] Tests
- [ ] Reference contract implementation
- [x] [Public Docs](https://docs.cexles.finance)

## Setup config 
- Create and fill _config.yaml_ file
  ```shell
  cp ./goerli.example.config.yaml ./config.yaml
  ```
- Config
  ```
    account:
    # Your private key with Link balance (https://faucets.chain.link/ free testnet faucet)
    private_key: ""
  network:
    rpc_url: "wss://"   # Node websocket (Quicknode best choice)
    chain_id: 5 # goerli chain_id, change to needed
  contracts:
    sample_test: # key of your contract
      name: "Sample test contract" # contract name, used just for logging
      short_name: "sampletest" # short contract name, used for creating new upkeep
      address: "0xCb1Ad0A9D12993CB614533Cd19e7f9f555F13816" # sepolia
      version: "0_1" # version of contract, used for creating new upkeep
      gas_limit: 5000000 # upkeep's gasLimit used for creating new upkeep
      execution_limit: 5 # max upkeep execution limit, check Scale Policy docs section https://docs.cexles.finance/architecture/scale-policy
  chainlink:
    registry_version: "2_0"
    topup_amount: "5000000000000000000" # amount in WEI of LINK to topup on Keepup creation
    upkeep_controller_address: "0x09970ef6d0E46F71E568538fD88B44127739D892" # upkeep controller, deploy your own
    token_address: "0x779877A7B0D9E8603169DdbD7836e478b4624789" # LINK token address, check Chainlink docs
    min_controller_approve: "300000000000000000000" # lower allowance threshold in WEI to execute IncreaseAllowance
  app:
    log_level: -1   # Log level from -1 (more logs) to 5 (fewer logs)
    active_contract_key: "sample_test" # put the key from contracts config section
  ```

## Installing
- Install Docker and Docker Compose
- Copy config file
  ```shell 
  cp goerli.example.config.yaml config.yaml
    ```
- Set config values according to docs
- Run build
  ```shell
  docker-compose build
  ```
- Run balancer
  ```shell
  docker-compose up -d
  ```
- Check program logs. If everything is fine you will see something similar to:
  ```shell
  INF upkeep controller found address=0x3911d7499d72dd5f4ea88270af6dfeced1a1bef6
  INF logic contract found name="Cexles platform contract" short_name=cexles_platform version=1_0
  INF subscribed to NewHead event func=StartBlockListener module=service.ChainService
  ```
  
    <i>If you have any questions about installing/usage feel free to create an issue</i>

### <i>Let Chainlink grow with your web3 infrastructure</i>
