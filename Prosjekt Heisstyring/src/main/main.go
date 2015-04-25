package main

import ( 
	"fmt"
	"udp"
	"driver"
	"control"
	"runtime" 
	//"net"
	//"os"
	//"sort"
	"functions"
	"time"
)

func main() {
	fmt.Println("FINN ET BEDRE STED FOR RUNNING=0 I DAG")
	runtime.GOMAXPROCS(runtime.NumCPU())
	test := time.Now()
	fmt.Println(test)
	fmt.Println(test.Hour())
	fmt.Println(test.Second())
	fmt.Println(test.Minute())
	fmt.Println(udp.GetID())
		
	//floorChan := make(chan int)
	var Data udp.Data
	PrimaryChan := make(chan int)
	SlaveChan := make(chan int)
	SortChan := make(chan int)

	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
	return
	}
	udp.UdpInit(30169, 39998, 1024, &Data,PrimaryChan,SlaveChan, SortChan)
	fmt.Println("Ferdig med å initialisere")
	fmt.Println("Currentfloor: ", Data.Statuses[udp.GetIndex(udp.GetID(),&Data)].CurrentFloor)
	fmt.Println("GetINDEX: ",udp.GetIndex(udp.GetID(),&Data))
	//fmt.Println("Currentfloor: ", Data.Statuses[udp.GetIndex(udp.GetID(),&Data)].CurrentFloor)
	fmt.Println("Test: ", udp.GetIndex(udp.GetID(),&Data))
	fmt.Println("Currentfloor[0]: ", Data.Statuses[0].CurrentFloor)
	//Status.ID = udp.GetID()	
	fmt.Println("Getfloor", driver.GetFloorSensorSignal())	


		
	//PrimaryChan<- 1
	//SlaveChan<-1
	fmt.Println("MIN INDEX ER: ", udp.GetIndex(udp.GetID(),&Data))
	Data.Lock.Lock()
	go control.GetDestination(&(Data.Statuses[udp.GetIndex(udp.GetID(),&Data)]))
	Data.Lock.Unlock()
	Data.Lock.Lock()
	go control.ElevatorControl(&(Data.Statuses[udp.GetIndex(udp.GetID(),&Data)]))
	Data.Lock.Unlock()
	fmt.Println("index fra main: ", udp.GetIndex(udp.GetID(), &Data))
	if(Data.Statuses[udp.GetIndex(udp.GetID(), &Data)].Primary){
		if len(Data.Statuses) >1{	
			for {
				Data.Statuses[1].OrderList = append(Data.Statuses[1].OrderList,3)
				udp.SendOrderlist(&Data,1)
				Data.Statuses[1].OrderList = functions.UpdateList(Data.Statuses[1].OrderList,0)
			}
		}
		Data.Lock.Lock()
		go control.CostFunction(&Data)
		Data.Lock.Unlock()
		//go udp.CleanDeadSlaves(&Data)
	}

	for {
		fmt.Println("for loop")
		select {
			case <-PrimaryChan:
					Data.Statuses[udp.GetIndex(udp.GetID(), &Data)].Primary = true
					Data.Lock.Lock()
					go control.CostFunction(&Data) 
					Data.Lock.Unlock()
			case <-SlaveChan:
				
			case <- SortChan:
					if len(Data.PrimaryQ)  > 1{
						temp := functions.SortUp(Data.PrimaryQ[1:])
						Data.PrimaryQ = Data.PrimaryQ[:1]
						Data.PrimaryQ = append(Data.PrimaryQ, temp...)
						fmt.Println(Data.PrimaryQ)
					}
			//default:
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
