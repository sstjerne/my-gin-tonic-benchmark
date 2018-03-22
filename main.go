package main

import (
	"github.com/gin-gonic/gin"
	"crypto/sha512"
	"encoding/base64"
	"crypto/sha1"
	"io/ioutil"
	"regexp"
	"fmt"
	"os"
)


var DB = make(map[string]string)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}


func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// throughput test
	r.GET("/throughput", func(c *gin.Context) {
		c.JSON(200, gin.H{"throughput": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In in ipsum a velit faucibus tempor vel nec odio. Interdum et malesuada fames ac ante ipsum primis in faucibus. Curabitur quis orci eget purus tempus aliquet eu eu risus. Ut velit elit, viverra et ex vel, scelerisque rhoncus odio. Donec vitae diam pellentesque, commodo velit et, lacinia leo. In vel pharetra purus, sed eleifend nunc. Maecenas porta rhoncus consectetur. In hac habitasse platea dictumst. Duis sed erat nibh. Morbi imperdiet lorem purus, vitae facilisis enim maximus et. Phasellus ullamcorper sapien eget neque eleifend malesuada."})
	})

	// cpu test
	r.GET("/cpu", func(c *gin.Context) {
		hasher := []byte("Sparkers doing some benchmarking")
		sha512 := sha512.New()
		for i := 0; i < 256; i++ {
			sha512.Write(hasher)
		}

		c.JSON(200, gin.H{"Hashed": base64.URLEncoding.EncodeToString(sha512.Sum(nil)) })
	})


	// ram test
	r.GET("/ram", func(c *gin.Context) {
		byteArray, err := ioutil.ReadFile("support/ram_test.txt") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		str := fmt.Sprintf("%s", byteArray)
		pattern := regexp.MustCompile("[aeiouAEIOU]")
		matches := pattern.FindAllStringIndex(str, -1)
		c.JSON(200, gin.H{"n_vowels": len(matches) })
	})

	// disk test
	r.GET("/disk", func(c *gin.Context) {
		byteArray, err := ioutil.ReadFile("support/disk_test.csv") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}

		f, err := os.Create("/tmp/disk_test.csv.tmp")
		check(err)	

		defer f.Close()

		n2, err := f.Write(byteArray)
		check(err)
		count := fmt.Sprintf("%d", n2)

		f, err := os.Remove("/tmp/disk_test.csv.tmp")
		check(err)

		c.JSON(200, gin.H{"bytes": count })
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:80
	r.Run(":80")
}
