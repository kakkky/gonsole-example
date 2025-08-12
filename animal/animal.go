package animal

import "fmt"

// 動物インターフェース
type Animal interface {
	// 動物の鳴き声を返す
	Sound() string
	// 動物の説明を返す
	Describe() string
	// メソッドチェーン用
	Feed() Animal
	// メソッドチェーン用
	Exercise() Animal
	// 動物の状態を返す
	GetStatus() string
	// 動物の名前を変更する
	ChangeName(newName string) Animal
}

// 基本動物構造体
type BaseAnimal struct {
	Name  string
	Age   int
	Fed   bool
	Tired bool
}

// 犬構造体
type Dog struct {
	BaseAnimal
	Breed string // 犬種
}

// 猫構造体
type Cat struct {
	BaseAnimal
	Color string // 毛色
}

// 定数
const (
	DefaultBreed = "雑種"
	DefaultColor = "茶色"
)

// 犬のコンストラクタ
func NewDog(name string, age int) *Dog {
	return &Dog{
		BaseAnimal: BaseAnimal{
			Name:  name,
			Age:   age,
			Fed:   false,
			Tired: false,
		},
		Breed: DefaultBreed,
	}
}

// 猫のコンストラクタ
func NewCat(name string, age int) *Cat {
	return &Cat{
		BaseAnimal: BaseAnimal{
			Name:  name,
			Age:   age,
			Fed:   false,
			Tired: false,
		},
		Color: DefaultColor,
	}
}

// 犬のインターフェース実装
func (d *Dog) Sound() string {
	return "ワンワン"
}

func (d *Dog) Describe() string {
	return fmt.Sprintf("犬の%s（%d歳、%s）", d.Name, d.Age, d.Breed)
}

func (d *Dog) Feed() Animal {
	d.Fed = true
	return d // メソッドチェーン用
}

func (d *Dog) Exercise() Animal {
	d.Tired = true
	return d // メソッドチェーン用
}

func (d *Dog) GetStatus() string {
	// 短縮変数宣言
	status := "元気"
	if d.Fed && d.Tired {
		status = "満足"
	} else if d.Fed {
		status = "満腹"
	} else if d.Tired {
		status = "疲れ気味"
	}
	return status
}

// 犬専用のメソッド
func (d *Dog) SetBreed(breed string) {
	d.Breed = breed
}

// 名前を変えるメソッド
func (d *Dog) ChangeName(newName string) Animal {
	d.Name = newName
	return d
}

func NewAnimal(name string, age int, animalType string) Animal {
	switch animalType {
	case "dog":
		return NewDog(name, age)
	case "cat":
		return NewCat(name, age)
	default:
		return nil // 未知の動物タイプ
	}
}

// 猫のインターフェース実装
func (c *Cat) Sound() string {
	return "ニャーニャー"
}

func (c *Cat) Describe() string {
	return fmt.Sprintf("猫の%s（%d歳、%s）", c.Name, c.Age, c.Color)
}

func (c *Cat) Feed() Animal {
	c.Fed = true
	return c // メソッドチェーン用
}

func (c *Cat) Exercise() Animal {
	c.Tired = true
	return c // メソッドチェーン用
}

func (c *Cat) GetStatus() string {
	// 短縮変数宣言
	status := "のんびり"
	if c.Fed && c.Tired {
		status = "ぐっすり"
	} else if c.Fed {
		status = "満足"
	} else if c.Tired {
		status = "だらけ気味"
	}
	return status
}

// 猫専用のメソッド
func (c *Cat) SetColor(color string) {
	c.Color = color
}

// 名前を変えるメソッド
func (c *Cat) ChangeName(newName string) Animal {
	c.Name = newName
	return c
}
