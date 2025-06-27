package lo

import "fmt"
import lo "github.com/samber/lo"


func HelloDash(){

    //--map-- ，
    new_list := lo.Map([]int64{1, 2, 3, 4}, func(x int64, index int) int64 {
        return x * 2
    })
    fmt.Println(new_list);


    //--forEach--
    lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
        fmt.Println(x)
    })


    //--filter--
    even := lo.Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
        return x%2 == 0
    })
    fmt.Println(even)

    //--find--
    str, ok := lo.Find([]string{"a", "b", "c", "d"}, func(i string) bool {
        return i == "b"
    })
    fmt.Println(str,ok)

    //--findIndexOf--
    str, index, ok := lo.FindIndexOf([]string{"a", "b", "a", "b"}, func(i string) bool {
        return i == "b"
    })
    fmt.Println(str,index,ok)
    

    //--reduce--
    sum := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
        return agg + item
    }, 0)
    fmt.Println(sum)


}