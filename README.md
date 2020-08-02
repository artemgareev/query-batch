# query-batch

![build](https://github.com/artemgareev/query-batch/workflows/build/badge.svg?branch=master&event=push)
![test](https://github.com/artemgareev/query-batch/workflows/test/badge.svg?branch=master&event=push)

Gives you desired *LIMIT* and *OFFSET* values depending on *BATCH SIZE* and *RECORDS NUMBER*

## Installing

```go get -u github.com/artemgareev/query-batch```

## Example

```go
package main

import (
	querybatch "github.com/artemgareev/query-batch"
)

func main() {
	_ = querybatch.GetQueryBatches(100, 300)
	//output
	//[]querybatch.QueryBatches{
	//	querybatch.QueryBatches{Limit:100, Offset:0},
	//	querybatch.QueryBatches{Limit:100, Offset:100},
	//	querybatch.QueryBatches{Limit:100, Offset:200}
	//}

	_ = querybatch.GetQueryBatchesWithOffset(100, 300, 100)
	//output
	//[]querybatch.QueryBatches{
	//	querybatch.QueryBatches{Limit:100, Offset:100},
	//	querybatch.QueryBatches{Limit:100, Offset:200}
	//}
}
```