#!/bin/bash

newDirectory="./tests/coverage"
if [ ! -d "$newDirectory" ]; then
    mkdir -p "$newDirectory"
fi

go test ./tests/unit/domain/tasklist_test.go -coverpkg="github.com/tweedledo/core/domain" -coverprofile="./tests/coverage/coverage-domain.out"
go test ./tests/unit/service/tasklist_service_test.go -coverpkg="github.com/tweedledo/core/service" -coverprofile="./tests/coverage/coverage-service.out"

sourceFile="./tests/coverage/coverage-domain.out"
appendFile="./tests/coverage/coverage-service.out"
thirdFile="./tests/coverage/coverage.out"

# Copy lines from line 2 to the end of the source file and save to a temporary file
tail -n +2 "$sourceFile" > temp.txt
tail -n +2 "$appendFile" > temp2.txt

# Append the content of the append file to the temporary file
cat temp2.txt >> temp.txt

# Rename the temporary file to the third file
mv temp.txt "$thirdFile"

echo "Complete coverage in $thirdFile file"

go tool cover -html="$thirdFile"

rm temp.txt
rm temp2.txt