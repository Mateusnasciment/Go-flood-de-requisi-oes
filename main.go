package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "https://archlinux-drab.vercel.app/"
	numRequests := 672

	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Erro ao acessar %s: %s\n", url, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Página %s acessada com sucesso! Status: %s\n", url, resp.Status)
		}()
	}

	wg.Wait()
	fmt.Println("Todas as solicitações concluídas.")
}
