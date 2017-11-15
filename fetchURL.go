package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)
func main(){

	for _,url := range os.Args[1:] {
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr,"fetch err %v %v", err, url)
			os.Exit(1)
		}
		b,err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"read error %v, %v", err, url)
			os.Exit(1)
		}
		fmt.Printf("%s",b)

	}

}
