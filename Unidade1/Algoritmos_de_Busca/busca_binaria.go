package main


func bin_search_int(val int, list []int) int{
	start := 0
	end := len(list) - 1
	meio:= int((end + start)/2)

	for start <= end{

		if val == list[meio]{
			return val
		} else if val > list[meio]{
			start = meio + 1

		} else if val < list[meio]{
			end = meio - 1
		}

		meio = int((end + start)/2)
	}

	return -1


}

func bin_search_rec(val int, list []int, start int, end int) int {
	meio := int((end + start)/2)

	if start > end{
		return -1

	} else if val == list[meio]{
		return val
	}else if val > list[meio]{
		start = meio + 1

	}else if val < list[meio]{
		end = meio - 1
	} 

	return bin_search_rec(val, list, start , end)

}

func main(){


}