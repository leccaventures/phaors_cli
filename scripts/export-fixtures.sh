#!/usr/bin/env bash
set -euo pipefail

BINARY_PATH="${1:-./bin/pharoscli}"
OUTPUT_DIR="./test"

POOL_ID_PRIMARY="0x8467c9cf1536642e27ed004d13af86753c312d2d7176a0ca48fe806894b573e2"
POOL_ID_SECONDARY="0x0a2c55a00df40b658738e1417622b69531d11b37c2cfac825a8e9f565a8064eb"
ADDRESS="0x00000B834695138Ffd7E4BF07CB4470c292F4eE4"

mkdir -p "${OUTPUT_DIR}"

"${BINARY_PATH}" getAllValidators > "${OUTPUT_DIR}/getAllValidators.json"
"${BINARY_PATH}" getActiveValidators > "${OUTPUT_DIR}/getActiveValidators.json"
"${BINARY_PATH}" getValidator "${POOL_ID_PRIMARY}" > "${OUTPUT_DIR}/getValidator.json"
"${BINARY_PATH}" validators "${POOL_ID_PRIMARY}" > "${OUTPUT_DIR}/validators.json"
"${BINARY_PATH}" isValidatorActive "${POOL_ID_PRIMARY}" > "${OUTPUT_DIR}/isValidatorActive.json"
"${BINARY_PATH}" getActiveValidatorCount > "${OUTPUT_DIR}/getActiveValidatorCount.json"
"${BINARY_PATH}" getValidatorCounts > "${OUTPUT_DIR}/getValidatorCounts.json"
"${BINARY_PATH}" getDelegator "${POOL_ID_PRIMARY}" "${ADDRESS}" > "${OUTPUT_DIR}/getDelegator.json"
"${BINARY_PATH}" getPendingStakeInfo "${POOL_ID_SECONDARY}" "${ADDRESS}" > "${OUTPUT_DIR}/getPendingStakeInfo.json"
