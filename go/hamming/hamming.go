package hamming

import "errors"

func Distance(a, b string) (count int, err error) {

    for i := 0; i < len(a); i++ {
        if (len(a) != len(b)) || (len(a) == 0)  || (len(b) == 0){
            return 0, errors.New("strings of unequal length")
        }
        if a[i] != b[i]{
            count++
        }
    }

    return
}
