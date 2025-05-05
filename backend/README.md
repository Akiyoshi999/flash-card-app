## 構成

### Protocol Buffers 定義（proto/）

-   CRUD 操作のためのインターフェースを定義
-   リクエスト、レスポンスの型定義

### DynamdDB リポジトリ（internal/repository/dynamodb/）

-   DynamoDB のやり取り担当
-   CRUD の具体的な処理

### サービス層（internal/service）

-   ビジネスロジックを実装
-   リポジトリを利用してデータ操作
-   バリデーションやデータ操作
