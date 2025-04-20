## DynamoDB の構成

コスト削減のため、必要になるまで単一のテーブルとする。

| 型               | PK            | SK                          |
| ---------------- | ------------- | --------------------------- |
| ユーザーメタ情報 | USER#<userId> | METADATA#                   |
| デッキ           | USER#<userId> | DECK#<deckId>               |
| カード           | USER#<userId> | DECK#<deckId>#CARD#<cardId> |

### GSI

#### GSI1: カード一覧取得用

```text
PK = USER#<userId>
SK = DECK#<deckId>#CARD#<cardId
```

#### GSI2: 期日到来カード

```text
PK = USER#<userId>
SK = nextReview
```
