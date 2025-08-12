# Gonsole Example

シンプルなGo言語のサンプルプロジェクトです。interface、struct、var、短縮変数宣言、method、const、functionなどの要素を含んでいます。

## プロジェクト構造

```
gonsole-example/
├── go.mod              # Go モジュールファイル
├── README.md           # このファイル
├── animal/
│   ├── animal.go       # 動物インターフェースと実装（犬、猫）
│   └── utils/
│       └── utils.go    # 動物関連のユーティリティ関数
├── plant/
│   ├── plant.go        # 植物インターフェースと実装（花、木）
│   └── utils/
│       └── utils.go    # 植物関連のユーティリティ関数
└── vehicle/
    ├── vehicle.go      # 乗り物インターフェースと実装（車、バイク）
    └── utils/
        └── utils.go    # 乗り物関連のユーティリティ関数
```

## 特徴

### 含まれている Go 言語の要素

1. **Interface**: `Animal`, `Plant`, `Vehicle` - 各カテゴリーの共通動作を定義
2. **Struct**: 
   - 基底構造体: `BaseAnimal`, `BasePlant`, `BaseVehicle`
   - 具象構造体: `Dog`, `Cat`, `Flower`, `Tree`, `Car`, `Bike`
3. **Var**: パッケージレベルの変数宣言
4. **短縮変数宣言**: `:=` を使用した変数宣言（各メソッド内で使用）
5. **Method**: 構造体のメソッド、メソッドチェーン対応
6. **Const**: 定数宣言（`DefaultBreed`, `DefaultColor`, `DefaultEngine` など）
7. **Function**: パッケージレベルの関数とユーティリティ関数

### パッケージ構成

#### animal パッケージ
- 動物の基本的な動作（鳴く、描写、餌やり、運動）を定義
- 犬（Dog）と猫（Cat）の実装
- 年齢換算や名前フォーマットのユーティリティ

#### plant パッケージ
- 植物の基本的な動作（成長、開花、描写、水やり、肥料）を定義
- 花（Flower）と木（Tree）の実装
- 成長段階判定や名前フォーマットのユーティリティ

#### vehicle パッケージ
- 乗り物の基本的な動作（始動、停止、描写、給油、加速）を定義
- 車（Car）とバイク（Bike）の実装
- 年数計算や燃費計算のユーティリティ

### メソッドチェーン

全てのオブジェクトでメソッドチェーンが可能です：

```go
// 動物
dog := animal.NewDog("ポチ", 3)
dog.Feed().Exercise()

// 植物
flower := plant.NewFlower("バラ", 2)
flower.Water().Fertilize()

// 乗り物
car := vehicle.NewCar("トヨタ", "プリウス", 2020)
car.Refuel().Accelerate()
```

## 学習ポイント

このプロジェクトは、Go言語の以下の概念を実践的に学習できます：

1. **インターフェース設計**: 共通の動作を抽象化
2. **構造体の組み込み**: BaseStructを使った継承的な設計
3. **メソッドレシーバー**: ポインタレシーバーを使ったメソッド実装
4. **メソッドチェーン**: 流暢なインターフェース（Fluent Interface）の実装
5. **パッケージ分割**: 機能別のパッケージ構成
6. **ユーティリティ関数**: 共通処理の分離
7. **定数とコンストラクタ**: オブジェクトの初期化パターン


