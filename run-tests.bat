@echo off
call go test .\tests\unit\domain\tasklist_test.go -coverpkg="github.com/tweedledo/core/domain" -coverprofile=".\tests\coverage\coverage-domain.out"
call go test .\tests\unit\service\tasklist_service_test.go -coverpkg="github.com/tweedledo/core/service" -coverprofile=".\tests\coverage\coverage-service.out"

set "sourceFile=.\tests\coverage\coverage-domain.out"
set "appendFile=.\tests\coverage\coverage-service.out"
set "thirdFile=.\tests\coverage\coverage.out"

:: Copy lines from line 2 to the end of the source file and save to a temporary file
more "%sourceFile%" > temp.txt
more +1 "%appendFile%" > temp2.txt

:: Append the content of the append file to the temporary file
type temp2.txt >> temp.txt

:: Rename the temporary file to the third file
move temp.txt "%thirdFile%"

echo Complete covarage in "%thirdFile%" file

call go tool cover -html=%thirdFile%
