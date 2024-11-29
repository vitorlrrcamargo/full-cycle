package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "1!2@3#")
	bookHotel(ctx, "Hotel Internacional")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Printf("Hotel Name: %v - Token: %v", name, token)
}
