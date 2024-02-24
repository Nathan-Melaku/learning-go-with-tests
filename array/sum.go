package array


func Sum(arr []int) int {
  accumulator := 0
  for _, num := range arr {
    accumulator += num
  }
  return accumulator
}

func SumAll(numberToSum ...[]int) []int {
  var sums []int
  for _, numbers := range numberToSum {
    sums = append(sums, Sum(numbers))
  }

  return sums 
}
