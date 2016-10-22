package main   // <-- this is a package name, similar to a Java program.  The main package is the starting point for any program. Packages do not necessarily correspond to filesystem path like they do in Java

//imports can either be done individually as you would in Java or as a list in a single statement, called a "factored import", like so
//  while there is no functional improvement to doing it this way,  this is the best practice for convenience and style reasons
import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	_ "log" // <-- Import for side effect, log isn't a great example, but some packages perform functions on initialization, the underbar(blank identifier) allows the program to compile even though an import is not specifically used
)

//if you are in development and you have an import you aren't ready to use yet but don't want to get rid of, you can silence the unused variable with the blank identifier as well
//var _ = fmt.Printf

func main() {

}
		