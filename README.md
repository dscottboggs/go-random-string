# go-random-string
Generate cryptographically safe random alphanumeric strings, or strings constructed from random words.

## usage
package main                                                                    
                                                                                
import (                                                                        
    "fmt"                                                                       
    "github.com/dscottboggs/go-random-string"
)                                                                                  
func main() {                                                                   
    fmt.Println(random.AlphanumericString(20)) //=> [twenty random alphanumeric characters]
    fmt.Println(random.Words(3, "--")) // [three random words separated by --]  
}
