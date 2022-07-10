# go-cloud-disk


## How to use goctl

### install

```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### quick start 

```bash
goctl api new core
```

### according api file to generate api

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
    // 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
    // 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
    u, _ := url.Parse("https://examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            // 通过环境变量获取密钥
            // 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
            SecretID: os.Getenv("SECRETID"),
            // 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
            SecretKey: os.Getenv("SECRETKEY"),
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