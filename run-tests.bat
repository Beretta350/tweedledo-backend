@echo off

set "newDirectory=.\tests\coverage"
if not exist "%newDirectory%" (
    mkdir "%newDirectory%"
)

go test github.com\tweedledo\tests\integration
go test github.com\tweedledo\tests\unit\domain -coverpkg="github.com/tweedledo/core/domain" -coverprofile=".\tests\coverage\coverage-domain.out"
go test github.com\tweedledo\tests\unit\service -coverpkg="github.com/tweedledo/core/service" -coverprofile=".\tests\coverage\coverage-service.out"

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

del temp2.txt
