#!/bin/bash
newDirectory="./tests/coverage"

# Create the coverage directory if it doesn't exist
mkdir -p "$newDirectory"

go test github.com/tweedledo/tests/integration
go test github.com/tweedledo/tests/unit/domain -coverpkg="github.com/tweedledo/core/domain" -coverprofile="./tests/coverage/coverage-domain.out"
go test github.com/tweedledo/tests/unit/service -coverpkg="github.com/tweedledo/core/service" -coverprofile="./tests/coverage/coverage-service.out"

coverageDomainFile="$newDirectory/coverage-domain.out"
coverageServiceFile="$newDirectory/coverage-service.out"
combinedCoverageFile="$newDirectory/coverage.out"

# Copy lines from line 2 to the end of the source file and save to a temporary file
tail -n +2 "$coverageDomainFile" > "$newDirectory/temp.txt"
tail -n +2 "$coverageServiceFile" > "$newDirectory/temp2.txt"
cat "$newDirectory/temp2.txt" >> "$newDirectory/temp.txt"
mv "$newDirectory/temp.txt" "$combinedCoverageFile"

echo "Complete coverage in \"$combinedCoverageFile\" file"

# Generate and display an HTML coverage report
go tool cover -html="$combinedCoverageFile"

# Clean up temporary files
rm "$newDirectory/temp2.txt"