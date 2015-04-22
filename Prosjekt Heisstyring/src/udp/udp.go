// go run networkUDP.cd ..go
package udp
import (."fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	//"runtime"
	"time"
	"net"
	//"bufio"
	"os"
	"strconv"
	"driver"
	//"sort"
	"encoding/json"

)


type Status struct {
	Running int
	CurrentFloor int
	NextFloor int
	Primary bool
	ID int
	//PrimaryQ [3]string
	CommandList []int
	UpList []int  // slice = slice[:0] for å tømme slicen når sendt til primary
	DownList[]int // slice = slice[:0] for å tømme slicen når sendt til primary
	//PriList[]int
	OrderList []int // sjekke for nye ordrer når primary sender
}

type Data struct {
	//Status Status
	//Timestamp???????
	ID int
	Statuses []Status // Oppdatere den her å i UdpInit()
	PrimaryQ []int
}


func SetStatus(status *Status, running int, NextFloor int) {
	ID := GetID()
	status.Running = running
	status.CurrentFloor = driver.GetFloorSensorSignal()
	status.NextFloor = NextFloor
	status.ID = ID
	
	/*
	(*data).Statuses[GetIndex(GetID(), data)].Running = running
	(*data).Statuses[GetIndex(GetID(), data)].CurrentFloor = driver.GetFloorSensorSignal()
	(*data).Statuses[GetIndex(GetID(), data)].NextFloor = NextFloor
	(*data).Statuses[GetIndex(GetID(), data)].ID = ID
	//Println(" id i func:", (*Status).ID)
	*/
}
func GetID() int {
	ut:=0
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
 	var ipAddr string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddr = ipnet.IP.String()
			}
		}
	}
	if len(ipAddr)==14{ 
		ut,_ = strconv.Atoi(ipAddr[12:14])
	}else{
		ut,_ = strconv.Atoi(ipAddr[12:15])
	}
	return ut
	
}


/////////// Primary functions ////////////

func PrimaryBroadcast(baddr *net.UDPAddr, data *Data) { // IMALIVE, oppdatere backup for alle
	//var temp Data
	//udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")
	//checkError(err)
	bconn, err := net.DialUDP("udp", nil, baddr)
	checkError(err)
	for {
		Println("SENDER")
		// WRITE
		b,_ := json.Marshal(data)
		bconn.Write(b)
		//json.Unmarshal(b[0:len(b)], temp) 
		//Println("b: ", b)
		//Println("PrimaryQ marshalled: ", len(temp.Statuses))
		checkError(err)
		time.Sleep(500*time.Millisecond)
	}

}

func SendOrderlist(data *Data) { // IMALIVE
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")
	bconn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	// WRITE
	b,_ := json.Marshal(data) // nok å bare sende en gang?
	bconn.Write(b)		
	checkError(err)
}

func PrimaryListen(data *Data, SortChan chan int) {
	buffer := make([]byte, 1024)
	temp := data
	udpAddr, err := net.ResolveUDPAddr("udp", ":39999")
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {	
		Println("HØRER")	
		n, err := conn.Read(buffer)
		checkError(err)
		//Data = buffer
		err = json.Unmarshal(buffer[0:n], temp)
		if temp.PrimaryQ[len(temp.PrimaryQ)-1] != data.PrimaryQ[len(temp.PrimaryQ)-1] {
			data.PrimaryQ = append(data.PrimaryQ, temp.PrimaryQ[1:]...)
			SortChan<- 1	
		}
		data.Statuses[temp.ID] = temp.Statuses[temp.ID]
		//(*data).Statuses[GetIndex((*data).Status.ID,data)] = (*data).Status // Oppdaterar mottatt status hos primary 
	}
}

/////////// Slave functions //////////// 

func ListenForPrimary(bconn *net.UDPConn, baddr *net.UDPAddr, data *Data, PrimaryChan chan int, SlaveChan chan int) { // Bruke chan muligens fordi den skal skrive til Data
	buffer := make([]byte, 1024)
	//udpAddr, err := net.ResolveUDPAddr("udp", ":39998")
	//conn, err := net.ListenUDP("udp", udpAddr)
	//checkError(err)
	for {
		Println("Hører")
		bconn.SetReadDeadline(time.Now().Add(5*time.Second))		
		n, err := bconn.Read(buffer)
		if err != nil && data.PrimaryQ[1] == GetID() {
			Println("Mottar ikke meldinger fra primary lenger, tar over")
			data.PrimaryQ = UpdateList(data.PrimaryQ,0)
			data.Statuses = data.Statuses[1:]
			go PrimaryBroadcast(baddr, data)
			// SendOrderlist(Data)
			PrimaryChan<- 1
			break
		}
		//Data = buffer
		err = json.Unmarshal(buffer[0:n], data)		
		Println("her er primaryQen:", data.PrimaryQ)		
		// Printf("Rcv %d bytes: %s\n",n, buffer)	
	}	
}



