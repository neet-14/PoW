### Proof of Work Consensus Algorithm - 

- Proof of work consensus is used to achieve a distributed consensus among the working nodes of a blockchain. 
- I have used a `difficulty` parameter to determine the complexity and computational power of the mining.

- As the value of `difficulty` increases the value of the `target` hash value decreases hence making it diffic

#### Difficulty:

Difficulty is a measure of how hard it is to find a valid hash for a block. In blockchain systems like Bitcoin, the difficulty is adjusted to ensure that blocks are mined at a consistent rate.
#### Target:

The target is a threshold value that the hash of a block must be less than or equal to in order for the block to be considered valid.
The hash of the block is a large number, and it must be less than this target value.

#### How Difficulty Affects Target

    When the difficulty increases:

    The target value becomes smaller.
    This means there are fewer valid hash values that meet the criteria.
    As a result, it becomes harder to find a valid hash.

    When the difficulty decreases:

    The target value becomes larger.
    This means there are more valid hash values that meet the criteria.
    As a result, it becomes easier to find a valid hash.