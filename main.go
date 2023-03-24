package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	scanner := bufio.NewScanner(os.stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.scan() {
		checkDomain(scanner.Text())

		if err := scanner.err(); err != nil {
			log.Fatal("Error: could not readfrom input: %v", err)
		}
	}
}

func checkDomain(domain string) {
   var hasMX, hasSPF, hasDMARC, bool
   var hasSPF, dmarcRecord string

   mxRecords, err := net.LookupMX(domain)

   if err != nil{
	 log.Printf("Error: %v", err)
   }
   if len(mxRecords) > 0 {
	hasMX = true
   }
   txtRecords, err := net.Lookup.TXT(domain)
   if err != nil {
	log.Printf("Error: %v", err)
   }
} 
