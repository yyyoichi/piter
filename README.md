# piter

Pipeline and Fun-Out/In patterns using the iter package

## install

```shell
go get "github.com/yyyoichi/piter"
```

## How to use

```golang
func main() {
    var strSrc iter.Seq[string] = func(yield func(string) bool) {
        _ = yield("1")
        _ = yield("2")
        _ = yield("a")
        _ = yield("10")
    }
    var ns = make([]int, 0, 4)
    for n, err := range piter.Pipeline12(ctx, strSrc, strconv.Atoi) {
        if err != nil {
            break
        }
        ns = append(ns, n)
    }
    for _, n := range ns {
        fmt.Println(n)
    } 

    // Output:
    // 1
    // 2
}

```
