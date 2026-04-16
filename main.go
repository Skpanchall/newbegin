package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Skpanchall/newbegin/handler"
	"github.com/Skpanchall/newbegin/utils"
)

type middlewareWriter struct {
	http.ResponseWriter
	code int
}

func (w *middlewareWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func middleWare(h http.HandlerFunc, middlewares ...MiddlewareFunc) http.HandlerFunc {

	for i := 0; i < len(middlewares); i++ {
		h = middlewares[i](h)
	}
	return h

}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")

		if token != "123" {
			utils.SendErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		mw := &middlewareWriter{ResponseWriter: w}
		next(mw, r)
		duration := time.Since(start)
		fmt.Println("[", r.Method, "]", r.URL.Path, "time : ", duration, "status :", mw.code)
	}
}

func main() {
	http.HandleFunc("/", middleWare(handler.WelcomeAPI, authMiddleware, loggingMiddleware))
	http.HandleFunc("/users", middleWare(handler.HandleUsers, authMiddleware, loggingMiddleware)) // order is here is first called loggingmiddleware and then authmiddleware
	http.HandleFunc("/user", middleWare(handler.HandleUser, authMiddleware, loggingMiddleware))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Err while listen server :", err)
	}
	return

	var choice int

	fmt.Println("1. Bio Maker")
	fmt.Println("2. Age Checker")
	fmt.Println("3. Work Status Checker")
	fmt.Println("4. Give Numbers and Print Even and Odd numbers")
	fmt.Println("5. functions")
	fmt.Println("6. arrya and slices")
	fmt.Println("7. sturcts")
	fmt.Println("8. Profile update")
	fmt.Println("9. API User Processor")
	fmt.Println("Enter choice:")

	fmt.Scanln(&choice)

	if choice == 1 {
		BioMaker()
	} else if choice == 2 {
		Agechecker()
	} else if choice == 3 {
		WorkStatusChecker()
	} else if choice == 4 {
		EvenNumberPrint()
	} else if choice == 5 {
		FunctionBased()
	} else if choice == 6 {
		ArrayAndSlices()
	} else if choice == 7 {
		GetProductTotalValues()
	} else if choice == 8 {
		GetProfile()
	} else if choice == 9 {
		ApiUserProcess()
	} else if choice == 10 {
		UserCli()
	} else {
		fmt.Println("Invalid choice")
	}

}
