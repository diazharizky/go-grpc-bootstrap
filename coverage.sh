#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

# Load Environment Variables
if [ -f ./config/.env.test ]; then
    export $(cat ./config/.env.test | grep -v '#' | awk '/=/ {print $1}')
fi

# Passed argument takes precedence over environment variables
TOTAL_COVERAGE=0
if [ "$1" ]; then
    TOTAL_COVERAGE=$1
else
    TOTAL_COVERAGE=$CODE_COVERAGE
fi

if (( $(echo "$TOTAL_COVERAGE <= 0" |bc -l) )); then
    printf "${RED}Code coverage benchmark is invalid or is not set\n"
    exit 1
fi

go test ./... -coverprofile cover.out
[[ "$?" == "0" ]] && test_passed=true || test_passed=false

status="FAIL"
if ! $test_passed; then
    printf "${RED}Current test results in failure\n${NC}"
else
    printf "${GREEN}Current test results passed\n${NC}"
    test_passed=true
fi

COVERAGE=$(go tool cover -func cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}')

coverage_passed=false

if (( $(echo "$COVERAGE < $TOTAL_COVERAGE" |bc -l) )); then
    printf "${RED}Current test coverage ($COVERAGE) is less than the expected ($TOTAL_COVERAGE)\n${NC}"
else
    printf "${GREEN}Current test coverage ($COVERAGE) is greater or equal to the expected ($TOTAL_COVERAGE)\n${NC}"
    coverage_passed=true
fi

if [ "$coverage_passed" = false ] || [ "$test_passed" = false ]; then
    exit 1
fi
