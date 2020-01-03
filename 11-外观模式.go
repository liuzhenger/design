package design

import "fmt"

type WeaponSystem struct {
}

func (w *WeaponSystem) Bullet() {
	fmt.Println("上子弹")
}

func (w *WeaponSystem) Fire() {
	fmt.Println("开火")
}

type UserSystem struct {
}

func (u *UserSystem) Score() {
	fmt.Println("得分")
}

type Facade struct {
	*WeaponSystem
	*UserSystem
}

func (f *Facade) Shooting() {
	f.Bullet()
	f.Fire()
	f.Score()
}
