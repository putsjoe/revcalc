// revcalc

package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
    fmt.Print("")
    var input string
    //var stack int
    
    for input != "exit" && input != "e" {
        
        reader := bufio.NewReader(os.Stdin)
        text,_ := reader.ReadString('\n')
        //fmt.Println(text)
        
        split := strings.Fields(text)
        
        //fmt.Println(split, len(split))
        
        // Take the first value, second and the operator from third entry.
        
        var first,_ = strconv.Atoi(split[0])
        var sec,_ = strconv.Atoi(split[1])
        
        switch split[2]{
            case "*":
            fmt.Println(first * sec)
            case "/": 
            fmt.Println(first / sec)
            case "-": 
            fmt.Println(first - sec)
            case "+": 
            fmt.Println(first + sec)
            
            default: fmt.Println("Unknown")
            
        }
        
       
    }
    


   /*
   fmt.Scanln(&input) << Think the issue at the moment is this line.
   The doc says that it takes each string seperated by space is taken as an argument.
   */
}


