package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"
)

func main() {

	// available commands:
	// adb shell am startservice -a com.blogspot.newapphorizons.fakegps.START -e "latitude" 1.11 -e "longitude" 0.21
	// adb shell am startservice -a com.blogspot.newapphorizons.fakegps.UPDATE -e "latitude" 1.11 -e "longitude" 0.21
	// adb shell am startservice -a com.blogspot.newapphorizons.fakegps.STOP

	track, err := os.Open("./track.csv")
	if err != nil {
		panic(err)
	}
	defer track.Close()

	scanner := bufio.NewScanner(track)
	// read heading
	if !scanner.Scan() {
		panic("failed to read heading")
	}
	// read 1st coordinates
	lat, long, alt, bearing, err := readData(scanner)
	if err != nil {
		panic(err)
	}

	// start GPS faking
	// NOTE: altitude & bearing doesn't work
	command := fmt.Sprintf("START -e \"latitude\" %v -e \"longitude\" %v -e \"altitude\" %v -e \"bearing\" %v", lat, long, alt, bearing)
	run(command)

	// stop faking when exiting
	defer stop()
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		stop()
		os.Exit(0)
	}()

	// update position
	for {
		lat, long, alt, bearing, err = readData(scanner)
		if err != nil {
			break
		}
		//fmt.Printf("update to lat %v, long %v, alt %v, bearing %v\n", lat, long, alt, bearing)
		command = fmt.Sprintf("UPDATE -e \"latitude\" %v -e \"longitude\" %v -e \"alt\" %v -e \"bear\" %v", lat, long, alt, bearing)
		run(command)
		time.Sleep(300 * time.Millisecond)
	}
	if err != nil {
		panic(err)
	}
}

func readData(s *bufio.Scanner) (lat, long, alt, bearing string, err error) {
	if !s.Scan() {
		err = s.Err()
		if err == nil {
			err = fmt.Errorf("EOF")
		}
		return
	}
	data := strings.Split(s.Text(), ",")
	lat, long, alt, bearing = data[0], data[1], data[2], data[4]
	return
}

func stop() {
	command := "STOP"
	run(command)
}

func run(command string) {
	fixedArgs := "shell am startservice -a com.blogspot.newapphorizons.fakegps."
	fullArgs := fixedArgs + command
	args := strings.Split(fullArgs, " ")
	cmd := exec.Command("adb", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("failed to execute command %q\nout: %q\n", args, strings.TrimSpace(string(out)))
		panic(err)
	} else {
		fmt.Printf("out: %q\n", strings.TrimSpace(string(out)))
	}
}
