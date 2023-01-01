package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var psCertificate sync.Mutex
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func ()  {
		defer wg.Done()

		psCertificate.Lock()
		defer psCertificate.Unlock()

		time.Sleep(5 * time.Second)
		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()

	go func ()  {
		defer wg.Done()

		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		time.Sleep(5 * time.Second)
		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("成功完成")
}