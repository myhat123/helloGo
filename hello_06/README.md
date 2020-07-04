练习目标
=======
第三方包集成

go第三方包下载代理

https://goproxy.io/zh/

Go 版本是 1.13 及以上

```sh
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

```sh
hzg@gofast:~/work/helloGo/hello_06/tools$ go env -w GO111MODULE=on
hzg@gofast:~/work/helloGo/hello_06/tools$ go env -w GOPROXY=https://goproxy.io,direct
hzg@gofast:~/work/helloGo/hello_06/tools$ go test
go: finding module for package github.com/stretchr/testify/assert
go: downloading github.com/stretchr/testify v1.5.1
go: found github.com/stretchr/testify/assert in github.com/stretchr/testify v1.5.1
go: downloading gopkg.in/yaml.v2 v2.2.2
go: downloading github.com/davecgh/go-spew v1.1.0
go: downloading github.com/pmezard/go-difflib v1.0.0
PASS
ok      hello_05/tools  0.003s
```

go.mod 文件
===========

    module hello_05

    go 1.14

    require github.com/stretchr/testify v1.5.1

新增了go.sum 文件

```
github.com/davecgh/go-spew v1.1.0 h1:ZDRjVQ15GmhC3fiQ8ni8+OwkZQO4DARzQgrnXU1Liz8=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.5.1 h1:nOGnQDM7FYENwehXlg/kFVnos3rEvtKTjRvOWSzb6H4=
github.com/stretchr/testify v1.5.1/go.mod h1:5W2xD1RspED5o8YsWQXVCued0rvSQ+mT+I5cxcmMvtA=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v2 v2.2.2 h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
```

下载这些包的位置存放
=================

```sh
export GOBIN=~/go/bin
```

放在 ~/go/pkg/mod 中

了解这些，主要是为了代码迁移