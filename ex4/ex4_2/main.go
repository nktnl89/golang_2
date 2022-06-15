package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("чагыр")

	input := make(chan string, 100)

	ctx, cancelFunc := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "is_running", true)

	for ctx.Value("is_running") == true {
		go func(ctx context.Context) {
			fmt.Print("Введите что-нибудь, получите ответ, введите 'SIGTERM', программа завершится: ")
			var u string
			fmt.Scanf("%s\n", &u)
			input <- u
		}(ctx)

		select {
		case val := <-input:
			if val == "SIGTERM" {
				fmt.Println("Конец.")
				cancelFunc()

				ctx = context.WithValue(ctx, "is_running", false)
			} else {
				fmt.Println("чагыр-чагыр")
			}
		}
	}
}
