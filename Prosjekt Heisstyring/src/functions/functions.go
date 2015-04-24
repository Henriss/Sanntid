package functions
import (//"fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	//"runtime"
	"time"
	"math"
	//"bufio"
	//"os"
	//"strconv"
	//"driver"
	//"sort"
	//"encoding/json"
	"sort"

)




func UpdateList(OrderList []int, j int) []int {
	temp := make([]int, len(OrderList)-1)
	for i:= 0; i<len(OrderList);i++ {
		if i<j {
			temp[i] = OrderList[i]
		} else if i>j {
			temp[i-1] = OrderList[i]
		}
	}
	return temp
}





func SortUp(UpList []int)  []int{ //Sorterer listen UpList i stigende rekkefølge og fjerner like tall og -1
	sort.Ints(UpList)
	temp := make([]int,1)
	var minusen int
	if(UpList[0]==-1){
		temp[0] = UpList[1]
		minusen  = 2
	}else{
		temp[0] = UpList[0]
		minusen = 1
	}
	
	
	counter := 0
	for i:= minusen;i<len(UpList); i++ {
		if UpList[i] > temp[counter] {
			counter ++
			temp = append(temp,UpList[i])
		}
	}
	return temp
}
func CheckList(list []int, check int) bool{ // Sjekker om listen list inneholder heltallet check
	for i:=0;i<len(list);i++{
		if(list[i] == check){
			return true
		}
	}
	return false
}



func SortDown(DownList []int)  []int{
	sort.Ints(DownList)
	var minusen int
	temp := make([]int,1)
	if(DownList[0]==-1){
		minusen = 1
	}else{
		minusen = 0
	}
	temp[0] = DownList[len(DownList)-1]
	counter := 0
	for i:= (len(DownList)-1); i>=minusen; i-- {
		
		if DownList[i] < temp[counter] {
			counter ++
			temp = append(temp,DownList[i])
		}
	}
	return temp
} 

func Delay(SlaveTime time.Time, PrimeTime time.Time) int{
	temp := SlaveTime.Sub(PrimeTime)
	return int(math.Floor(temp.Seconds()))
}

