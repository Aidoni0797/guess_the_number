func foo() {
	// паникуем
	panic("unexpected!")
}

func main() {
	// выполняется при завершении main
	defer func() {
		// вызываем recover и сравниваем результат с nil
		if r := recover(); r != nil {
			fmt.Println(r) // выведет "unexpected"
		}
	}()
	foo() // внутри foo срабатывает паника
	fmt.Println("Вы не увидите это сообщение, так как в foo возникла паника")
}