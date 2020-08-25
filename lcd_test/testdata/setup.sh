#!/usr/bin/env bash

PASSWORD="1234567890"
ADDR="fetch16xyempempp92x9hyzz9wrgf94r6j9h5fu8gzym"
RECEIVER="fetch17gx5vwpm0y2k59tw0x00ccug234n56cgvk0wva"
VALIDATOR="fetchvaloper16xyempempp92x9hyzz9wrgf94r6j9h5ferhphu"
AMOUNT="1000000stake"
CHAIN="lcd"
PROPOSALID="2"
HOME="/tmp/contract_tests/.wasmcli"
SWAGGER='/tmp/contract_tests/swagger.yaml'

# sleeping a whole second between each step is a conservative precaution
# check lcd_test/testdata/state.tar.gz -> .wasmd/config/config.toml precommit_timeout = 500ms
sleep 1s
echo ${PASSWORD} | ./build/wasmcli tx gov submit-proposal --home ${HOME} --from ${ADDR} --chain-id ${CHAIN} --type text --title test --description test_description --deposit 10000stake --yes
sleep 1s
echo ${PASSWORD} | ./build/wasmcli tx gov deposit --home ${HOME} --from ${ADDR} --chain-id ${CHAIN} ${PROPOSALID} 1000000000stake --yes
sleep 1s
echo ${PASSWORD} | ./build/wasmcli tx gov vote --home ${HOME} --from ${ADDR} --yes --chain-id ${CHAIN} ${PROPOSALID} Yes
sleep 1s
HASH=$(echo ${PASSWORD} | ./build/wasmcli tx send --home ${HOME} ${ADDR} ${RECEIVER} ${AMOUNT} --yes --chain-id ${CHAIN} | awk '/txhash.*/{print $2}')
sed -i.bak -e "s/BCBE20E8D46758B96AE5883B792858296AC06E51435490FBDCAE25A72B3CC76B/${HASH}/g" "${SWAGGER}"
echo "Replaced dummy with actual transaction hash ${HASH}"
sleep 1s
echo ${PASSWORD} | ./build/wasmcli tx staking unbond --home ${HOME} --from ${ADDR} ${VALIDATOR} 100stake --yes --chain-id ${CHAIN}

