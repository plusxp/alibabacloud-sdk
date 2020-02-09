English | [简体中文](Usage-CN.md)

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# Go SDK instructions

## Requirements
- It's necessary for you to make sure your system have installed a Go environment which is new than 1.12.x.

## Installation
If you use `go mod` to manage your dependence, you can use the following command:

```sh
$ go mod tidy
```

## Quick Examples
Before you begin, you need to sign up for an Alibaba Cloud account and retrieve your [Credentials](https://usercenter.console.aliyun.com/#/manage/ak).

## Demo
```go
import (
    "fmt"
    "strings"

    common "github.com/aliyun/alibabacloud-rpc-util-sdk/golang/common"
	facebody "github.com/aliyun/alibabacloud-sdk/facebody-20191230/golang/client"
)

func main(){
    // init config
    var config = new(facebody.Config).SetAccessKeyId("ACCESS_KEY_ID").
	SetAccessKeySecret("ACCESS_KEY_SECRET").
	SetType("access_key")

    // create client
    client, err := facebody.NewClient(config)
    if err != nil {
        panic(err)
    }

    // init runtimeObject
    var runtimeObject = new(common.RuntimeObject).SetAutoretry(false).
	SetMaxIdleConns(3)

    // init request
    var request := new(facebody.DetectFaceRequest).SetFaceCount(10)

    // call api
    resp, err := client.DetectFace(request, runtimeObject)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println(resp)


    // file upload
    var uploadRequest := new(facebody.DetectFaceAdvanceRequest).SetFaceCount(10).
       SetImageURLObject(strings.NewReader("demo"))

    // call api
    uploadResp, err := client.DetectFaceAdvance(request, runtimeObject)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println(uploadResp)
}
```

## Parameter Specification
```go
type Config struct {
	AccessKeyId          *string // AccessKey Id
	AccessKeySecret      *string // AccessKey Secret
	Type                 *string // credential type，If you have any questions, please refer to it https://github.com/aliyun/credentials-go/blob/master/README-CN.md#%E5%87%AD%E8%AF%81%E7%B1%BB%E5%9E%8B
	SecurityToken        *string // Security Token
	Endpoint             *string // endpoint
	Protocol             *string // http or https
	RegionId             *string // region
    UserAgent            *string // UserAgent
	ReadTimeout          *int    // read timeout
	ConnectTimeout       *int    // connect tiemout
	LocalAddr            *string // the local network adapter ip
	HttpProxy            *string // http proxy
	HttpsProxy           *string // https proxy
	NoProxy              *string // agent whitelist
	Socks5Proxy          *string // socks5 proxy
	Socks5NetWork        *string // socks5 agency agreement
	MaxIdleConns         *int    // maximum number of connections
	EndpointType         *string // ednpoint type：internal，accelerate or null
	OpenPlatformEndpoint *string // endpoint used when the file is uploaded(Not at the moment)
}

type RuntimeObject struct {
	Autoretry        *bool       // whether to enable retry
	IgnoreSSL        *bool       // whether SSL validation is ignored
	MaxAttempts      *int        // maximum number of retries， defaut: 3
	BackoffPolicy    *string     // retry the sleep strategy, default: no
	BackoffPeriod    *int        // retry the sleep time,default: 1
	ReadTimeout      *int        // read timeout
	ConnectTimeout   *int        // connect timeout
	LocalAddr        *string     // the local network adapter ip
	HttpProxy        *string     // http proxy
	HttpsProxy       *string     // https proxy
	NoProxy          *string     // agent whitelist
	MaxIdleConns     *int        // maximum number of connections
	Socks5Proxy      *string     // socks5 proxy
	Socks5NetWork    *string     // socks5 agency agreement
}
```