package utils

import (
	"testing"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BenchmarkBinarySearchAsnIPv4(b *testing.B) {
	for i:=0;i<b.N;i++ {
		x := i%254
		y := i%8
		m := i%255
		n := i%255
		res, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/check/%d.%d.%d.%d",y,m,n,x))
		if err != nil {
			b.Fatal(err)
		}
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			b.Fatal(err)
		}
		b.Log(fmt.Sprintf("%s", robots))
	}
}
