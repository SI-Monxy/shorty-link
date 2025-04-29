# 🔗 shorty-link

URL短縮サービスをGo（Gin）で構築したポートフォリオ用プロジェクトです。
クリーンアーキテクチャ・Docker・GORM・MySQL を採用し、今後TerraformやCI/CDも導入予定です。


## ✨ 主な技術スタック

- Go 1.24
- Gin（HTTPサーバー）
- GORM（ORM）
- MySQL（Docker）
- Docker / Docker Compose
- クリーンアーキテクチャ
- REST API


## 🚀 現在の実装機能

### ✅ POST `/api/v1/shorten`

**URLを短縮するAPI**

#### 📥 リクエスト

```json
{
  "original_url": "https://example.com"
}
```

#### 📤 レスポンス
``` json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### ✅ GET `/:code`

**短縮URLにアクセスすると元のURLにリダイレクトするAPI**

#### 📤 レスポンス
```
HTTP/1.1 302 Found
Location: https://example.com
```

## 📚 APIドキュメント（Swagger）

このAPIは Swagger（OpenAPI） を導入しており、ブラウザ上から仕様を確認できます。

- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### 使用ライブラリ

- [swaggo/swag](https://github.com/swaggo/swag)
- [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)

### ドキュメント生成コマンド

```bash
swag init --dir ./cmd/server,./internal/interface/handler/http --output ./docs --parseDependency --parseInternal
```

## 🛠 ローカル環境構築
```
# リポジトリをクローン
git clone https://github.com/SI-Monxy/shorty-link.git

# プロジェクトに移動
cd shorty-link

# Dockerで起動
docker compose up --build
```


## 📁 ディレクトリ構成（抜粋）
```
shorty-link/
├── cmd/                    # エントリポイント
├── internal/              # アプリケーション本体（クリーンアーキテクチャ）
│   ├── config/            # DB設定
│   ├── domain/            # エンティティ、リポジトリインターフェース
│   ├── usecase/           # インターフェース、DTO、ユースケース
│   ├── presenter/         # プレゼンター（出力変換）
│   ├── interface/handler/ # Ginハンドラ（API定義）
│   └── infrastructure/db/ # MySQL + GORMモデル
├── docker/                # Dockerfileなど
├── docker-compose.yml
└── README.md
```

## 🧪 今後の実装予定（TODO）
- URLの有効期限の設定
- アクセスログの保存
- CI/CD（GitHub Actions）
- Terraformによるインフラ構築
- 短縮URLのカスタムエイリアス対応


## 🧑‍💻 開発者
Shimon Iwata

GitHub: [@SI-Monxy](https://github.com/SI-Monxy)

X: [@SI-Monxy](https://x.com/SI_Monxy)

## 📄 ライセンス
MIT License
