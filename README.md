# cosmosChain
=======
# cosmos Blockchain

This guide provides step-by-step instructions for setting up a Cosmos-SDK-based blockchain using `cosmosd`. 

## Prerequisites  

Ensure you have `cosmosd` installed and configured on your system before proceeding by running `make install`.  

## Steps  

### 1. Initialize the Blockchain  

Initialize your blockchain with the following command, specifying your node's moniker and chain ID. Set the default denomination to `cosmos`.  

```bash  
sudo ./cosmosd init cosmos --chain-id cosmos --default-denom cosmos
```


### 2. Create Key Pairs

Create a key pair for your validator node. This key will be used to sign blocks and transactions.  

```bash  
sudo ./cosmosd keys add  cosmos-node 
```

Example output:

```bash  
- address: cosmos18muf4rc9jgzghmcjz2mzkkpe0l70t54uedwlu7  
  name: cosmos-node  
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Atkt4/rSXc4FCn3jXq98uWg77OfJcYG8MbtfhwNi0laZ"}'  
  type: local
```

### 3. Add Genesis Accounts

Add accounts to the genesis file with a specified starting balance. You can add multiple accounts and specify tokens to multiple accounts if needed.

```bash 
sudo ./cosmosd genesis add-genesis-account cosmos18muf4rc9jgzghmcjz2mzkkpe0l70t54uedwlu7 1000000cosmos  
sudo ./cosmosd genesis add-genesis-account cosmos18muf4rc9jgzghmcjz2mzkkpe0l70t54uedwlu7 2000000cosmos --append
```


### 4. Create and Collect Genesis Transactions

Create the genesis transaction file and collect signatures from the validator and other accounts.

```bash
sudo ./cosmosd genesis gentx cosmos-node 1000000cosmos --chain-id cosmos
sudo ./cosmosd genesis collect-gentxs  
```

### 5. Validate the Genesis File

Ensure your genesis file is correctly structured without errors.

```bash
sudo ./cosmosd genesis validate-genesis  
```


### 6. Start the Blockchain Node
Begin running your blockchain node with a specified minimum gas price for transactions.

```bash
sudo ./cosmosd start --minimum-gas-prices "0.01cosmos"
```
