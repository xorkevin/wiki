# Cryptocurrency

## Blockchain

#### proof of work
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

## Layer 2

#### lightning network
