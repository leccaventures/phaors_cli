# Pharos Staking Contract ABI Reference

This document provides a comprehensive ABI reference for the Pharos Network Staking smart contract. This contract is responsible for validator management, delegation, reward distribution, epoch management, and inflation control.

## Overview

- **Network**: Pharos Network
- **Compiler**: Solidity 0.8.33
- **Core Features**: Validator registration and management, staking delegation, reward claiming, epoch advancement, administrative functions, etc.
- **Inheritance**: AccessControl, Pausable, UUPSUpgradeable, ReentrancyGuard

---

## Table of Contents

1. [Data Structures (Structs)](#data-structures-structs)
2. [Constants](#constants)
3. [Validator Management](#validator-management)
4. [Delegation](#delegation)
5. [Rewards](#rewards)
6. [Commission](#commission)
7. [Pool Ownership](#pool-ownership)
8. [Epoch & Inflation](#epoch--inflation)
9. [Admin](#admin)
10. [Proxy/Upgrade](#proxyupgrade)
11. [Utility](#utility)
12. [Events](#events)
13. [Errors](#errors)

---

## Data Structures (Structs)

### Validator

A structure containing validator information.


| Field                 | Type    | Description                                                     |
| --------------------- | ------- | --------------------------------------------------------------- |
| description           | string  | Validator description                                           |
| publicKey             | string  | Public key                                                      |
| publicKeyPop          | string  | Proof of Possession (PoP) for the public key                    |
| blsPublicKey          | string  | BLS public key                                                  |
| blsPublicKeyPop       | string  | Proof of Possession (PoP) for the BLS public key                |
| endpoint              | string  | Network endpoint                                                |
| status                | uint8   | Validator status (0: Inactive, 1: Active, 2: PendingExit, etc.) |
| poolId                | bytes32 | Pool ID                                                         |
| totalStake            | uint256 | Total staked amount                                             |
| owner                 | address | Owner address                                                   |
| stakeSnapshot         | uint256 | Stake snapshot                                                  |
| pendingWithdrawStake  | uint256 | Stake amount pending withdrawal                                 |
| pendingWithdrawWindow | uint8   | Pending withdrawal window                                       |
| pendingOwner          | address | Pending owner address                                           |


### Delegator

A structure containing delegator information.


| Field                     | Type    | Description                         |
| ------------------------- | ------- | ----------------------------------- |
| principalStake            | uint256 | Principal staked amount             |
| stake                     | uint256 | Current stake (may include rewards) |
| accumulatedRewardPerShare | uint256 | Accumulated reward per share        |
| rewards                   | uint256 | Unclaimed rewards                   |
| pendingStake              | uint256 | Stake pending activation            |
| pendingWithdrawStake      | uint256 | Stake amount pending withdrawal     |
| pendingWithdrawWindow     | uint256 | Pending withdrawal window           |
| totalRewardsClaimed       | uint256 | Total rewards claimed               |
| isPendingUndelegate       | bool    | Whether an undelegation is pending  |


### PendingUndelegation

A structure containing pending undelegation information.


| Field                 | Type    | Description                                 |
| --------------------- | ------- | ------------------------------------------- |
| delegator             | address | Delegator address                           |
| poolId                | bytes32 | Pool ID                                     |
| amount                | uint256 | Undelegation amount                         |
| unlockEpoch           | uint256 | Epoch when the stake will be unlocked       |
| pendingWithdrawWindow | uint256 | Pending withdrawal window                   |
| processed             | bool    | Whether the undelegation has been processed |
| principalAmount       | uint256 | Principal amount                            |
| rewardAmount          | uint256 | Reward amount                               |


---

## Constants

### COMMISSION_SCALE

`COMMISSION_SCALE() -> uint256` 🔍 view

- **Selector**: `d142205f`
- **Description**: Scale value used for commission calculations.
- **Return Values**:

  | Return Value | Type    | Description            |
  | ------------ | ------- | ---------------------- |
  | (unnamed)    | uint256 | Commission scale value |


### DEFAULT_ADMIN_ROLE

`DEFAULT_ADMIN_ROLE() -> bytes32` 🔍 view

- **Selector**: `a217fddf`
- **Description**: Identifier for the default admin role.
- **Return Values**:

  | Return Value | Type    | Description     |
  | ------------ | ------- | --------------- |
  | (unnamed)    | bytes32 | Admin role hash |


### DEFAULT_COMMISSION_RATE

`DEFAULT_COMMISSION_RATE() -> uint256` 🔍 view

- **Selector**: `b8db983e`
- **Description**: Default commission rate.
- **Return Values**:

  | Return Value | Type    | Description             |
  | ------------ | ------- | ----------------------- |
  | (unnamed)    | uint256 | Default commission rate |


### DEFAULT_WITHDRAW_WINDOW

`DEFAULT_WITHDRAW_WINDOW() -> uint256` 🔍 view

- **Selector**: `ae631b7c`
- **Description**: Default withdrawal waiting period (window).
- **Return Values**:

  | Return Value | Type    | Description               |
  | ------------ | ------- | ------------------------- |
  | (unnamed)    | uint256 | Default withdrawal window |


### INFLATION_ADJUSTMENT_INTERVAL

`INFLATION_ADJUSTMENT_INTERVAL() -> uint256` 🔍 view

- **Selector**: `d6b5f87f`
- **Description**: Interval for inflation adjustment.
- **Return Values**:

  | Return Value | Type    | Description                      |
  | ------------ | ------- | -------------------------------- |
  | (unnamed)    | uint256 | Adjustment interval (in seconds) |


### INFLATION_ADJUSTMENT_RATE

`INFLATION_ADJUSTMENT_RATE() -> uint256` 🔍 view

- **Selector**: `c7c39bcb`
- **Description**: Rate for inflation adjustment.
- **Return Values**:

  | Return Value | Type    | Description     |
  | ------------ | ------- | --------------- |
  | (unnamed)    | uint256 | Adjustment rate |


### INITIAL_INFLATION_RATE

`INITIAL_INFLATION_RATE() -> uint256` 🔍 view

- **Selector**: `0e8601fc`
- **Description**: Initial inflation rate.
- **Return Values**:

  | Return Value | Type    | Description            |
  | ------------ | ------- | ---------------------- |
  | (unnamed)    | uint256 | Initial inflation rate |


### INITIAL_TOTAL_SUPPLY

`INITIAL_TOTAL_SUPPLY() -> uint256` 🔍 view

- **Selector**: `c04fcad8`
- **Description**: Initial total supply.
- **Return Values**:

  | Return Value | Type    | Description    |
  | ------------ | ------- | -------------- |
  | (unnamed)    | uint256 | Initial supply |


### MAX_COMMISSION_RATE

`MAX_COMMISSION_RATE() -> uint256` 🔍 view

- **Selector**: `207239c0`
- **Description**: Maximum configurable commission rate.
- **Return Values**:

  | Return Value | Type    | Description             |
  | ------------ | ------- | ----------------------- |
  | (unnamed)    | uint256 | Maximum commission rate |


### MAX_POOL_STAKE

`MAX_POOL_STAKE() -> uint256` 🔍 view

- **Selector**: `02778a83`
- **Description**: Maximum stake allowed per pool.
- **Return Values**:

  | Return Value | Type    | Description          |
  | ------------ | ------- | -------------------- |
  | (unnamed)    | uint256 | Maximum stake amount |


### MAX_QUEUE_LENGTH

`MAX_QUEUE_LENGTH() -> uint256` 🔍 view

- **Selector**: `6e73ba5b`
- **Description**: Maximum length of the waiting queue.
- **Return Values**:

  | Return Value | Type    | Description          |
  | ------------ | ------- | -------------------- |
  | (unnamed)    | uint256 | Maximum queue length |


### MIN_DELEGATOR_STAKE

`MIN_DELEGATOR_STAKE() -> uint256` 🔍 view

- **Selector**: `0f952573`
- **Description**: Minimum amount a delegator must stake.
- **Return Values**:

  | Return Value | Type    | Description               |
  | ------------ | ------- | ------------------------- |
  | (unnamed)    | uint256 | Minimum delegation amount |


### MIN_INFLATION_RATE

`MIN_INFLATION_RATE() -> uint256` 🔍 view

- **Selector**: `0fbccaae`
- **Description**: Minimum inflation rate.
- **Return Values**:

  | Return Value | Type    | Description            |
  | ------------ | ------- | ---------------------- |
  | (unnamed)    | uint256 | Minimum inflation rate |


### MIN_POOL_STAKE

`MIN_POOL_STAKE() -> uint256` 🔍 view

- **Selector**: `568735b0`
- **Description**: Minimum stake required for a pool to be active.
- **Return Values**:

  | Return Value | Type    | Description               |
  | ------------ | ------- | ------------------------- |
  | (unnamed)    | uint256 | Minimum pool stake amount |


### PRECISION

`PRECISION() -> uint256` 🔍 view

- **Selector**: `aaf5eb68`
- **Description**: Constant for calculation precision.
- **Return Values**:

  | Return Value | Type    | Description        |
  | ------------ | ------- | ------------------ |
  | (unnamed)    | uint256 | Precision constant |


### UPGRADE_INTERFACE_VERSION

`UPGRADE_INTERFACE_VERSION() -> string` 🔍 view

- **Selector**: `ad3cb1cc`
- **Description**: Version of the upgrade interface.
- **Return Values**:

  | Return Value | Type   | Description    |
  | ------------ | ------ | -------------- |
  | (unnamed)    | string | Version string |


---

## Validator Management

### registerValidator

`registerValidator(string _description, string _publicKey, string _publicKeyPop, string _blsPublicKey, string _blsPublicKeyPop, string _endpoint) -> bytes32` 💰 payable

- **Selector**: `7404f1e1`
- **Description**: Registers a new validator. A certain amount of stake must be sent with the call.
- **Parameters**:

  | Parameter        | Type   | Description                                |
  | ---------------- | ------ | ------------------------------------------ |
  | _description     | string | Validator description                      |
  | _publicKey       | string | Public key                                 |
  | _publicKeyPop    | string | Proof of Possession for the public key     |
  | _blsPublicKey    | string | BLS public key                             |
  | _blsPublicKeyPop | string | Proof of Possession for the BLS public key |
  | _endpoint        | string | Network endpoint                           |

- **Return Values**:

  | Return Value | Type    | Description       |
  | ------------ | ------- | ----------------- |
  | (unnamed)    | bytes32 | Generated pool ID |


### updateValidator

`updateValidator(bytes32 _poolId, string _description, string _endpoint)` ✏️ nonpayable

- **Selector**: `6b2f9def`
- **Description**: Updates information for an existing validator.
- **Parameters**:

  | Parameter    | Type    | Description     |
  | ------------ | ------- | --------------- |
  | _poolId      | bytes32 | Pool ID         |
  | _description | string  | New description |
  | _endpoint    | string  | New endpoint    |


### exitValidator

`exitValidator(bytes32 _poolId)` ✏️ nonpayable

- **Selector**: `3ebb52ec`
- **Description**: Requests to exit as a validator.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


### getValidator

`getValidator(bytes32 _poolId) -> Validator` 🔍 view

- **Selector**: `d5f20ff6`
- **Description**: Retrieves validator information for a specific pool ID.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type  | Description           |
  | ------------ | ----- | --------------------- |
  | validator    | tuple | Validator struct data |


### getActiveValidators

`getActiveValidators() -> bytes32[]` 🔍 view

- **Selector**: `9de70258`
- **Description**: Returns a list of pool IDs for all currently active validators.
- **Return Values**:

  | Return Value | Type      | Description              |
  | ------------ | --------- | ------------------------ |
  | (unnamed)    | bytes32[] | Array of active pool IDs |


### getAllValidators

`getAllValidators() -> bytes32[] poolIds` 🔍 view

- **Selector**: `f3513a37`
- **Description**: Returns a list of pool IDs for all registered validators.
- **Return Values**:

  | Return Value | Type      | Description           |
  | ------------ | --------- | --------------------- |
  | poolIds      | bytes32[] | Array of all pool IDs |


### getActiveValidatorCount

`getActiveValidatorCount() -> uint256` 🔍 view

- **Selector**: `37deea70`
- **Description**: Returns the number of currently active validators.
- **Return Values**:

  | Return Value | Type    | Description            |
  | ------------ | ------- | ---------------------- |
  | (unnamed)    | uint256 | Active validator count |


### getValidatorCounts

`getValidatorCounts() -> uint256 activeCount, uint256 inactiveCount, uint256 pendingExitCount` 🔍 view

- **Selector**: `8f493ae0`
- **Description**: Returns the number of validators by status.
- **Return Values**:

  | Return Value     | Type    | Description                      |
  | ---------------- | ------- | -------------------------------- |
  | activeCount      | uint256 | Active validator count           |
  | inactiveCount    | uint256 | Inactive validator count         |
  | pendingExitCount | uint256 | Count of validators pending exit |


### isValidatorActive

`isValidatorActive(bytes32 _poolId) -> bool` 🔍 view

- **Selector**: `2324e5e1`
- **Description**: Checks if a specific validator is active.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type | Description                     |
  | ------------ | ---- | ------------------------------- |
  | (unnamed)    | bool | Whether the validator is active |


### isValidatorPendingAdd

`isValidatorPendingAdd(bytes32 _poolId) -> bool` 🔍 view

- **Selector**: `66b54522`
- **Description**: Checks if a specific validator is pending addition.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type | Description                               |
  | ------------ | ---- | ----------------------------------------- |
  | (unnamed)    | bool | Whether the validator is pending addition |


### isValidatorPendingExit

`isValidatorPendingExit(bytes32 _poolId) -> bool` 🔍 view

- **Selector**: `9b4aae66`
- **Description**: Checks if a specific validator is pending exit.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type | Description                           |
  | ------------ | ---- | ------------------------------------- |
  | (unnamed)    | bool | Whether the validator is pending exit |


### isValidatorPendingUpdate

`isValidatorPendingUpdate(bytes32 _poolId) -> bool` 🔍 view

- **Selector**: `e61da203`
- **Description**: Checks if a specific validator is pending an information update.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type | Description                             |
  | ------------ | ---- | --------------------------------------- |
  | (unnamed)    | bool | Whether the validator is pending update |


### getPendingAddValidators

`getPendingAddValidators() -> bytes32[]` 🔍 view

- **Selector**: `cc8964df`
- **Description**: Returns a list of validators pending addition.
- **Return Values**:

  | Return Value | Type      | Description               |
  | ------------ | --------- | ------------------------- |
  | (unnamed)    | bytes32[] | Array of pending pool IDs |


### getPendingExitValidators

`getPendingExitValidators() -> bytes32[]` 🔍 view

- **Selector**: `6d6f0d84`
- **Description**: Returns a list of validators pending exit.
- **Return Values**:

  | Return Value | Type      | Description               |
  | ------------ | --------- | ------------------------- |
  | (unnamed)    | bytes32[] | Array of pending pool IDs |


### getPendingUpdateValidators

`getPendingUpdateValidators() -> bytes32[]` 🔍 view

- **Selector**: `1a30c776`
- **Description**: Returns a list of validators pending update.
- **Return Values**:

  | Return Value | Type      | Description               |
  | ------------ | --------- | ------------------------- |
  | (unnamed)    | bytes32[] | Array of pending pool IDs |


### getValidatorApr

`getValidatorApr(bytes32 _poolId) -> uint256 apr` 🔍 view

- **Selector**: `f03f0031`
- **Description**: Retrieves the Annual Percentage Rate (APR) for a specific validator.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type    | Description            |
  | ------------ | ------- | ---------------------- |
  | apr          | uint256 | Annual Percentage Rate |


### validators

`validators(bytes32 poolId) -> (string, string, string, string, string, string, uint8, bytes32, uint256, address, uint256, uint256, uint8, address)` 🔍 view

- **Selector**: `9bdafcb3`
- **Description**: Directly accesses the validator mapping data.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | poolId    | bytes32 | Pool ID     |

- **Return Values**: Field values of the Validator struct.

### slashValidator

`slashValidator(bytes32 _poolId)` ✏️ nonpayable

- **Selector**: `e265bca1`
- **Description**: Slashes (disciplines) a specific validator.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


---

## Delegation

### delegate

`delegate(bytes32 _poolId)` 💰 payable

- **Selector**: `c3254c23`
- **Description**: Delegates stake to a specific pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


### undelegate

`undelegate(bytes32 _poolId) -> uint256 amount, uint256 unlockEpoch` ✏️ nonpayable

- **Selector**: `9b0dc6ce`
- **Description**: Requests to undelegate the staked amount.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type    | Description                           |
  | ------------ | ------- | ------------------------------------- |
  | amount       | uint256 | Undelegated amount                    |
  | unlockEpoch  | uint256 | Epoch when the stake will be unlocked |


### addStake

`addStake(bytes32 _poolId)` 💰 payable

- **Selector**: `66da754b`
- **Description**: Adds stake to an existing delegation.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


### claimStake

`claimStake(bytes32 _poolId)` ✏️ nonpayable

- **Selector**: `4007c5ad`
- **Description**: Claims unlocked stake and receives it in the wallet.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


### withdrawStake

`withdrawStake(bytes32 _poolId, uint256 _withdrawStake)` ✏️ nonpayable

- **Selector**: `19dc5716`
- **Description**: Requests to withdraw a specific amount of stake.
- **Parameters**:

  | Parameter      | Type    | Description        |
  | -------------- | ------- | ------------------ |
  | _poolId        | bytes32 | Pool ID            |
  | _withdrawStake | uint256 | Amount to withdraw |


### getDelegator

`getDelegator(bytes32 _poolId, address _delegator) -> Delegator` 🔍 view

- **Selector**: `3a6d6bae`
- **Description**: Retrieves delegator information for a specific pool.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _poolId    | bytes32 | Pool ID           |
  | _delegator | address | Delegator address |

- **Return Values**: Delegator struct data.

### getDelegatorReward

`getDelegatorReward(bytes32 _poolId, address _delegator) -> uint256` 🔍 view

- **Selector**: `2cc0b49c`
- **Description**: Retrieves the current reward amount available for a delegator.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _poolId    | bytes32 | Pool ID           |
  | _delegator | address | Delegator address |

- **Return Values**: Reward amount.

### getDelegatorPendingWithdraw

`getDelegatorPendingWithdraw(bytes32 _poolId, address _delegator) -> uint256 pendingAmount, uint256 remainingWindow, bool isPending, uint256 unlockEpoch` 🔍 view

- **Selector**: `fc2e6398`
- **Description**: Retrieves the pending withdrawal status for a delegator.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _poolId    | bytes32 | Pool ID           |
  | _delegator | address | Delegator address |

- **Return Values**: Pending amount, remaining window, pending status, unlock epoch.

### getPendingStakeInfo

`getPendingStakeInfo(bytes32 _poolId, address _delegator) -> uint256 pendingAmount, uint256 activationEpoch, bool canActivateNow` 🔍 view

- **Selector**: `f40bd6f4`
- **Description**: Retrieves information about stake pending activation.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _poolId    | bytes32 | Pool ID           |
  | _delegator | address | Delegator address |

- **Return Values**: Pending amount, activation epoch, whether it can be activated now.

### hasPendingActivation

`hasPendingActivation(bytes32 _poolId, address _delegator) -> bool` 🔍 view

- **Selector**: `db59539c`
- **Description**: Checks if a delegator has stake pending activation.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _poolId    | bytes32 | Pool ID           |
  | _delegator | address | Delegator address |

- **Return Values**: Whether pending activation exists.

### delegators

`delegators(bytes32 poolId, address delegator) -> (uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, bool)` 🔍 view

- **Selector**: `3cd57a06`
- **Description**: Directly accesses the delegator mapping data.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | poolId    | bytes32 | Pool ID           |
  | delegator | address | Delegator address |

- **Return Values**: Field values of the Delegator struct.

### delegatorCounts

`delegatorCounts(bytes32 poolId) -> uint256 delegatorCount` 🔍 view

- **Selector**: `8482a439`
- **Description**: Returns the total number of delegators for a specific pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | poolId    | bytes32 | Pool ID     |

- **Return Values**: Delegator count.

### pendingActivationEpoch

`pendingActivationEpoch(bytes32 poolId, address delegator) -> uint256 activationEpoch` 🔍 view

- **Selector**: `63a33a97`
- **Description**: Retrieves the epoch when a delegator's stake will be activated.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | poolId    | bytes32 | Pool ID           |
  | delegator | address | Delegator address |

- **Return Values**: Activation epoch.

### pendingWithdrawStakes

`pendingWithdrawStakes(address delegator) -> uint256 pendingStake` 🔍 view

- **Selector**: `375d87ee`
- **Description**: Retrieves the total amount of stake pending withdrawal for a delegator.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | delegator | address | Delegator address |

- **Return Values**: Pending stake amount.

### getUserPendingUndelegations

`getUserPendingUndelegations(address _delegator) -> PendingUndelegation[] undelegations` 🔍 view

- **Selector**: `a5bd8279`
- **Description**: Retrieves all pending undelegations for a user.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _delegator | address | Delegator address |

- **Return Values**: Array of PendingUndelegation structs.

### getPendingUndelegationCount

`getPendingUndelegationCount() -> uint256` 🔍 view

- **Selector**: `c2fb291a`
- **Description**: Returns the total number of pending undelegations in the system.
- **Return Values**: Pending undelegation count.

### pendingUndelegations

`pendingUndelegations(uint256) -> (address, bytes32, uint256, uint256, uint256, bool, uint256, uint256)` 🔍 view

- **Selector**: `58333db7`
- **Description**: Directly accesses the pending undelegation queue.
- **Return Values**: Field values of the PendingUndelegation struct.

### undelegationReverseIndex

`undelegationReverseIndex(uint256 globalIndex) -> uint256 positionInUserArray` 🔍 view

- **Selector**: `419b87bf`
- **Description**: Converts a global index to a position within the user's array.
- **Parameters**:

  | Parameter   | Type    | Description  |
  | ----------- | ------- | ------------ |
  | globalIndex | uint256 | Global index |

- **Return Values**: Position in the user's array.

### userUndelegationHead

`userUndelegationHead(address delegator) -> uint256 headPosition` 🔍 view

- **Selector**: `bb81cbf3`
- **Description**: Returns the starting position of a user's undelegation list.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | delegator | address | Delegator address |

- **Return Values**: Starting position.

### userUndelegationIndices

`userUndelegationIndices(address delegator, uint256) -> uint256 undelegationIndices` 🔍 view

- **Selector**: `43191b6c`
- **Description**: Accesses the list of undelegation indices for a user.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | delegator | address | Delegator address |
  | (unnamed) | uint256 | Array index       |

- **Return Values**: Undelegation index.

### getTotalDelegatorCount

`getTotalDelegatorCount() -> uint256` 🔍 view

- **Selector**: `5e8f90c9`
- **Description**: Returns the total number of delegators across the entire system.
- **Return Values**: Total delegator count.

---

## Rewards

### claimReward

`claimReward(bytes32 _poolId) -> uint256` ✏️ nonpayable

- **Selector**: `f5414023`
- **Description**: Claims rewards generated from a specific pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value | Type    | Description           |
  | ------------ | ------- | --------------------- |
  | (unnamed)    | uint256 | Claimed reward amount |


### compoundRewards

`compoundRewards(bytes32 _poolId)` ✏️ nonpayable

- **Selector**: `62ed5f4e`
- **Description**: Re-stakes (compounds) generated rewards back into the pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


### settleReward

`settleReward(bytes32 _poolId)` ✏️ nonpayable

- **Selector**: `73512734`
- **Description**: Settles rewards for a specific pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


### accumulatedRewardPerShares

`accumulatedRewardPerShares(bytes32 poolId) -> uint256 rewardPerShare` 🔍 view

- **Selector**: `62e9bc1c`
- **Description**: Retrieves the accumulated reward per share (RPS) for a specific pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | poolId    | bytes32 | Pool ID     |

- **Return Values**:

  | Return Value   | Type    | Description             |
  | -------------- | ------- | ----------------------- |
  | rewardPerShare | uint256 | Reward per share amount |


### getHistoricalRPS

`getHistoricalRPS(uint256 _epoch, bytes32 _poolId) -> uint256 rps, uint256 refcount` 🔍 view

- **Selector**: `70ad2790`
- **Description**: Retrieves historical RPS data for a specific epoch.
- **Parameters**:

  | Parameter | Type    | Description  |
  | --------- | ------- | ------------ |
  | _epoch    | uint256 | Epoch number |
  | _poolId   | bytes32 | Pool ID      |

- **Return Values**: RPS value, reference count.

### getHistoricalRPSBatch

`getHistoricalRPSBatch(uint256 _epoch, bytes32[] _poolIds) -> uint256[] rpsValues, uint256[] refcounts` 🔍 view

- **Selector**: `1576f03b`
- **Description**: Batch retrieves historical RPS data for multiple pools for a specific epoch.
- **Parameters**:

  | Parameter | Type      | Description       |
  | --------- | --------- | ----------------- |
  | _epoch    | uint256   | Epoch number      |
  | _poolIds  | bytes32[] | Array of pool IDs |

- **Return Values**: Array of RPS values, array of reference counts.

### historicalRPS

`historicalRPS(uint256 epoch, bytes32 poolId) -> (uint256 rps, uint256 refcount)` 🔍 view

- **Selector**: `b1f66e13`
- **Description**: Directly accesses the historical RPS mapping data.
- **Parameters**:

  | Parameter | Type    | Description  |
  | --------- | ------- | ------------ |
  | epoch     | uint256 | Epoch number |
  | poolId    | bytes32 | Pool ID      |

- **Return Values**: RPS value, reference count.

---

## Commission

### setCommissionRate

`setCommissionRate(bytes32 _poolId, uint256 _newRate)` ✏️ nonpayable

- **Selector**: `031e7f5d`
- **Description**: Sets the commission rate for a validator pool.
- **Parameters**:

  | Parameter | Type    | Description         |
  | --------- | ------- | ------------------- |
  | _poolId   | bytes32 | Pool ID             |
  | _newRate  | uint256 | New commission rate |


### getCommissionRate

`getCommissionRate(bytes32 _poolId) -> uint256` 🔍 view

- **Selector**: `a15050cc`
- **Description**: Retrieves the current commission rate for a specific pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |

- **Return Values**: Commission rate.

### commissionRates

`commissionRates(bytes32 poolId) -> uint256 commissionRate` 🔍 view

- **Selector**: `aef30a18`
- **Description**: Directly accesses the commission rate mapping data.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | poolId    | bytes32 | Pool ID     |

- **Return Values**: Commission rate.

---

## Pool Ownership

### nominateOwner

`nominateOwner(bytes32 _poolId, address _newOwner)` ✏️ nonpayable

- **Selector**: `86edecdc`
- **Description**: Nominates a new owner for the pool.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | _poolId   | bytes32 | Pool ID           |
  | _newOwner | address | New owner address |


### acceptOwnership

`acceptOwnership(bytes32 _poolId)` ✏️ nonpayable

- **Selector**: `1b4f6c46`
- **Description**: Accepts pool ownership by the nominated new owner.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | _poolId   | bytes32 | Pool ID     |


### setDelegationEnabled

`setDelegationEnabled(bytes32 _poolId, bool _enabled)` ✏️ nonpayable

- **Selector**: `674602e9`
- **Description**: Sets whether delegation is enabled for a specific pool.
- **Parameters**:

  | Parameter | Type    | Description                  |
  | --------- | ------- | ---------------------------- |
  | _poolId   | bytes32 | Pool ID                      |
  | _enabled  | bool    | Whether to enable delegation |


### delegationEnabledMapping

`delegationEnabledMapping(bytes32 poolId) -> bool delegationEnabled` 🔍 view

- **Selector**: `c0ef4e13`
- **Description**: Accesses the mapping for delegation enabled status per pool.
- **Parameters**:

  | Parameter | Type    | Description |
  | --------- | ------- | ----------- |
  | poolId    | bytes32 | Pool ID     |

- **Return Values**: Whether delegation is enabled.

### addDelegatorToWhitelist

`addDelegatorToWhitelist(bytes32 _poolId, address _delegator)` ✏️ nonpayable

- **Selector**: `95d6e00f`
- **Description**: Adds a delegator to the whitelist of a specific pool.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _poolId    | bytes32 | Pool ID           |
  | _delegator | address | Delegator address |


### removeDelegatorFromWhitelist

`removeDelegatorFromWhitelist(bytes32 _poolId, address _delegator)` ✏️ nonpayable

- **Selector**: `5bdd5d0b`
- **Description**: Removes a delegator from the whitelist of a specific pool.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _poolId    | bytes32 | Pool ID           |
  | _delegator | address | Delegator address |


### validatorWhitelists

`validatorWhitelists(bytes32 poolId, address delegator) -> bool isWhitelisted` 🔍 view

- **Selector**: `ceb29963`
- **Description**: Accesses the mapping for whitelist status per pool.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | poolId    | bytes32 | Pool ID           |
  | delegator | address | Delegator address |

- **Return Values**: Whether the delegator is whitelisted.

---

## Epoch & Inflation

### advanceEpoch

`advanceEpoch()` ✏️ nonpayable

- **Selector**: `3cf80e6c`
- **Description**: Advances the epoch by one step.

### advanceEpoch

`advanceEpoch(bytes32[] _poolIds, uint256[] _priorityFees)` ✏️ nonpayable

- **Selector**: `d93b1a4f`
- **Description**: Advances the epoch while applying priority fees to specific pools.
- **Parameters**:

  | Parameter     | Type      | Description            |
  | ------------- | --------- | ---------------------- |
  | _poolIds      | bytes32[] | Array of pool IDs      |
  | _priorityFees | uint256[] | Array of priority fees |


### currentEpoch

`currentEpoch() -> uint256` 🔍 view

- **Selector**: `76671808`
- **Description**: Returns the current epoch number.
- **Return Values**: Epoch number.

### currentInflationRate

`currentInflationRate() -> uint256` 🔍 view

- **Selector**: `a563c3d5`
- **Description**: Returns the current inflation rate.
- **Return Values**: Inflation rate.

### lastInflationAdjustmentTime

`lastInflationAdjustmentTime() -> uint256` 🔍 view

- **Selector**: `e749a90a`
- **Description**: Returns the time when inflation was last adjusted.
- **Return Values**: Timestamp.

### lastInflationTotalSupplySnapshot

`lastInflationTotalSupplySnapshot() -> uint256` 🔍 view

- **Selector**: `d53a0f9c`
- **Description**: Returns the total supply snapshot at the time of the last inflation adjustment.
- **Return Values**: Supply snapshot.

### totalStake

`totalStake() -> uint256` 🔍 view

- **Selector**: `8b0e9f3f`
- **Description**: Returns the total staked amount across the entire system.
- **Return Values**: Total stake amount.

### totalSupply

`totalSupply() -> uint256` 🔍 view

- **Selector**: `18160ddd`
- **Description**: Returns the total supply across the entire system.
- **Return Values**: Total supply.

### getBlockchainInfo

`getBlockchainInfo() -> uint256 currentEpoch_, uint256 currentBlock_, uint256 totalStake`_ 🔍 view

- **Selector**: `600fc734`
- **Description**: Retrieves integrated information about the current blockchain and staking status.
- **Return Values**: Current epoch, current block number, total stake amount.

---

## Admin

### initialize

`initialize(address adminAddress_, address chainConfigAddress_)` ✏️ nonpayable

- **Selector**: `485cc955`
- **Description**: Initializes the contract. (Called when using the proxy pattern)
- **Parameters**:

  | Parameter           | Type    | Description                 |
  | ------------------- | ------- | --------------------------- |
  | adminAddress_       | address | Admin address               |
  | chainConfigAddress_ | address | Chain configuration address |


### pause

`pause()` ✏️ nonpayable

- **Selector**: `8456cb59`
- **Description**: Pauses the core functions of the contract.

### unpause

`unpause()` ✏️ nonpayable

- **Selector**: `3f4ba83a`
- **Description**: Resumes the paused contract functions.

### paused

`paused() -> bool` 🔍 view

- **Selector**: `5c975abb`
- **Description**: Checks if the contract is currently paused.
- **Return Values**: Whether it is paused.

### grantRole

`grantRole(bytes32 role, address account)` ✏️ nonpayable

- **Selector**: `2f2ff15d`
- **Description**: Grants a role to a specific account.
- **Parameters**:

  | Parameter | Type    | Description            |
  | --------- | ------- | ---------------------- |
  | role      | bytes32 | Role identifier        |
  | account   | address | Target account address |


### revokeRole

`revokeRole(bytes32 role, address account)` ✏️ nonpayable

- **Selector**: `d547741f`
- **Description**: Revokes a role from a specific account.
- **Parameters**:

  | Parameter | Type    | Description            |
  | --------- | ------- | ---------------------- |
  | role      | bytes32 | Role identifier        |
  | account   | address | Target account address |


### renounceRole

`renounceRole(bytes32 role, address callerConfirmation)` ✏️ nonpayable

- **Selector**: `36568abe`
- **Description**: Renounces a role held by the caller.
- **Parameters**:

  | Parameter          | Type    | Description                 |
  | ------------------ | ------- | --------------------------- |
  | role               | bytes32 | Role identifier             |
  | callerConfirmation | address | Caller confirmation address |


### hasRole

`hasRole(bytes32 role, address account) -> bool` 🔍 view

- **Selector**: `91d14854`
- **Description**: Checks if a specific account holds a given role.
- **Parameters**:

  | Parameter | Type    | Description            |
  | --------- | ------- | ---------------------- |
  | role      | bytes32 | Role identifier        |
  | account   | address | Target account address |

- **Return Values**: Whether the account holds the role.

### getRoleAdmin

`getRoleAdmin(bytes32 role) -> bytes32` 🔍 view

- **Selector**: `248a9ca3`
- **Description**: Retrieves the admin role for a specific role.
- **Parameters**:

  | Parameter | Type    | Description     |
  | --------- | ------- | --------------- |
  | role      | bytes32 | Role identifier |

- **Return Values**: Admin role identifier.

### setRoleAdmin

`setRoleAdmin(bytes32 role, bytes32 adminRole)` ✏️ nonpayable

- **Selector**: `1e4e0091`
- **Description**: Sets the admin role for a specific role.
- **Parameters**:

  | Parameter | Type    | Description       |
  | --------- | ------- | ----------------- |
  | role      | bytes32 | Target role       |
  | adminRole | bytes32 | Admin role to set |


### updateChainConfigAddress

`updateChainConfigAddress(address chainConfigAddress_)` ✏️ nonpayable

- **Selector**: `7ba04992`
- **Description**: Updates the chain configuration contract address.
- **Parameters**:

  | Parameter           | Type    | Description                     |
  | ------------------- | ------- | ------------------------------- |
  | chainConfigAddress_ | address | New chain configuration address |


### cfg

`cfg() -> address` 🔍 view

- **Selector**: `0457dad2`
- **Description**: Returns the currently set chain configuration contract address.
- **Return Values**: Contract address.

### migrateV1ToV2

`migrateV1ToV2()` ✏️ nonpayable

- **Selector**: `21b3917e`
- **Description**: Migrates data from V1 to V2.

### v1ToV2Migrated

`v1ToV2Migrated() -> bool` 🔍 view

- **Selector**: `fab7de7f`
- **Description**: Checks if the migration from V1 to V2 has been completed.
- **Return Values**: Whether migration is complete.

---

## Proxy/Upgrade

### proxiableUUID

`proxiableUUID() -> bytes32` 🔍 view

- **Selector**: `52d1902d`
- **Description**: Returns the unique identifier for UUPS upgrades.
- **Return Values**: UUID.

### upgradeToAndCall

`upgradeToAndCall(address newImplementation, bytes data)` 💰 payable

- **Selector**: `4f1ef286`
- **Description**: Upgrades the contract to a new implementation and executes additional data.
- **Parameters**:

  | Parameter         | Type    | Description                |
  | ----------------- | ------- | -------------------------- |
  | newImplementation | address | New implementation address |
  | data              | bytes   | Data to execute            |


### getImplAddress

`getImplAddress() -> address` 🔍 view

- **Selector**: `7311ad81`
- **Description**: Returns the address of the current implementation.
- **Return Values**: Implementation address.

### supportsInterface

`supportsInterface(bytes4 interfaceId) -> bool` 🔍 view

- **Selector**: `01ffc9a7`
- **Description**: Checks if a specific interface is supported.
- **Parameters**:

  | Parameter   | Type   | Description  |
  | ----------- | ------ | ------------ |
  | interfaceId | bytes4 | Interface ID |

- **Return Values**: Whether it is supported.

---

## Utility

### unionArrays

`unionArrays(bytes32[] _poolId1, bytes32[] _poolId2) -> bytes32[]` 🔍 pure

- **Selector**: `d3c26581`
- **Description**: Combines two pool ID arrays and returns the result.
- **Parameters**:

  | Parameter | Type      | Description  |
  | --------- | --------- | ------------ |
  | _poolId1  | bytes32[] | First array  |
  | _poolId2  | bytes32[] | Second array |

- **Return Values**: Combined array.

### getUserTotalInfo

`getUserTotalInfo(address _delegator) -> uint256 totalStake_, uint256 totalRewards_, uint256 totalClaimed`_ 🔍 view

- **Selector**: `25667b3b`
- **Description**: Retrieves information about a user's total stake, total rewards, and total claimed amount.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _delegator | address | Delegator address |

- **Return Values**: Total stake, total rewards, total claimed amount.

### getUserValidators

`getUserValidators(address _delegator) -> bytes32[] poolIds, uint256[] stakes, uint256[] rewards` 🔍 view

- **Selector**: `4e3266e3`
- **Description**: Retrieves all validator pools a user has delegated to, along with the corresponding stake and reward information.
- **Parameters**:

  | Parameter  | Type    | Description       |
  | ---------- | ------- | ----------------- |
  | _delegator | address | Delegator address |

- **Return Values**: Array of pool IDs, array of stakes, array of rewards.

### activePoolIds

`activePoolIds(uint256) -> bytes32` 🔍 view

- **Selector**: `1f9c324e`
- **Description**: Accesses the active pool ID list by index.
- **Return Values**: Pool ID.

### pendingAddPoolIds

`pendingAddPoolIds(uint256) -> bytes32` 🔍 view

- **Selector**: `104b1a29`
- **Description**: Accesses the pending addition pool ID list by index.
- **Return Values**: Pool ID.

### pendingExitPoolIds

`pendingExitPoolIds(uint256) -> bytes32` 🔍 view

- **Selector**: `893dfe68`
- **Description**: Accesses the pending exit pool ID list by index.
- **Return Values**: Pool ID.

### pendingUpdatePoolIds

`pendingUpdatePoolIds(uint256) -> bytes32` 🔍 view

- **Selector**: `483bb601`
- **Description**: Accesses the pending update pool ID list by index.
- **Return Values**: Pool ID.

---

## Events

### ChainConfigUpdate

`event ChainConfigUpdate(address indexed oldChainConfigAddress, address indexed newChainConfigAddress)`

- **Description**: Emitted when the chain configuration address is updated.
- **Parameters**:

  | Parameter             | Type    | Indexed | Description                    |
  | --------------------- | ------- | ------- | ------------------------------ |
  | oldChainConfigAddress | address | Yes     | Previous configuration address |
  | newChainConfigAddress | address | Yes     | New configuration address      |


### CoinMinted

`event CoinMinted(address indexed beneficiary, uint8 indexed reason, uint256 amount)`

- **Description**: Emitted when new coins are minted.
- **Parameters**:

  | Parameter   | Type    | Indexed | Description         |
  | ----------- | ------- | ------- | ------------------- |
  | beneficiary | address | Yes     | Beneficiary address |
  | reason      | uint8   | Yes     | Reason for minting  |
  | amount      | uint256 | No      | Minted amount       |


### CommissionRateChanged

`event CommissionRateChanged(bytes32 indexed poolId, uint256 oldRate, uint256 newRate)`

- **Description**: Emitted when the commission rate is changed.
- **Parameters**:

  | Parameter | Type    | Indexed | Description   |
  | --------- | ------- | ------- | ------------- |
  | poolId    | bytes32 | Yes     | Pool ID       |
  | oldRate   | uint256 | No      | Previous rate |
  | newRate   | uint256 | No      | New rate      |


### DelegatorWithdrawClaimed

`event DelegatorWithdrawClaimed(address indexed delegator, bytes32 indexed poolId, uint256 amount, uint256 indexed epochNumber)`

- **Description**: Emitted when a delegator receives stake that was pending withdrawal.
- **Parameters**:

  | Parameter   | Type    | Indexed | Description       |
  | ----------- | ------- | ------- | ----------------- |
  | delegator   | address | Yes     | Delegator address |
  | poolId      | bytes32 | Yes     | Pool ID           |
  | amount      | uint256 | No      | Received amount   |
  | epochNumber | uint256 | Yes     | Epoch number      |


### DomainUpdate

`event DomainUpdate(bytes32 indexed poolId, string description, string publicKey, string blsPublicKey, string endpoint, uint64 effectiveBlockNum, uint8 status)`

- **Description**: Emitted when validator domain information is updated.
- **Parameters**:

  | Parameter         | Type    | Indexed | Description            |
  | ----------------- | ------- | ------- | ---------------------- |
  | poolId            | bytes32 | Yes     | Pool ID                |
  | description       | string  | No      | Description            |
  | publicKey         | string  | No      | Public key             |
  | blsPublicKey      | string  | No      | BLS public key         |
  | endpoint          | string  | No      | Endpoint               |
  | effectiveBlockNum | uint64  | No      | Effective block number |
  | status            | uint8   | No      | Status                 |


### EpochChange

`event EpochChange(uint256 indexed epochNumber, uint256 indexed blockNumber, uint256 timestamp, uint256 totalStake, bytes32[] activeValidators)`

- **Description**: Emitted when the epoch changes.
- **Parameters**:

  | Parameter        | Type      | Indexed | Description               |
  | ---------------- | --------- | ------- | ------------------------- |
  | epochNumber      | uint256   | Yes     | Epoch number              |
  | blockNumber      | uint256   | Yes     | Block number              |
  | timestamp        | uint256   | No      | Timestamp                 |
  | totalStake       | uint256   | No      | Total stake               |
  | activeValidators | bytes32[] | No      | List of active validators |


### ErrorOccurred

`event ErrorOccurred(uint256 indexed epochNumber, uint256 indexed blockNumber, uint256 errorCode, address recipient, uint256 amount, uint256 balance)`

- **Description**: Emitted when an error occurs during system processing.
- **Parameters**:

  | Parameter   | Type    | Indexed | Description    |
  | ----------- | ------- | ------- | -------------- |
  | epochNumber | uint256 | Yes     | Epoch number   |
  | blockNumber | uint256 | Yes     | Block number   |
  | errorCode   | uint256 | No      | Error code     |
  | recipient   | address | No      | Target address |
  | amount      | uint256 | No      | Amount         |
  | balance     | uint256 | No      | Balance        |


### InflationRateAdjusted

`event InflationRateAdjusted(uint256 oldRate, uint256 newRate, uint256 timestamp)`

- **Description**: Emitted when the inflation rate is adjusted.
- **Parameters**:

  | Parameter | Type    | Indexed | Description   |
  | --------- | ------- | ------- | ------------- |
  | oldRate   | uint256 | No      | Previous rate |
  | newRate   | uint256 | No      | New rate      |
  | timestamp | uint256 | No      | Timestamp     |


### Initialized

`event Initialized(uint64 version)`

- **Description**: Emitted when the contract is initialized.
- **Parameters**:

  | Parameter | Type   | Indexed | Description            |
  | --------- | ------ | ------- | ---------------------- |
  | version   | uint64 | No      | Initialization version |


### OwnerNominated

`event OwnerNominated(bytes32 indexed poolId, address indexed currentOwner, address indexed pendingOwner)`

- **Description**: Emitted when a new owner is nominated for a pool.
- **Parameters**:

  | Parameter    | Type    | Indexed | Description     |
  | ------------ | ------- | ------- | --------------- |
  | poolId       | bytes32 | Yes     | Pool ID         |
  | currentOwner | address | Yes     | Current owner   |
  | pendingOwner | address | Yes     | Nominated owner |


### OwnershipTransferred

`event OwnershipTransferred(bytes32 indexed poolId, address indexed newOwner, address indexed previousOwner)`

- **Description**: Emitted when pool ownership is transferred.
- **Parameters**:

  | Parameter     | Type    | Indexed | Description    |
  | ------------- | ------- | ------- | -------------- |
  | poolId        | bytes32 | Yes     | Pool ID        |
  | newOwner      | address | Yes     | New owner      |
  | previousOwner | address | Yes     | Previous owner |


### Paused

`event Paused(address account)`

- **Description**: Emitted when the contract is paused.
- **Parameters**:

  | Parameter | Type    | Indexed | Description                      |
  | --------- | ------- | ------- | -------------------------------- |
  | account   | address | No      | Account that paused the contract |


### RewardClaimed

`event RewardClaimed(address indexed delegator, bytes32 indexed poolId, uint256 amount)`

- **Description**: Emitted when rewards are claimed.
- **Parameters**:

  | Parameter | Type    | Indexed | Description       |
  | --------- | ------- | ------- | ----------------- |
  | delegator | address | Yes     | Delegator address |
  | poolId    | bytes32 | Yes     | Pool ID           |
  | amount    | uint256 | No      | Claimed amount    |


### RewardsCompounded

`event RewardsCompounded(address indexed delegator, bytes32 indexed poolId, uint256 amount)`

- **Description**: Emitted when rewards are compounded (re-staked).
- **Parameters**:

  | Parameter | Type    | Indexed | Description       |
  | --------- | ------- | ------- | ----------------- |
  | delegator | address | Yes     | Delegator address |
  | poolId    | bytes32 | Yes     | Pool ID           |
  | amount    | uint256 | No      | Compounded amount |


### RoleAdminChanged

`event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)`

- **Description**: Emitted when the admin role for a role is changed.
- **Parameters**:

  | Parameter         | Type    | Indexed | Description         |
  | ----------------- | ------- | ------- | ------------------- |
  | role              | bytes32 | Yes     | Target role         |
  | previousAdminRole | bytes32 | Yes     | Previous admin role |
  | newAdminRole      | bytes32 | Yes     | New admin role      |


### RoleGranted

`event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)`

- **Description**: Emitted when a role is granted to an account.
- **Parameters**:

  | Parameter | Type    | Indexed | Description                   |
  | --------- | ------- | ------- | ----------------------------- |
  | role      | bytes32 | Yes     | Role identifier               |
  | account   | address | Yes     | Account granted the role      |
  | sender    | address | Yes     | Account that granted the role |


### RoleRevoked

`event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)`

- **Description**: Emitted when a role is revoked from an account.
- **Parameters**:

  | Parameter | Type    | Indexed | Description                   |
  | --------- | ------- | ------- | ----------------------------- |
  | role      | bytes32 | Yes     | Role identifier               |
  | account   | address | Yes     | Account role was revoked from |
  | sender    | address | Yes     | Account that revoked the role |


### StakeAdded

`event StakeAdded(address indexed delegator, bytes32 indexed poolId, uint256 amount)`

- **Description**: Emitted when stake is added to an existing delegation.
- **Parameters**:

  | Parameter | Type    | Indexed | Description       |
  | --------- | ------- | ------- | ----------------- |
  | delegator | address | Yes     | Delegator address |
  | poolId    | bytes32 | Yes     | Pool ID           |
  | amount    | uint256 | No      | Added amount      |


### StakeCancelled

`event StakeCancelled(address indexed delegator, bytes32 indexed poolId, uint256 amount, uint256 indexed epochNumber)`

- **Description**: Emitted when a pending stake is cancelled.
- **Parameters**:

  | Parameter   | Type    | Indexed | Description       |
  | ----------- | ------- | ------- | ----------------- |
  | delegator   | address | Yes     | Delegator address |
  | poolId      | bytes32 | Yes     | Pool ID           |
  | amount      | uint256 | No      | Cancelled amount  |
  | epochNumber | uint256 | Yes     | Epoch number      |


### StakeClaimed

`event StakeClaimed(address indexed user, uint256 amount)`

- **Description**: Emitted when stake is claimed.
- **Parameters**:

  | Parameter | Type    | Indexed | Description    |
  | --------- | ------- | ------- | -------------- |
  | user      | address | Yes     | User address   |
  | amount    | uint256 | No      | Claimed amount |


### StakeDelegated

`event StakeDelegated(address indexed delegator, bytes32 indexed poolId, uint256 amount, uint256 effectiveEpoch)`

- **Description**: Emitted when stake is delegated.
- **Parameters**:

  | Parameter      | Type    | Indexed | Description       |
  | -------------- | ------- | ------- | ----------------- |
  | delegator      | address | Yes     | Delegator address |
  | poolId         | bytes32 | Yes     | Pool ID           |
  | amount         | uint256 | No      | Delegated amount  |
  | effectiveEpoch | uint256 | No      | Effective epoch   |


### StakeUndelegated

`event StakeUndelegated(address indexed delegator, bytes32 indexed poolId, uint256 amount, uint256 unlockEpoch)`

- **Description**: Emitted when stake is undelegated.
- **Parameters**:

  | Parameter   | Type    | Indexed | Description        |
  | ----------- | ------- | ------- | ------------------ |
  | delegator   | address | Yes     | Delegator address  |
  | poolId      | bytes32 | Yes     | Pool ID            |
  | amount      | uint256 | No      | Undelegated amount |
  | unlockEpoch | uint256 | No      | Unlock epoch       |


### TotalSupplyIncreased

`event TotalSupplyIncreased(uint256 indexed epochNumber, uint256 amount, uint256 newTotalSupply)`

- **Description**: Emitted when the total supply increases.
- **Parameters**:

  | Parameter      | Type    | Indexed | Description      |
  | -------------- | ------- | ------- | ---------------- |
  | epochNumber    | uint256 | Yes     | Epoch number     |
  | amount         | uint256 | No      | Increase amount  |
  | newTotalSupply | uint256 | No      | New total supply |


### Unpaused

`event Unpaused(address account)`

- **Description**: Emitted when the contract is unpaused.
- **Parameters**:

  | Parameter | Type    | Indexed | Description                        |
  | --------- | ------- | ------- | ---------------------------------- |
  | account   | address | No      | Account that unpaused the contract |


### Upgraded

`event Upgraded(address indexed implementation)`

- **Description**: Emitted when the contract is upgraded.
- **Parameters**:

  | Parameter      | Type    | Indexed | Description                |
  | -------------- | ------- | ------- | -------------------------- |
  | implementation | address | Yes     | New implementation address |


### ValidatorExitRequested

`event ValidatorExitRequested(bytes32 indexed poolId)`

- **Description**: Emitted when a validator requests to exit.
- **Parameters**:

  | Parameter | Type    | Indexed | Description |
  | --------- | ------- | ------- | ----------- |
  | poolId    | bytes32 | Yes     | Pool ID     |


### ValidatorRegistered

`event ValidatorRegistered(address indexed validator, bytes32 indexed poolId, uint256 amount)`

- **Description**: Emitted when a validator is registered.
- **Parameters**:

  | Parameter | Type    | Indexed | Description       |
  | --------- | ------- | ------- | ----------------- |
  | validator | address | Yes     | Validator address |
  | poolId    | bytes32 | Yes     | Pool ID           |
  | amount    | uint256 | No      | Staked amount     |


### ValidatorReward

`event ValidatorReward(bytes32 indexed poolId, address indexed owner, uint256 indexed epochNumber, uint256 blockNumber, uint256 baseReward, uint256 feeReward, uint256 totalReward)`

- **Description**: Emitted when a validator receives a reward.
- **Parameters**:

  | Parameter   | Type    | Indexed | Description   |
  | ----------- | ------- | ------- | ------------- |
  | poolId      | bytes32 | Yes     | Pool ID       |
  | owner       | address | Yes     | Owner address |
  | epochNumber | uint256 | Yes     | Epoch number  |
  | blockNumber | uint256 | No      | Block number  |
  | baseReward  | uint256 | No      | Base reward   |
  | feeReward   | uint256 | No      | Fee reward    |
  | totalReward | uint256 | No      | Total reward  |


### ValidatorUpdated

`event ValidatorUpdated(bytes32 indexed poolId)`

- **Description**: Emitted when validator information is updated.
- **Parameters**:

  | Parameter | Type    | Indexed | Description |
  | --------- | ------- | ------- | ----------- |
  | poolId    | bytes32 | Yes     | Pool ID     |


---

## Errors

### AccessControlBadConfirmation

`error AccessControlBadConfirmation()`

- **Description**: Thrown when the confirmation address for renouncing a role is incorrect.

### AccessControlUnauthorizedAccount

`error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)`

- **Description**: Thrown when an account attempts an action without the required role.
- **Parameters**:

  | Parameter  | Type    | Description     |
  | ---------- | ------- | --------------- |
  | account    | address | Account address |
  | neededRole | bytes32 | Required role   |


### AddressEmptyCode

`error AddressEmptyCode(address target)`

- **Description**: Thrown when an address that should contain code is empty.
- **Parameters**:

  | Parameter | Type    | Description    |
  | --------- | ------- | -------------- |
  | target    | address | Target address |


### ERC1967InvalidImplementation

`error ERC1967InvalidImplementation(address implementation)`

- **Description**: Thrown when the provided implementation address is invalid.
- **Parameters**:

  | Parameter      | Type    | Description            |
  | -------------- | ------- | ---------------------- |
  | implementation | address | Implementation address |


### ERC1967NonPayable

`error ERC1967NonPayable()`

- **Description**: Thrown when a non-payable function receives Ether.

### EnforcedPause

`error EnforcedPause()`

- **Description**: Thrown when an action is attempted while the contract is paused.

### EpochNotCompleted

`error EpochNotCompleted(uint256 lastEpochStart, uint256 epochDuration, uint256 currentTime)`

- **Description**: Thrown when an action requiring epoch completion is attempted before the epoch ends.
- **Parameters**:

  | Parameter      | Type    | Description                  |
  | -------------- | ------- | ---------------------------- |
  | lastEpochStart | uint256 | Start time of the last epoch |
  | epochDuration  | uint256 | Duration of an epoch         |
  | currentTime    | uint256 | Current time                 |


### ExpectedPause

`error ExpectedPause()`

- **Description**: Thrown when an action is attempted while the contract is not paused, but should be.

### FailedCall

`error FailedCall()`

- **Description**: Thrown when a low-level call fails.

### InvalidInitialization

`error InvalidInitialization()`

- **Description**: Thrown when contract initialization fails or is attempted multiple times.

### NotInitializing

`error NotInitializing()`

- **Description**: Thrown when an initialization function is called outside of the initialization process.

### ReentrancyGuardReentrantCall

`error ReentrancyGuardReentrantCall()`

- **Description**: Thrown when a reentrant call is detected.

### SafeCastOverflowedUintDowncast

`error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)`

- **Description**: Thrown when a downcast operation results in an overflow.
- **Parameters**:

  | Parameter | Type    | Description      |
  | --------- | ------- | ---------------- |
  | bits      | uint8   | Number of bits   |
  | value     | uint256 | Value being cast |


### StringsInvalidChar

`error StringsInvalidChar()`

- **Description**: Thrown when an invalid character is encountered in a string operation.

### UUPSUnauthorizedCallContext

`error UUPSUnauthorizedCallContext()`

- **Description**: Thrown when a UUPS upgrade is attempted from an unauthorized context.

### UUPSUnsupportedProxiableUUID

`error UUPSUnsupportedProxiableUUID(bytes32 slot)`

- **Description**: Thrown when the proxiable UUID is not supported.
- **Parameters**:

  | Parameter | Type    | Description  |
  | --------- | ------- | ------------ |
  | slot      | bytes32 | Storage slot |

