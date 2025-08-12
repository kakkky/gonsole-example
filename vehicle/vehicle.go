package vehicle

import "fmt"

// 乗り物インターフェース
type Vehicle interface {
	Start() string
	Stop() string
	Describe() string
	Refuel() Vehicle     // メソッドチェーン用
	Accelerate() Vehicle // メソッドチェーン用
	GetStatus() string
	ChangeName(newBrand, newModel string) Vehicle // 乗り物の名前（ブランド、モデル）を変更する
}

// 基本乗り物構造体
type BaseVehicle struct {
	Brand    string
	Model    string
	Year     int
	Running  bool
	FuelFull bool
}

// 車構造体
type Car struct {
	BaseVehicle
	Engine string // エンジンタイプ
}

// バイク構造体
type Bike struct {
	BaseVehicle
	Type string // バイクタイプ
}

// 定数
const (
	DefaultEngine = "ガソリン"
	DefaultType   = "スポーツ"
)

// 車のコンストラクタ
func NewCar(brand, model string, year int) *Car {
	return &Car{
		BaseVehicle: BaseVehicle{
			Brand:    brand,
			Model:    model,
			Year:     year,
			Running:  false,
			FuelFull: false,
		},
		Engine: DefaultEngine,
	}
}

// バイクのコンストラクタ
func NewBike(brand, model string, year int) *Bike {
	return &Bike{
		BaseVehicle: BaseVehicle{
			Brand:    brand,
			Model:    model,
			Year:     year,
			Running:  false,
			FuelFull: false,
		},
		Type: DefaultType,
	}
}

// 車のインターフェース実装
func (c *Car) Start() string {
	c.Running = true
	return "ブルルル... エンジン始動"
}

func (c *Car) Stop() string {
	c.Running = false
	return "エンジン停止"
}

func (c *Car) Describe() string {
	return fmt.Sprintf("%s %s (%d年式、%sエンジン)", c.Brand, c.Model, c.Year, c.Engine)
}

func (c *Car) Refuel() Vehicle {
	c.FuelFull = true
	return c // メソッドチェーン用
}

func (c *Car) Accelerate() Vehicle {
	c.Running = true
	return c // メソッドチェーン用
}

func (c *Car) GetStatus() string {
	// 短縮変数宣言
	status := "停車中"
	if c.Running && c.FuelFull {
		status = "快調走行中"
	} else if c.Running {
		status = "走行中（燃料不足）"
	} else if c.FuelFull {
		status = "満タン停車中"
	}
	return status
}

// 車専用のメソッド
func (c *Car) SetEngine(engine string) {
	c.Engine = engine
}

// 名前を変えるメソッド（ブランドとモデル）
func (c *Car) ChangeName(newBrand, newModel string) Vehicle {
	c.Brand = newBrand
	c.Model = newModel
	return c
}

// バイクのインターフェース実装
func (b *Bike) Start() string {
	b.Running = true
	return "ブロロロ... バイク始動"
}

func (b *Bike) Stop() string {
	b.Running = false
	return "バイク停止"
}

func (b *Bike) Describe() string {
	return fmt.Sprintf("%s %s (%d年式、%sタイプ)", b.Brand, b.Model, b.Year, b.Type)
}

func (b *Bike) Refuel() Vehicle {
	b.FuelFull = true
	return b // メソッドチェーン用
}

func (b *Bike) Accelerate() Vehicle {
	b.Running = true
	return b // メソッドチェーン用
}

func (b *Bike) GetStatus() string {
	// 短縮変数宣言
	status := "停車中"
	if b.Running && b.FuelFull {
		status = "爽快走行中"
	} else if b.Running {
		status = "走行中（燃料不足）"
	} else if b.FuelFull {
		status = "満タン停車中"
	}
	return status
}

// バイク専用のメソッド
func (b *Bike) SetType(bikeType string) {
	b.Type = bikeType
}

// 名前を変えるメソッド（ブランドとモデル）
func (b *Bike) ChangeName(newBrand, newModel string) Vehicle {
	b.Brand = newBrand
	b.Model = newModel
	return b
}
