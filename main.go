package main

import (
	"github.com/sgjp/LindaExperimentClient/client"
	"log"
	"strings"
	"strconv"
	"os"
	"encoding/csv"
	"time"
)

var mode int
var primeNumsQty int
var taskDurationFile = "/Users/jsanchez/workspace/src/github.com/sgjp/LindaExperimentClient/TaskDuration.csv"

func main() {
	//mode sets the device to Worker: 0 or Manager: 1
	mode = 0;

	//sets the number of prime numbers needed, this only applies for when the device is set to 1 Manager
	primeNumsQty = 5000

	start()

}

func start(){
	if mode ==1{
		log.Printf("Starting as manager, generating W tuples...")
		//Generate requests for primer numbers
		for i:=1;i<=primeNumsQty;i++{
			//log.Printf("Writing tuple: %v","W,"+strconv.Itoa(i))
			client.OutTuple("W,"+strconv.Itoa(i))
		}
		/*log.Printf("Searching for results...")

		//Check for the same ammount of results

				for true{
					tuple := client.InTuple("R")
					if tuple != "0"{
						//log.Printf("Adding to i, Proccessed Tuple found: %v",tuple)
						i++
					}else{
					}
					//log.Printf("i: %v",i)
					if i>primeNumsQty{
						break
					}
					time.Sleep(100*time.Millisecond)


				}
				elapsed := time.Since(t)
				durationExtra := int64(100*delays)
				log.Printf("Prime numbers calculated!, it took %v Ms and %v S",int64(elapsed/time.Millisecond)-durationExtra,(int64(elapsed/time.Millisecond))/1000)
				saveTaskDuration((int64(elapsed/time.Millisecond)-durationExtra),primeNumsQty)
				*/
		log.Printf("Done producing !")
	}else{
		//Get requests, process them and return the result
		var tuplesToSend []string
		i := 1
		log.Printf("Starting as worker, looking for W tuples...")
		for  {
			time.Sleep(10*time.Millisecond)
			tuple := client.InTuple("W")
			if tuple != "0"{
				log.Printf("Found!: %v",tuple)
				i++
				splittedTuple := strings.Split(tuple,",")
				qty, err := strconv.Atoi(splittedTuple[1])
				if err!= nil{
					log.Printf("Error parsing tuple %v",err)
				}
				result := calcPrimeNumber(qty)

				tuplesToSend = append(tuplesToSend,"R,"+splittedTuple[1]+","+strconv.Itoa(result))
			}
			if i>primeNumsQty{
				break
			}
		}
		for b:=0;b<=len(tuplesToSend)-1;b++{
			//log.Printf("Writing tuple: %v", tuplesToSend[b])
			client.OutTuple(tuplesToSend[b])
		}
		log.Printf("Finished working !")

	}
}


func calcPrimeNumber(qty int) int {
	var num int
	for i := 2; i < qty; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			} else if i == j+1 {
				num = i
			}
		}
	}
	return num
}

func saveTaskDuration(elapsed int64, qty int){
	record := []string{
		strconv.Itoa(qty), strconv.FormatInt(elapsed,10)}

	file, er := os.OpenFile(taskDurationFile, os.O_RDWR|os.O_APPEND, 0666)

	if er != nil {
		log.Fatal(er)
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	err := writer.Write(record)


	if err != nil {
		log.Fatal(er)
	}

	defer writer.Flush()
}
