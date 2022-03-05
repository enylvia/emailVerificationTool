package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Sprintf("domain, hasMX, hasSPF,sprRecord,hasDMARC,dmarcRecord \n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read data from input: %v\n",err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords,err := net.LookupMX(domain)
	if err != nil {
		log.Printf("%s, hasNoMxRecord \n",domain)
		return
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecord,err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("%s, hasNoTxtRecord \n",domain)
		return
	}

	for _, val := range txtRecord {
		if strings.HasPrefix(val, "v=spf1") {
			hasSPF = true
			spfRecord = val
			break
		}
	}
	dmarcRecords,err := net.LookupTXT("_dmarc."+domain)
	if err != nil {
		log.Printf("%s, hasNoDMARCRecord \n",domain)
		return
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v",domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord)
}
