## システム構成/ディレクトリ構成

### ディレクトリ構成

```bash
/flashmaster-monorepo
├── frontend/       # React + Vite + gRPC‑Web client
├── backend/        # AppSync schema, Lambda (optional)
├── grpc-server/    # Go (or TS) の gRPC サーバー & proto 定義
└── infra/          # CDK(TS) or Terraform で ALB, ECS, AppSync, DynamoDB 定義
```
