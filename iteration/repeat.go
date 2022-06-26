package iteration

func Repeat(chracter string, repeatCount int) string {
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += chracter
    }
    return repeated
}
