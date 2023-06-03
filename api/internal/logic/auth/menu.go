package logic

import (
	"api/internal/service"
	"context"
	"fmt"
)

type sMenu struct{}

func NewMenu() *sMenu {
	return &sMenu{}
}

func init() {
	fmt.Println("Menu")
	service.RegisterMenu(NewMenu())
}

func (logic *sMenu) List(ctx context.Context) {
	fmt.Println("Menu1")

}
