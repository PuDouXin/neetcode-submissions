func topKFrequent(nums []int, k int) []int {
 frequency := make(map[int]int)
 for _,n := range nums{
	frequency[n]++
 }
 arr := make([][2]int, 0, len(frequency))
 for key,value := range frequency{
	arr = append(arr, [2]int{key,value})
 }

 sort.Slice(arr, func(i,j int) bool{
	return arr[i][1]>arr[j][1]
 })

 res := make([]int,k)
 for i:=0;i<k;i++{
	res[i]=arr[i][0]
 }
 return res
}
