# go-cloud-disk


## How to use goctl

### Installation

```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### Quick start 

```bash
goctl api new core
```

### Generate handler file and logic file according .api file

```bash
goctl api go core.api -dir . -style gozero
```

## File upload (Tencent Cloud COS)

https://console.cloud.tencent.com/cos

### Install SDK
```bash
go get -u github.com/tencentyun/cos-go-sdk-v5
```

### Upload Object Example

```go
package main

import (
    "context"
    "github.com/tencentyun/cos-go-sdk-v5"
    "net/http"
    "net/url"
    "os"
)

func main() {
    u, _ := url.Parse("https://examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID: "YOUR-COS-SECRET-ID",
            SecretKey: "YOUR-COS-SECRET-KEY",
        },
    })

    key := "exampleobject"

    _, _, err := client.Object.Upload(
        context.Background(), key, "localfile", nil,
    )
    if err != nil {
        panic(err)
    }
}
```