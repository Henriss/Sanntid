package control

import ( 
	"fmt"
	"udp"
	//"runtime"
	"time"
	"driver"
	//"control"
	//"os"
)

func GoToFloor(floorChan chan int,Status *udp.Status) {
	floor := <-floorChan
	button := Status.Button	
	fmt.Println("Floor: ",floor)
	fmt.Println("Button: ",button)	
	
	if driver.GetFloorSensorSignal() == -1 {
		driver.SetMotorDirection(driver.DIRN_DOWN)
	}
	var done int
	temp:= floor	
	//polse:
	for {	
		/*if driver.GetStopSignal() != 0 {
			driver.SetMotorDirection(driver.DIRN_STOP)
			fmt.Println("Stop button pressed")			
			os.Exit(2)	
			} */	
		select {
		
		case temp = <-floorChan:
			//fmt.Println("Her er temp: %d", temp)
			//fmt.Println("Her er DONE: %d", done)
			//if done == 1{
									
			//	floor = temp
			//	done = 0
			//}

		default:			
			/*if driver.GetStopSignal() != 0 {
				driver.SetMotorDirection(driver.DIRN_STOP)
				os.Exit(3)	
			}*/		
			driver.SetFloorIndicator(driver.GetFloorSensorSignal())	
			if done == 1{
				//fmt.Printf("GAA IN EHFE")				
				floor = temp
				done = 0
				
				
			}	
			//fmt.Printf("Hva er done? %d\n",done)
			driver.SetButtonLamp(Status.Button,floor,1)
			//fmt.Printf("Her er flooooooooor: %d\n", floor)
				
			if floor == driver.GetFloorSensorSignal()  {
				
				fmt.Println("Framme på:", floor)
				udp.SetStatus(Status,0,floor,button)

				driver.SetDoorOpenLamp(true)				
				driver.SetMotorDirection(driver.DIRN_STOP)
				time.Sleep(1*time.Second)
				driver.SetDoorOpenLamp(false)
				driver.SetFloorIndicator(floor)
				driver.SetButtonLamp(Status.Button,floor,0)
				done = 1
				
				//temp = -1
				//driver.SetDoorOpenLamp(false)	
				//fmt.Println("Done: %d", done)
				break
		
			} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 {
			udp.SetStatus(Status,2, floor,button)
			driver.SetMotorDirection(driver.DIRN_UP) 
		
			} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 {
			udp.SetStatus(Status,1, floor,button)
			driver.SetMotorDirection(driver.DIRN_DOWN)
			}

		}
	
	}
}

func GetDestination(Status *udp.Status) {	
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			for button := 0; button < 2; button++ {
				 if(driver.GetButtonSignal(button,floor) == 1) {
					udp.SetStatus(Status,-1,floor,button)
				}
			}
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
//udp.SetStatus(Status,-1,driver.GetFloorSensorSignal(), -1)
}

func GetCommand(Status *udp.Status) {
	
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			if(driver.GetButtonSignal(2,floor) == 1) {
				udp.SetStatus(Status,-1,floor,2)
			}
			
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
//udp.SetStatus(Status,-1,driver.GetFloorSensorSignal(), -1)
}








