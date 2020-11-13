package main
import(
	"fmt"
	"log"
	"net/http"
	
)
func homepage(w http.ResponseWriter,r *http.Request){
fmt.Fprintf(w,"rest api perform the arithmetic operations")
}
func handleRequests(){
http.HandleFunc("/",homepage)
var num1 int=8
var num2 int=4
add:=num1+num2
fmt.Println("%d\n",add)
sub:=num1-num2
fmt.Println("%d",sub)
mul:=num1*num2
fmt.Println("%d",mul)
div:=num1/num2
fmt.Println("%d",div)
log.Fatal(http.ListenAndServe(":8080",nil))
}
func main(){
handleRequests()

}

