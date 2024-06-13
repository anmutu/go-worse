# go-worse

本项目，有可能让您的项目变得不简洁，不易懂。需要谨慎使用。

## 示例1

`ternary.go`实现了一个泛型函数 `Ternary`，用于在 Go 语言中模拟三元表达式的功能。

### 介绍

许多编程语言都支持三元表达式，用于简洁地执行条件判断并返回结果。Go 语言虽然没有内置的三元表达式语法，但我们可以通过定义泛型函数来实现类似的功能。本项目展示了如何使用 Go 语言的泛型特性来实现一个简单的三元表达式函数。

### 使用示例

第一步，获取库。

`go get github.com/anmutu/go-worse`

第二步，调用方法。你看跟原始的相比还是变麻烦了点吧。

```go
package main

import (
   "fmt"
   "github.com/yourusername/go-ternary"
)

func main() {
    result := ternary.Ternary(1 > 0, true, false)
    fmt.Println(result) // 输出: true
}
```
