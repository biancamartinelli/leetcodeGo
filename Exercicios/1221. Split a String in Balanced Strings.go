func balancedStringSplit(s string) int {
    r := 0
    l := 0

    count := 0
    for _, v := range s {
        if v == 'R'{
            r++
        }else{
            l++
        }
        if r != 0 && r == l {
            count++
            r = 0
            l = 0
        }
    }

    return count 
}