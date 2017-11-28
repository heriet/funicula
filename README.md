# funicula

WIP

go SDK for NIFCLOUD


## Installation

```
go get github.com/heriet/funicula/nifcloud
go get github.com/heriet/funicula/service/rdb
```

## Example

```
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/heriet/funicula/nifcloud"
	"github.com/heriet/funicula/nifcloud/credential"
	"github.com/heriet/funicula/service/rdb"
)

func main() {

	config := &nifcloud.Config{
		Region: "east-1",
		Credential: &credential.Credential{
			AccessKeyId:     "<YOUR ACCESS KEY>",
			SecretAccessKey: "<YOUR SECRET KEY>",
		},
	}

	svc := rdb.New(config)

	input := &rdb.NiftyGetMetricStatisticsInput{
		Dimensions: []rdb.Dimension{
			{
				Name:  "DBInstanceIdentifier",
				Value: "mydbinstance",
			},
		},
		MetricName: "CPUUtilization",
	}

	output, err := svc.NiftyGetMetricStatistics(input)
	if err != nil {
		log.Fatal(err)
	}

	datapoints := output.Datapoints
	j, err := json.Marshal(datapoints)
	os.Stdout.Write(j)
}
```

## Supported Service

### RDB

- NiftyGetMetricStatistics


## License

This project is distributed under the Apache License, Version 2.0, see LICENSE.txt.