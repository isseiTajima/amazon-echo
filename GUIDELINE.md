## Go クリーンアーキテクチャ + DDD ガイドライン
### 基本原則
•	依存関係の方向: 外側から内側へ（ドメイン層に依存する）
•	関心の分離: 各レイヤーは単一の責務を持つ
•	テスタビリティ: 依存をインターフェースで抽象化し、テスト容易性を確保
### DDD実践のポイント
1.	ユビキタス言語: コードとビジネス用語を一致させる
2.	集約: 関連するエンティティと値オブジェクトをグループ化
3.	境界づけられたコンテキスト: 異なるドメインを明確に分離
4.	値オブジェクト: 不変で同一性を持たないオブジェクトを活用
### コーディング規約
•	インターフェースはドメイン層で定義し、実装は外側のレイヤーで行う
•	ドメインロジックはドメイン層に閉じ込める
•	依存性注入を活用し、具象型への依存を避ける
•	エラーハンドリングは各レイヤーで適切に行う
### プロジェクト構造
```bash
├── main.go                  # エントリーポイント
├── domain/                  # エンティティ、値オブジェクト、ドメインサービス
├── usecase/                 # ユースケース実装
├── adapter/                 # コントローラー、リポジトリ実装
├── infrastructure/          # DB、サーバー、DI設定
└── test/                    # テストコード

```

### 各レイヤーの役割
#### ドメイン層
```go
// ドメインエンティティ
type User struct {
    ID    uint
    Name  string
    Email string
}

// リポジトリインターフェース
type UserRepository interface {
    FindByID(id uint) (*User, error)
    Save(user *User) error
}
```

#### ユースケース層
```go
type UserUseCase struct {
    repo UserRepository
}

func (uc *UserUseCase) GetUser(id uint) (*User, error) {
    return uc.repo.FindByID(id)
}
```

#### アダプター層
```go
// コントローラー
func (c *UserController) GetUser(ctx echo.Context) error {
    id, _ := strconv.Atoi(ctx.Param("id"))
    user, err := c.useCase.GetUser(uint(id))
    if err != nil {
        return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
    }
    return ctx.JSON(http.StatusOK, user)
}
```

#### インフラストラクチャ層
```go

// リポジトリ実装
func (r *UserRepositoryImpl) FindByID(id uint) (*User, error) {
    var user User
    result := r.db.First(&user, id)
    return &user, result.Error
}
// DI設定
func BuildContainer() *fx.App {
    return fx.New(
        fx.Provide(
            database.NewDB,
            repository.NewUserRepository,
            usecase.NewUserUseCase,
            controller.NewUserController,
            server.NewEchoServer,
        ),
        fx.Invoke(func(e *echo.Echo) {
            e.Start(":8080")
        }),
    )
}


```
