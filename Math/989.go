/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
package math

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

// A + K
// func addToArrayForm(A []int, K int) []int {
// 	n := len(A)
// 	k := 1
// 	buf := 0
// 	stk := make([]int, 0, cap(A)+1)

// 	for i := n - 1; i >= 0; i-- {

// 		// fetch the addition
// 		add := K / k % 10
// 		bit := A[i] + add + buf
// 		if bit >= 10 {
// 			buf = 1
// 			bit %= 10
// 		} else {
// 			buf = 0
// 		}

// 		A[i] = bit
// 		stk = append(stk, A[i])
// 		k *= 10
// 	}

// 	for bit := K / k % 10; K / k != 0; bit = K/k %10 {
// 		if buf == 1 {
// 			bit =(bit + buf)%10
// 			buf = 0
// 			continue
// 		}
// 		stk = append(stk, bit)
// 		k *= 10
// 	}
// 	if buf == 1 {
// 		stk = append(stk, buf)
// 	}
// 	reverse(0,len(stk)-1,&stk)
// 	// consider len(K) > len(A)

// 	return stk

// }


func addToArrayForm(A []int, K int) (ans []int) {
    for i := len(A) - 1; i >= 0; i-- {
        sum := A[i] + K%10
        K /= 10
        if sum >= 10 {
            K++
            sum -= 10
        }
        ans = append(ans, sum)
    }
    for ; K > 0; K /= 10 {
        ans = append(ans, K%10)
    }
    reverse(ans)
    return
}

func reverse(A []int) {
    for i, n := 0, len(A); i < n/2; i++ {
        A[i], A[n-1-i] = A[n-1-i], A[i]
    }
}


// Representation
// E.g.	134
//	- K / 1 % 10  1
//	- K / 10 % 10 3
//	- K / 100 % 10 4
// 1. not overflow
// 		calc bitwise addition
// 2. overflow
// 		buf = 1
//		num[i] + buf + add

//		res %= 10
// 3. highest overflow
//		new an array
// 4. consider: [0]  k=23
//		new
