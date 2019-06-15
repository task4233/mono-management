#!/bin/bash

BASE_URI="localhost:8080/api/v1/"

PARAMS=""

TEST_ARRAY=(
    "GET" "mono/" '"get monos"'
    "POST" "mono/new/" 
)

for (( i=0; i<${#TEST_ARRAY[@]}; i+=3))
do
    case "${TEST_ARRAY[$i]}" in
        "GET" )
            REQ_URI="${BASE_URI}${TEST_ARRAY[("$i+1" | bc)]}"
            EXPECT="${TEST_ARRAY[("$i+2" | bc)]}"
            RESPONSE=$(curl --silent $REQ_URI)
            
            if [ "$RESPONSE" = "$EXPECT" ]; then
                echo "${REQ_URI} is OK"
            else
                echo "${REQ_URI} is NG"
                echo "Expected: ${EXPECT}, Response: ${RESPONSE}"
                exit
            fi
            break;;
        * )
            break;;
    esac
done

