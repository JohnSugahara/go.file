#file go, calcular baskhara

// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import ()
    "fmt"
    "math"
)

func main() {
  
  var a, b, c float64
  
    fmt.Print("Digite o valor de a: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de b: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de c: ")
	fmt.Scan(&c)
	
    delta = b*b (- (4*a*c))
    
    if delta < 0
    {
        fmt.Print = "por delta ser menor que zero, nao possui resultados"
    }
    else{
        
        resultado1 = (-b + math.Sqrt(delta)) / (2*a)
        resultado2 = (-b - math.Sqrt(delta)) / (2*a)
        
        fmt.Printf("As soluções são: x1 = %.2f e x2 = %.2f\n", resultado1, resultado2)
    }
}



