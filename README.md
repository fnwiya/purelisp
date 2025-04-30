# PureLISP

純 LISP の Go 言語実装です。

## 特徴

- 基本的な LISP の機能を実装
- 以下の基本的な関数をサポート:
  - `atom`: 値がアトムかどうかを判定
  - `eq`: 2 つの値が同一かどうかを判定
  - `car`: リストの最初の要素を取得
  - `cdr`: リストの残りの要素を取得
  - `cons`: 2 つの要素から新しいリストを作成
- 特殊形式:
  - `quote`
  - `lambda`
  - `cond`
  - `define`

## 使用方法

```bash
go run main.go
```

## ライセンス

MIT License