func SlaveUpdate(data *Data) { // chan muligens, bare oppdatere når det er endringar
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187."+ strconv.Itoa((*data).PrimaryQ[0]) + ":39999")
	conn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	for {
		 //WRITE
		b,_ := json.Marshal(data)
		data.Statuses[GetIndex(GetID(), data)].UpList = data.Statuses[GetIndex(GetID(), data)].UpList[:0]
		data.Statuses[GetIndex(GetID(), data)].DownList = data.Statuses[GetIndex(GetID(), data)].DownList[:0]
		conn.Write(b)	
		checkError(err)
		time.Sleep(2500*time.Millisecond) // bytte til bare ved endringar etterhvert
		if data.Statuses[GetIndex(GetID(), data)].Primary == true {
			break
		}
	}
}

// send_ch, receive_ch chan Udp_message
func UdpInit(localListenPort int, broadcastListenPort int, message_size int, data *Data, PrimaryChan chan int, SlaveChan chan int) (err error) {
	buffer := make([]byte, message_size)
	//var temp Data
	var status Status
	//data.Statuses = append(data.Statuses, temp)
	status.Primary = false
		
	SetStatus(&status,0,driver.GetFloorSensorSignal())	
	//InitStatus(*Status)
	//Println("SE HER::::: ", (Status).ID)
	
	//Generating broadcast address
	baddr, err = net.ResolveUDPAddr("udp4", "129.241.187.255:"+strconv.Itoa(broadcastListenPort))
	if err != nil {
		return err
	}

	//Generating localaddress
	tempConn, err := net.DialUDP("udp4", nil, baddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	laddr, err = net.ResolveUDPAddr("udp4", tempAddr.String())
	laddr.Port = localListenPort

	//Creating local listening connections
	localListenConn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}

	//Creating listener on broadcast connection
	broadcastListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {
		localListenConn.Close()
		return err
	}

	//go udp_receive_server(localListenConn, broadcastListenConn, message_size receive_ch)
	//go udp_transmit_server(localListenConn, broadcastListenConn ,send_ch)

	//Setting first primary
	broadcastListenConn.SetReadDeadline(time.Now().Add(3*time.Second))
	n, err := broadcastListenConn.Read(buffer)
	if err != nil {
		Println("Tar over som primary!")
		(*data).PrimaryQ = append((*data).PrimaryQ, GetID())
		(*data).Statuses = append((*data).Statuses, status)
		(*data).Statuses[GetIndex(GetID(), data)].Primary = true
		//PrimaryChan <- 1
		//go ChannelFunc(PrimaryChan)
		go PrimaryBroadcast(baddr,data)
		go PrimaryListen(data, PrimaryChan)
		
	
	} else {
		err = json.Unmarshal(buffer[0:n], data)
		//(*data).PrimaryQ = temp.PrimaryQ
		(*data).PrimaryQ = append((*data).PrimaryQ, GetID())
		//(*data).Statuses = temp.Statuses
		(*data).Statuses = append((*data).Statuses, status)	
		Println("PrimaryQ: ", data.PrimaryQ)
		//(*Data).PrimaryQ = append((*Data).PrimaryQ, string(buffer))
		//SlaveChan<- 1
		go ChannelFunc(SlaveChan)		
		go SlaveUpdate(data)
		time.Sleep(2500*time.Millisecond) // Vente for å la Primary oppdatere PrimaryQen
		go ListenForPrimary(broadcastListenConn, baddr, data,PrimaryChan, SlaveChan)
	}
	


	//	fmt.Printf("Generating local address: \t Network(): %s \t String(): %s \n", laddr.Network(), laddr.String())
	//	fmt.Printf("Generating broadcast address: \t Network(): %s \t String(): %s \n", baddr.Network(), baddr.String())
	return err
}

/*
func SendCommandList() { // Bare sende siste tal for simplicity
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (endre ip nettverket du sitter på)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}

	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("Sent: ",currentStruct.Teller) 		
		currentStruct.Teller = currentStruct.Teller + 1
		time.Sleep(1*time.Second)
	}
}*/





/*
func SendCommand(floorChan chan int) {
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (endre ip nettverket du sitter på)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}

	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("Sent: ",currentStruct.Teller) 		
		currentStruct.Teller = currentStruct.Teller + 1
		time.Sleep(1*time.Second)
	}

}*/

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err)
		return
	}
}

func GetIndex(ID int, data *Data) int { 
	for i:=0; i<len(data.PrimaryQ); i++ {
		if data.PrimaryQ[i] == ID {
			return i
		}
	}
	return -1
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


func ChannelFunc(Channel chan int) {
	Channel <-1
}






