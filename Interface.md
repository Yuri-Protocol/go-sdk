## API instruction

### 1、Mint721Nft 
####  Mint ERC-721 NFT
`
Mint(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error)
`

### 2、Account 
#### computes the token bound account address
`
Account(opts *bind.CallOpts, implementation common.Address, salt [32]byte, chainId *big.Int, 
         tokenContract common.Address, tokenId *big.Int) (common.Address, error)
`

### 3、CreateAccount

`
CreateAccount(opts *bind.TransactOpts, implementation common.Address, salt [32]byte, 
   chainId *big.Int, tokenContract common.Address, tokenId *big.Int) (*types.Transaction, error) 
`

### 4、Token Transfer
#### 4.1、native token transfer
`
Execute(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error)
`

#### 4.2 ERC20 token transer
`TransferToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) 
`
