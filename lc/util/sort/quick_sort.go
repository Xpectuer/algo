package sort

// import "fmt"
import "math/rand"

// QuickSort is a
func QuickSort(A []int, i int, j int) []int {
  quickSort(A, i, j)
  return A
}

func quickSort(A []int, i int, j int) {
  if i >= j  {return}
  pivot := partitionRandom(A,i,j)

  //fmt.Println(pivot)

  quickSort(A,i,pivot-1)
  quickSort(A,pivot+1,j)
}

func partition(A []int,s int,t int) int {
  pivot := t
  count := s
  for i := s;i<t;i++ {
    if A[i] < A[pivot] {
      A[i], A[count] = A[count], A[i]
      count++
    }
  }
  A[pivot] , A[count] = A[count],  A[pivot]
  return count
}

func partitionRandom(A []int, s int, t int) int {
   // generate [s,t]
  pivot := rand.Intn(t-s) + s
  count := s
  A[pivot], A[t] = A[t], A[pivot]
  for i := s;i<t;i++ {
    if A[i] < A[pivot] {
      A[i], A[count] = A[count], A[i]
      count++
    }
  }
  A[pivot] , A[count] = A[count],  A[pivot]
  return count
}
