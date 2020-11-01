package main
import "fmt"
func main(){
    var n int
    var ss string
    var tt string
    fmt.Scan(&n, &ss)
    fmt.Println(n, ss)
    
    for i := 0; i < n; n-- {
        fmt.Println(ss)
        if rcomp(ss, n){
            tt = tt + string(ss[i])
            ss = ss[1:]
        } else {
            tt = tt + string(ss[n-i-1])
            ss = ss[:n-i-1]
            
        }
    }
    fmt.Println(tt)
}

func rcomp(ss string, n int) bool{
    for i := 0 ; i < n; n--{
        if ss[i] < ss[n-i-1] {
            return true
        } else if ss[i] > ss[n-i-1]{
            return false
        }
    }
    return true
}
