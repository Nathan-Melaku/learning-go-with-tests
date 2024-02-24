package array


func Sum(arr []int) int {
  accumulator := 0
  for _, num := range arr {
    accumulator += num
  }
  return accumulator
}
