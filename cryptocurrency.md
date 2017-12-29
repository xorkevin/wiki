# Cryptocurrency

## Blockchain

### proof of work

- hash of a block includes:
  - transaction data
  - previous block hash
  - nonce
- object is to find a nonce s.t. the hash of the block produces a hash with a certain number of zeros at its start
  - the number of zeros required can be adjusted as a parameter to vary the difficulty and ultimately the "mining" rate
- once a nonce is found, the rest of the network is notified via some sort of gossip protocol
- the network will then begin work on mining the next block
- the longest chain will be chosen by the network since it has the most "work" invested in it
  - this drives the network to consensus

### implications

- the puzzle in a proof of work based blockchain must be difficult
  - a high hash rate is necessary to make it difficult to cheat the system
  - the cost must be economically difficult to pay, whether in the form of electricity or computers
- 51% attack is expensive as it requires gaining a large proportion of the computing resources on the network
- attempting to rewrite history would require remining all the blocks that have been appended to the chain since then
  - must be done so at a rate significantly faster than newer blocks are being added to the chain
  - the further down the chain the larger the cost, hence the more secure the block
  - total cost is extremely high only if the puzzle is expensive to complete

## Layer 2

### lightning network

- smart contract payment channel between two accounts on the blockchain
  - payment channels allow both participants to agree on the new distribution of funds within the channel via pub/priv crypto
- a network of these payment channels could be overlaid on top of the existing blockchain
- payment channels can be settled back onto the blockchain

#### example
- suppose two accounts Alice and Bob each commit 500 to a payment channel, AB
- suppose two accounts Bob and Carol each commit 500 to a payment channel, BC
- suppose Alice wants to give Carol 50
- Alice can sign an atomic transaction that changes AB to 450 | 550, and BC to 450 | 550
- Once Bob and Carol sign, does the transaction become committed
- Bob still has 1000 total, Alice is short 50, and Carol has gained 50

### implications

- supports a much higher rate of transactions of several magnitudes since they are not reliant on mining
- settlement rate of payment channels must have a certain minimum to prevent the mining rate from dropping too low
  - a 51% attack could potentially occur otherwise
- centralization is another concern where a couple hub nodes could broker many channels and create single points of failure
