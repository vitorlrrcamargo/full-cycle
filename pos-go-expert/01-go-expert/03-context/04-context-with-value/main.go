package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "A1B2C3D4D5")
	bookHotel(ctx, "International Hotel")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println("Hotel:", name, "| Token:", token)
}
