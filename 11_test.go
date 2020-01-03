package design

import (
	"fmt"
	"testing"
)

func Test_11_1(t *testing.T) {
	client := new(Facade)
	client.WeaponSystem = new(WeaponSystem)
	client.UserSystem = new(UserSystem)
	client.Fire()
	client.Bullet()
	fmt.Println("-----------")
	client.Shooting()
}
