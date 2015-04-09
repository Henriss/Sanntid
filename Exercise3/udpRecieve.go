// go run oving2_go.go
package main
import (. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
	."net" 
)

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err) //err.Error()
		return //os.exit(1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// recvSock = 
	buffer := make([]byte, 1024)
	udpAddr, err := ResolveUDPAddr("udp", ":30000") // Mulig net. trengs foran Dial
	checkError(err)
	conn, err := ListenUDP("udp", udpAddr)
	checkError(err)
	
	for {
		time.Sleep(1000*time.Millisecond)
		//Println("Hei!")
		n,err := conn.Read(buffer)
		checkError(err)
		Printf("Rcv %d bytes: %s\n",n, buffer)
	}
	

}
