package control

import ( 
	//"fmt"
	"udp"
	//"runtime"
	"time"
	"driver"
	//"control"
	//"os"
	"sort"
	
)
/*
func GoToFloor(button int,  floorChan chan int,data *udp.Data) {
	floor := <-floorChan
	if driver.GetFloorSensorSignal() == -1 {
		driver.SetMotorDirection(driver.DIRN_DOWN)
	}
	var done int
	temp:= floor	
	//polse:
	for {		
		select {
		
		case temp = <-floorChan:
			//fmt.Println("Her er temp: %d", temp)
			//fmt.Println("Her er DONE: %d", done)
			//if done == 1{
									
			//	floor = temp
			//	done = 0
			//}

		default:			

			driver.SetFloorIndicator(driver.GetFloorSensorSignal())	
			if done == 1{
				//fmt.Printf("GAA IN EHFE")				
				floor = temp
				done = 0
				
				
			}	
			//fmt.Printf("Hva er done? %d\n",done)
			driver.SetButtonLamp(button,floor,1)
			//fmt.Printf("Her er flooooooooor: %d\n", floor)
				
			if floor == driver.GetFloorSensorSignal()  {
				
				fmt.Println("Framme på:", floor)
				udp.SetStatus(data,0,floor)

				driver.SetDoorOpenLamp(true)				
				driver.SetMotorDirection(driver.DIRN_STOP)
				time.Sleep(1*time.Second)
				driver.SetDoorOpenLamp(false)
				driver.SetFloorIndicator(floor)
				driver.SetButtonLamp(button,floor,0)
				done = 1
				
				//temp = -1
				//driver.SetDoorOpenLamp(false)	
				//fmt.Println("Done: %d", done)
				break
		
			} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			udp.SetStatus(data,2, floor)
			driver.SetMotorDirection(driver.DIRN_UP) 
		
			} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			udp.SetStatus(data ,1, floor)
			driver.SetMotorDirection(driver.DIRN_DOWN)
			}

		}
	
	}
}
*/
func GoToFloor(floor int) {
	
	for {
		driver.SetFloorIndicator(driver.GetFloorSensorSignal())
		if floor == driver.GetFloorSensorSignal() {
				driver.SetFloorIndicator(floor)
				driver.SetMotorDirection(driver.DIRN_STOP)
				driver.SetDoorOpenLamp(true)				
				time.Sleep(1*time.Second)
				driver.SetDoorOpenLamp(false)
				
				break
		} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			driver.SetMotorDirection(driver.DIRN_UP) 
		} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			driver.SetMotorDirection(driver.DIRN_DOWN)
		}
	}
}

func ElevatorControl(data *udp.Data){
	i := 0
	for {
		//empty := len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList) > 0
		if len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList) > 0 {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] > driver.GetFloorSensorSignal() {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] > data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i] {
					GoToFloor(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i])
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList,i)
					
					
				} else if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i] {
					GoToFloor(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i])
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList,i)
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,i)
					
					
				} else if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] < data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i] {
					GoToFloor(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i])
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,i)
					
					
				}
			} else if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] < driver.GetFloorSensorSignal() {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = -1
				if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] < data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i] {
					GoToFloor(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i])
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList,i)
					
					
				} else if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i] {
					GoToFloor(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i])
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList,i)
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,i)
					
					
				} else if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] > data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CommandList[i] {
					GoToFloor(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i])
					UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,i)
					
					
				}
			} else if 	data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[i] == driver.GetFloorSensorSignal() {
				UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,i)
				
			}
		}
	
	
		
	}
}	
	
		
func GetDestination(status *udp.Status) { //returnerer bare button, orderlist oppdateres
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
				 if(driver.GetButtonSignal(0,floor) == 1 && status.UpList[len(status.UpList)-1] != floor) {
					status.UpList = append(status.UpList,floor)
				
				} else if driver.GetButtonSignal(1,floor) == 1  && status.DownList[len(status.DownList)-1] != floor {
					status.DownList = append(status.DownList,floor)
			
				} else if driver.GetButtonSignal(2,floor) == 1  && status.CommandList[len(status.CommandList)-1] != floor {
					status.CommandList = append(status.CommandList, floor)
					/*
					if status.Running == 0 {
						status.OrderList = append(status.OrderList, floor)
						// tenne lampe?
					} else if status.Running == 1 {
						if floor < status.OrderList[len(status.OrderList)-1] && floor > status.OrderList[0] {
					}*/
					
				}
			
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}

}
/*
func GetCommand() (int,int) {
	button := 2	
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			if(driver.GetButtonSignal(button,floor) == 1) {
				return button,floor
			}
			
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
return -1,-1
}
*/

func CostFunction(data *udp.Data) {
	var DownList []int
	var UpList []int
	for {
	for k := 0; k < len(data.PrimaryQ);k++ {
		DownList = append(DownList,data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList...)
	}
	
	for l := 0; l < len(data.PrimaryQ);l++ {
		UpList = append(UpList,data.Statuses[udp.GetIndex(data.PrimaryQ[l], data)].UpList...)
	}
	handled := 0
	
	UpList = SortUp(UpList)
	DownList = SortDown(DownList)
	
	
	for down:=0; down<len(DownList);down++ {
		
		handled = 0
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i samme floor, og står stille
			if DownList[down] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				
				DownList = UpdateList(DownList,down) //Må modifiseres
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				
				handled = 1
				break
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og på veg nedover
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == DownList[down]+1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og står stille
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == DownList[down]+1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg nedover
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor > DownList[down] && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1  && handled != 1{
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				DownList = UpdateList(DownList,down)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg oppover, men siste skal stoppe på denne etasjen
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] == DownList[down] { 
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				handled = 1
				break 
			}	
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg oppover, men siste stopp er under
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] < DownList[down] {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				handled = 1
				break 
			}
		}

		if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,DownList[down])
			data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList)
			DownList = UpdateList(DownList,down)
			handled = 1 
		}
	}

for up:=0; up<len(UpList);up++ {
		
		handled = 0
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if UpList[up] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				UpList = UpdateList(UpList,up) //Må modifiseres
				handled = 1
				break 
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == UpList[up]-1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				UpList = UpdateList(UpList,up)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == UpList[up]-1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				UpList = UpdateList(UpList,up)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor < UpList[up] && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1  && handled != 1{
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				UpList = UpdateList(UpList,up)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] == UpList[up] { 
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				
				UpList = UpdateList(UpList,up)
				handled = 1
				break 		
			}	
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] > UpList[up] {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				
				UpList = UpdateList(UpList,up)
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				
				UpList = UpdateList(UpList,up)
				handled = 1
				break 
			}
		}

		if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,UpList[up])
			UpList = UpdateList(UpList,up)
			handled = 1 
		}
	}
	}
}



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

func SortUp(UpList []int)  []int{
	sort.Ints(UpList)
	temp := make([]int,1)
	temp[0] = UpList[0]
	counter := 0
	for i:= 1;i<len(UpList); i++ {
		if UpList[i] > temp[counter] {
			counter ++
			temp = append(temp,UpList[i])
		}
	}
	return temp
}	

func SortDown(DownList []int)  []int{
	sort.Ints(DownList)
	temp := make([]int,1)
	temp[0] = DownList[len(DownList)-1]
	counter := 0
	for i:= (len(DownList)-1); i>=0; i-- {
		
		if DownList[i] < temp[counter] {
			counter ++
			temp = append(temp,DownList[i])
		}
	}
	return temp
} 
	





