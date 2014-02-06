moviehash
========= 
This is a small go package for calculating movie hashes
used by Media Player Classic. For more information, see 
http://trac.opensubtitles.org/projects/opensubtitles/wiki/HashSourceCodes

## Install

``` shell
go get github.com/radenling/moviehash
```

## Usage
Import the library:

``` go
import "github.com/radenling/moviehash"
```

Call it using ComputeFileHash or ComputeHash:
``` go
hash, err := ComputeFileHash(filename)
```

## Tests

To run the tests you need two test files from [opensubtitles.org](http://trac.opensubtitles.org/projects/opensubtitles/wiki/HashSourceCodes)

Run the tests with the go tool:
``` shell
go test
```
