package main

func somar(a int, b int) int {
	return a + b
}

func calculosMatematicos(a int, b int) (int, int, int) {
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	return soma, subtracao, multiplicacao
}

func main() {
	soma := somar(10, 20)
	println(soma)

	var f = func(txt string) string {
		println(txt)
		return txt
	}

	resultado := f("Texto de retorno")
	println(resultado)

	rSoma, rSub, _ := calculosMatematicos(10, 20)
	println(rSoma, rSub)
}
