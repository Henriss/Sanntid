package control

import ( 
	"fmt"
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
func GoToFloor(floor int, status *udp.Status) {
	fmt.Println("control 82: går til floor floor:",floor)
	for {
		driver.SetFloorIndicator(driver.GetFloorSensorSignal())
		if floor == driver.GetFloorSensorSignal() {
				driver.SetFloorIndicator(floor)
				driver.SetMotorDirection(driver.DIRN_STOP)
				driver.SetDoorOpenLamp(true)				
				time.Sleep(2*time.Second)
				driver.SetDoorOpenLamp(false)
				
				break
		} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			driver.SetMotorDirection(driver.DIRN_UP) 
		} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			driver.SetMotorDirection(driver.DIRN_DOWN)
		}
		if driver.GetFloorSensorSignal() != -1{
			status.CurrentFloor = driver.GetFloorSensorSignal()
		}	
	}
}


func ElevatorControl(status *udp.Status){
	//time.Sleep(1*time.Second)
	temp := 0
	temp = temp + 0
	for {
		if driver.GetFloorSensorSignal() != -1 {
			status.CurrentFloor = driver.GetFloorSensorSignal()
		}
		//fmt.Println(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor) 
		//fmt.Println("control 109: OrderList",status.OrderList)
		//time.Sleep(1*time.Second)
		if len(status.OrderList) == 0 && len(status.CommandList) == 0 {
			status.Running = 0
		}
		if len(status.CommandList) == 0 {
			status.CommandList = append(status.CommandList, -1)
		}
		if len(status.OrderList)==0 {
			status.OrderList = append(status.OrderList, -1)
			
		}

		if len(status.OrderList) > 0 && len(status.CommandList)>0 {
			if !(status.OrderList[0] == -1 && status.CommandList[0] ==-1){
				//fmt.Println("OrderList: ", status.OrderList)
				// 
				if status.OrderList[0] > status.CurrentFloor  {
					//status.Running =1
					//fmt.Println(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					// Sjekker om heisens ordreliste
					if status.CommandList[0] == -1{
						temp = status.OrderList[0]
						status.OrderList = UpdateList(status.OrderList,0)
						GoToFloor(temp, status)
						temp = 0
					}else if status.OrderList[0] == -1{
						temp = status.CommandList[0]
						status.CommandList = UpdateList(status.CommandList,0)
						GoToFloor(temp, status)
						temp = 0
					}else if status.OrderList[0]>status.CommandList[0]{
						temp = status.CommandList[0]
						status.CommandList = UpdateList(status.CommandList,0)	
						GoToFloor(temp, status)
						temp = 0
					}else if status.CommandList[0]>status.OrderList[0]{
						temp = status.OrderList[0]
						status.OrderList = UpdateList(status.OrderList,0)
						GoToFloor(temp, status)
						temp = 0
					}else if status.OrderList[0] == status.CommandList[0]{
						temp = status.OrderList[0]
						status.CommandList=UpdateList(status.CommandList,0)
						status.OrderList=UpdateList(status.OrderList,0)
						GoToFloor(temp, status)
						temp = 0
					}
				}else if status.OrderList[0] < status.CurrentFloor{
					//status.Running = -1
					if status.CommandList[0] == -1 {
						temp = status.OrderList[0]
						status.OrderList = UpdateList(status.OrderList,0)
						GoToFloor(temp, status)
						temp = 0
					}else if status.OrderList[0] == -1 {
						temp = status.CommandList[0] 
						status.CommandList = UpdateList(status.CommandList,0)
						GoToFloor(temp, status)
						temp = 0
					}else if status.OrderList[0] < status.CommandList[0]{
						temp = status.CommandList[0]
						status.CommandList = UpdateList(status.CommandList,0)
						GoToFloor(temp, status)
						temp = 0
					}else if status.CommandList[0] < status.OrderList[0]{
						temp = status.OrderList[0]
						status.OrderList = UpdateList(status.OrderList,0)
						GoToFloor(temp, status)
						temp = 0
					}else if status.OrderList[0] == status.CommandList[0]{
						temp = status.OrderList[0]
						status.CommandList=UpdateList(status.CommandList,0)
						status.OrderList=UpdateList(status.OrderList,0)
						GoToFloor(temp, status)
						temp = 0						
					}						
				}else if 	status.OrderList[0] == driver.GetFloorSensorSignal() {
						status.OrderList=UpdateList(status.OrderList,0)
						GoToFloor(driver.GetFloorSensorSignal(), status)						
				}
			}
	
		}
		
	}
}	
	
		
func GetDestination(status *udp.Status) { //returnerer bare button, orderlist oppdateres
	//time.Sleep(1*time.Second)
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
				if driver.GetButtonSignal(0,floor) == 1 && len(status.UpList) == 0 {
					status.UpList = append(status.UpList, floor) 
				}else if driver.GetButtonSignal(0,floor) == 1 && len(status.UpList) > 0 {
					if udp.CheckList(status.UpList,floor) == false {
						status.UpList = append(status.UpList,floor)
						fmt.Println("status.UpList: ", status.UpList) 
					}				
				}else if driver.GetButtonSignal(1,floor) == 1 && len(status.DownList)==0 {	
					status.DownList = append(status.DownList, floor)
				}else if driver.GetButtonSignal(1,floor) == 1 && len(status.DownList) > 0 {
					if udp.CheckList(status.DownList,floor) == false {
						status.DownList = append(status.DownList,floor)
						fmt.Println("status.DownList: ", status.DownList)
					}
				}else if driver.GetButtonSignal(2,floor) == 1 && len(status.CommandList) == 0{
						if status.Running == 0 {
							if status.CurrentFloor < floor{
								status.Running = 1
								status.CommandList = append(status.CommandList,floor)
							}else if status.CurrentFloor > floor{
								status.Running = -1
								status.CommandList = append(status.CommandList,floor) 
							}else{
								status.Running = 0
							}
						}else{
							status.CommandList = append(status.CommandList, floor)
						}
				}else if driver.GetButtonSignal(2,floor) == 1  && status.CommandList[0] == -1 {
						if status.Running == 0{
							if status.CurrentFloor < floor{
								status.Running = 1
								status.CommandList[0] = floor
							}else if status.CurrentFloor > floor{
								status.Running = -1
								status.CommandList[0] = floor
							}else{
								status.Running = 0
							}
						}else{
							status.CommandList[0] = floor
						}
				}else if driver.GetButtonSignal(2,floor) == 1 && len(status.CommandList) > 0 {
					if status.CommandList[len(status.CommandList)-1] != floor {
						//status.CommandList = append(status.CommandList, floor)
						if status.Running == 1{
							if floor <= status.CurrentFloor{
								status.CommandList = SortUp(status.CommandList)
								status.CommandList = append(status.CommandList, floor)
							}else{
								status.CommandList = append(status.CommandList, floor)
								status.CommandList = SortUp(status.CommandList)
							}
						}else if status.Running == -1{
							if floor >= status.CurrentFloor{
								status.CommandList = SortDown(status.CommandList)
								status.CommandList = append(status.CommandList, floor)
							}else{
								status.CommandList = append(status.CommandList, floor)
								status.CommandList = SortDown(status.CommandList)
							}
						}
					}
				}
					/*
					if status.Running == 0 {
						status.OrderList = append(status.OrderList, floor)
						// tenne lampe?
					} else if status.Running == 1 {
						if floor < status.OrderList[len(status.OrderList)-1] && floor > status.OrderList[0] {
					}*/
					
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
	handled := 0
	var DownList []int
	var UpList []int
	for {
	//fmt.Println("control 243, handled: ",handled)
	handled = 0
	/*if len(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList) > 0 {
		if(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList[0] == -1){
			data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList = UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList, 0)
		}
	}*/
	for k := 0; k < len(data.PrimaryQ);k++ {
		if udp.GetIndex(data.PrimaryQ[k],data) != -1 {
			DownList = append(DownList,data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList...)
			data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList = data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList[:0]
			fmt.Println("Sjekk om DOWNLIST oppdateres riktig: ", DownList)
		}
	}
	
	
	for l := 0; l < len(data.PrimaryQ);l++ {
		if 	udp.GetIndex(data.PrimaryQ[l],data) != -1 {
			UpList = append(UpList,data.Statuses[udp.GetIndex(data.PrimaryQ[l], data)].UpList...)
			data.Statuses[udp.GetIndex(data.PrimaryQ[l], data)].UpList = data.Statuses[udp.GetIndex(data.PrimaryQ[l], data)].UpList[:0]
			fmt.Println("Sjekk om UPLIST oppdateres riktig: ", UpList)
		}
	}
	//fmt.Println("control 258: PrimaryQ i cost function: ", data.PrimaryQ)
	//fmt.Println("control 259: Down List i cost function: ", DownList)
	time.Sleep(1*time.Second)
	
	if len(UpList) > 0 {
		UpList = SortUp(UpList)
	}
	if len(DownList) > 0 {
		DownList = SortDown(DownList)
	}
	
	//fmt.Println("OrderList: ", data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList)
	
	//fmt.Println(DownList)
	for down:=0; down<len(DownList);down++ {
		
		handled = 0
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i samme floor, og står stille
			if DownList[down] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {

				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 0
				fmt.Println("control 280: Heis i samma floor og står stille. Downlist:", DownList)
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
				fmt.Println("control 370: Heis i etasjen over og på vei nedover. Downlist:", DownList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og står stille
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == DownList[down]+1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = -1
				fmt.Println("control 385: Heis i etasjen over og står stille")
				DownList = UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg nedover
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor > DownList[down] && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1  && handled != 1{
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				DownList = UpdateList(DownList,down)
				fmt.Println("control 398: Heis på vei nedover og er over floor. Downlist:", DownList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		/*for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg oppover, men siste skal stoppe på denne etasjen
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] == DownList[down] { 
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				fmt.Println("e")
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}	
		}*/
		/*for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg oppover, men siste stopp er under
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] < DownList[down] {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				fmt.Println("f")
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				DownList = UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}*/
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,DownList[down])
				fmt.Println("control 437: heis står stille generelt. Downlist:", DownList)
				if DownList[down] > data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				}else if DownList[down] < data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = -1
				}
				DownList = UpdateList(DownList,down)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}

		/*if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,DownList[down])
			data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList)
			DownList = UpdateList(DownList,down)
			fmt.Println("h")
			handled = 1 
		}*/
	}

for up:=0; up<len(UpList);up++ {
		//fmt.Println("Up: ",up)
		//fmt.Println(data.PrimaryQ)
		handled = 0
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if UpList[up] == data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				fmt.Println("control 387: heis i samme etasjen og står stille. UpList:", UpList)
				UpList = UpdateList(UpList,up) //Må modifiseres
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == UpList[up]-1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				fmt.Println("control 402: heis i etasjen under og på vei oppover. UpList:", UpList)

				UpList = UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor == UpList[up]-1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 && handled != 1 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				fmt.Println("control 417: heis i etasjen under og står stille. UpList:", UpList)
				UpList = UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor < UpList[up] && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 1  && handled != 1{
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
				fmt.Println("control 430: floor over heis.currentfloor og på vei oppover. UpList:", UpList)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].NextFloor = UpList[up]
				UpList = UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		/*		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] == UpList[up] { 
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				fmt.Println("m")
				UpList = UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 		
			}	
		}
		*/
		/*		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == -1 && data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList[len(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)-1] > UpList[up] {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				fmt.Println("n")
				UpList = UpdateList(UpList,up)
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		*/
		fmt.Println("Her er handled: ",handled)		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			fmt.Println("RUNNING RUNNING RUNNING",data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running  )
			if data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running == 0 {
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = append(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList,UpList[up])
				
				fmt.Println("control 473: heisen står stille. UpList:", UpList)
				UpList = UpdateList(UpList,up)
				data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				if UpList[up] > data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortUp(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = 1
				}else if UpList[up] < data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor{
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					data.Statuses[udp.GetIndex(data.PrimaryQ[i], data)].Running = -1
				}
				if i != 0 {
					udp.SendOrderlist(data) // , udp.GetIndex(data.PrimaryQ[i], data))
				}
				handled = 1
				break 
			}
		}
		/*
		if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,UpList[up])
			UpList = UpdateList(UpList,up)
			handled = 1
			fmt.Println("p") 
		}*/
	}
	}
}





func SortUp(UpList []int)  []int{
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




