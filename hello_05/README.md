练习目标
=======
go 自带 测试框架

Golang单元测试对文件名和方法名，参数都有很严格的要求。  
1、文件名必须以xx_test.go命名  
2、方法必须是Test[^a-z]开头  
3、方法参数必须 t *testing.T  

对于 tools 下的 utils.go，增加 utils_test.go

```go
package tools

import "testing"

func Test_Sum(t *testing.T) {
    x := Sum(2, 3)
    expect := 6
    if x != expect {
        t.Errorf("got [%d] expected [%d]", x, expect)
    } 
}
```

```sh
hzg@gofast:~/work/helloGo/hello_05/tools$ go test
PASS
ok      hello_05/tools  0.002s
hzg@gofast:~/work/helloGo/hello_05/tools$ go test
--- FAIL: Test_Sum (0.00s)
    utils_test.go:9: got [5] expected [6]
FAIL
exit status 1
FAIL    hello_05/tools  0.001s
```