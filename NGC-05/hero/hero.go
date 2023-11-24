package hero

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

/*
Buatlah sebuah struct `Hero` dengan properti sebagai berikut

```
Name int
BaseAttack int
Defence int
CriticalDamage int
HealthPoint int
```

Dan buatlah sebuah struct `Weapon` dengan properti `Attack int` dan tambahkan properti `Weapon` ke dalam struct `Hero`.

# Lalu buatlah method untuk menghitung total damage `CountDamage` yang diberikan jika `Hero` tersebut menyerang, dengan ketentuan

```
total damage adalah base attack ditambahkan dengan attack dari weapon
milik hero.dan ditambahkan dengan critical damage, namun critical damage
hanya akan ditambahkan dengan peluang 50:50, gunakan `Math.Rand
````
*/

/*
	Dan buatlah method `isAttackedBy` yang menerima input `Hero`, dan didalam
	method tersbut akan melakukan kalkulasi total damage yang diterima dan berapa pengurangan health point nya. dengan ketentuan
	```
	total damage yang diterima adalah hasil pengurangan dari total damage dari
	Hero yg diinput (dari CountDamage) dikurangi Defencesetelah diketahui total damage yang diterima, maka health point akan dikurangi sebesar total damage yang diterima. Jika total damage serangan lebih kecil daripada Defence maka healhpoint tidak berkurang maupun bertambah
	```
	Buatlah sebuah function dengan nama `Battle` yang menerima input 2
	argument `Hero`dimana argument pertama adalah `Hero` yang menyerang, dan
	input ke 2 adalah `Hero` yang diserang.Lakukanlah kalkulasi HealthPoint
	dari hero yang mengalami serangan. Manfaatkan method `isAttackedBy`
*/

type Weapon struct {
	Attack int
}

type Hero struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         Weapon
}

func (h Hero) GetInfo() string {
	var format string
	rv := reflect.ValueOf(&h)
	rt := rv.Elem().Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := reflect.Indirect(rv).FieldByName(field.Name)
		format += fmt.Sprintf("%v: %v\n", field.Name, value)
	}
	return format
}

func (h Hero) CountDamage() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	iscritical := random.Intn(11) > 5

	if iscritical {
		damagecrit := h.Weapon.Attack + h.CriticalDamage + h.BaseAttack
		fmt.Println(h.Name, "attack critical hit...", (damagecrit))
		return damagecrit
	}

	fmt.Println(h.Name, "attack...", h.Weapon.Attack)
	return h.Weapon.Attack + h.BaseAttack
}

func (h *Hero) IsAttackedBy(attacker Hero) {
	damage := attacker.CountDamage() - h.Defence
	if damage > 0 {
		h.HealthPoint -= damage
		fmt.Printf("%s receive critical damage %v\n", h.Name, damage)
		return
	}

	fmt.Printf("%s has good defence (%d) | damage receive (%d)\n", h.Name, h.Defence, damage)
}

func Challenge_3_4() {
	batman := Hero{
		Name:           "Batman",
		BaseAttack:     20,
		Defence:        50,
		CriticalDamage: 80,
		HealthPoint:    100,
		Weapon: Weapon{
			Attack: 20,
		},
	}

	superman := Hero{
		Name:           "Superman",
		BaseAttack:     20,
		Defence:        50,
		CriticalDamage: 80,
		HealthPoint:    100,
		Weapon: Weapon{
			Attack: 20,
		},
	}

	Battle(batman, superman)
}

func Battle(hero_a, hero_b Hero) {
	fmt.Printf("Hero Battle: \n\n")
	fmt.Println(hero_a.GetInfo())
	fmt.Printf("VS\n\n")
	fmt.Println(hero_b.GetInfo())

	fmt.Println(hero_a.Name, "Is attacked by", hero_b.Name)
	hero_a.IsAttackedBy(hero_b)
	fmt.Printf("\n%v\n", hero_a.GetInfo())

}
