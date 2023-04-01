package main

import (
	//"clippy/brain"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const clippy = `/‾‾\
|  |
O  O
|| |/
|| ||
|\_/|
\___/`

func main() {
	// check args
	//goal := flag.String("goal", "Run a minecraft server", "What would you like clippy to help you with today?")
	flag.Parse()

	//brain, err := brain.NewBrain("gpt-3.5-turbo")
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	return
	//}

	fmt.Println(clippy)
	fmt.Printf("\nHello there.\n")
	fmt.Printf("Waiting for terminal connection...\n")

	// Set up channel to listen for terminal data
	termCh := make(chan string)
	defer close(termCh)
	go func() {
		for str := range termCh {
			fmt.Println("output:", str)
		}
	}()

	// Create the named pipe
	pipePath := "/tmp/clippy-pipe"
	os.Remove(pipePath) // Remove the named pipe if it already exists

	err := syscall.Mkfifo(pipePath, 0666)
	if err != nil {
		log.Fatal("Error creating pipe: ", err)
	}

	// Open the named pipe for reading
	pipe, err := os.OpenFile(pipePath, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatal("Error opening pipe: ", err)
	}

	// Read from the named pipe until null character is encountered
	data := make([]byte, 1)
	buf := make([]byte, 0)
	for {
		n, err := pipe.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("script complete.")
				break
			}
			log.Fatal("Error reading from pipe: ", err)
		}

		if data[0] == 0 {
			// Null character encountered, send the current buffer to the goroutine
			//fmt.Printf("ouputting to termch\n")
			str := string(buf)
			go func() {
				termCh <- str
			}()

			// Reset the buffer
			buf = make([]byte, 0)
		} else {
			//fmt.Printf("appending %s\n", data)
			// Append the byte to the buffer
			buf = append(buf, data[:n]...)
		}
	}

	// NOTE: to write to that, do something like
	// script -F /tmp/clippy-pipe
	// ls ; echo -n "\0"
	// exit

	// Print the data read from the pipe
	//log.Println("Received data:", string(buf))

	// Close the named pipe
	err = pipe.Close()
	if err != nil {
		log.Fatal("Error closing pipe: ", err)
	}

	return

	//intepretedGoal, err := brain.InterpretGoal(*goal)
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	return
	//}
	//fmt.Printf("%s\n", intepretedGoal)

}

func cleanupOnExit(pipePath string) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Remove(pipePath)
		os.Exit(0)
	}()
}
