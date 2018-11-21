package main
import(
	"fmt"
	"regexp"
)
const(
	HEADER_1=`^#\w*#$`
)

func main(){
	matcher,_:=regexp.Compile(HEADER_1)
	fmt.Println(matcher.FindAllString("#hello world#",-1))
}