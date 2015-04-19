package main

import ( 
	"fmt"
	"udp"
	"driver"
	"control"
	"runtime" 
	//"net"
	//"os"
)

func main() {
	
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(udp.GetID())
		
	//floorChan := make(chan int)
	//var Status udp.Status
	var Data udp.Data
	PrimaryChan := make(chan int)
	SlaveChan := make(chan int)
	//go udp.ChannelFunc(PrimaryChan)
		

	//Data := make(map[int]udp.Status)
	//var PrimaryQ [3]string

	udp.UdpInit(30169, 39998, 1024, &Data,PrimaryChan,SlaveChan)
	fmt.Println("Ferdig med å initialisere")
	//Status.ID = udp.GetID()	
	fmt.Println("Getfloor", driver.GetFloorSensorSignal())	

	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
		
	//PrimaryChan<- 1
	//SlaveChan<-1
		
	go control.GetDestination(&(Data.Statuses[udp.GetIndex(udp.GetID(),&Data)]))
	go control.ElevatorControl(&Data)
	go control.CostFunction(&Data)
	
	
	for {
		//fmt.Println("for loop")
		select {
			case <-PrimaryChan:
				if len(Data.PrimaryQ)  > 1{
					temp := control.SortUp(Data.PrimaryQ[1:])
					Data.PrimaryQ = Data.PrimaryQ[:1]
					Data.PrimaryQ = append(Data.PrimaryQ, temp...)
					fmt.Println(Data.PrimaryQ)
				}
			case <-SlaveChan:
			
			default:
				//fmt.Println("default case")
		}
	}
	

	
	
	
	
	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	//if Status.Primary == true {
	//	go udp.Send()
	//} else {
	//	go udp.Listen()
	//}	
		
	//go control.GoToFloor(2,floorChan,&Data)
	
	/*
	for {
		//_, temp := control.GetCommand()
		//floorChan<- temp
		//PrintStatus(Data.Status)
		fmt.Println("Stop signal pressed ", driver.GetStopSignal())
		if driver.GetStopSignal() != 0 {
			fmt.Println("Stop signal pressed ", driver.GetStopSignal())			
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	
	}
	*/
}		 

func PrintStatus(Status udp.Status) {
	fmt.Println("Running: ", Status.Running)
	fmt.Println("CurrentFloor: ", Status.CurrentFloor)
	fmt.Println("NextFloor: ", Status.NextFloor)
	fmt.Println("Primary: ", Status.Primary)
	fmt.Println("ID: ", Status.ID)
	fmt.Println("OrderList: ", Status.OrderList)
}
