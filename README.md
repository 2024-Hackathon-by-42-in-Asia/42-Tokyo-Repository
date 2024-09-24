# 42Hackathon_2024

## **カーボントークン: 環境に配慮した配送をインセンティブ化するエコシステム**

#### 1. **概要**
カーボントークンは、食品配送における環境負荷を削減し、持続可能な運転を促進するために設計されたインセンティブシステムです。配達者がエコフレンドリーな運転を行うたびにトークンが付与され、そのトークンはブロックチェーン技術を用いて管理されます。ドライバーはこのトークンを、エコフレンドリーなサービスやクーポンとして利用することができます。

---

#### 2. **提案するソリューション**
カーボントークンは、以下のエコドライブ行動に基づいて付与されます。
- **ふんわりアクセル「eスタート」**: 発進時に穏やかな加速（0→20km/hを5秒以上かける）を実施。
- **一定速度の維持**: 走行中に一定速度を保ち、車間距離を維持。
- **減速時のアクセルオフ**: 減速時に早めにアクセルを離す行動。
- **アイドリング回避**: 停車中にアイドリングを避ける運転。
- **渋滞回避**: リアルタイム交通情報を活用し、渋滞を避けるルートを選択。
- **駐車マナー**: 交差点付近で違法駐車を避ける行動。

これらの行動は、スマートフォンのGPSやセンサーによって監視され、運転データがリアルタイムで収集されます。

---

#### 3. **トークンの利用方法**
カーボントークンは、さまざまな用途で利用可能です。その一つとして、**foodpandaクーポン**として利用することができます。これにより、ドライバーはトークンをfoodpandaのサービスで割引や特典として利用でき、配送業務を通じて得たトークンを有効活用できる仕組みが提供されます。

他の利用手段としては、エコフレンドリーなサービスや製品への交換、パートナー企業の特典などがあります。

---

#### 4. **技術基盤**
- **データ収集**: スマートフォンのGPS、加速度センサー、振動センサーを活用して、エコドライブ行動を監視します。
- **ブロックチェーン**: トークンはブロックチェーン上で管理され、透明性とセキュリティが保証されます。

---

#### 5. **持続可能性の拡大**
カーボントークンのシステムは、他の業界やサービスにも拡大することが可能です。エコフレンドリーな運転行動を促進することで、持続可能な未来に向けた新しい取り組みの土台を作ります。

---

### 結論
カーボントークンは、環境に優しい運転を奨励するだけでなく、配達者が食品配送業務を通じてトークンを貯め、foodpandaのクーポンとして利用できる新しい形のエコシステムを構築します。持続可能な運転行動を促進し、デリバリー業界全体にメリットをもたらします。

---

##　技術要件

### 1. **フロントエンド**
- **技術スタック**: Next.js (App Router)
- **機能**
  - **ログイン**: MetaMaskやウォレットを使用した認証機能。
  - **ウォレット接続**: Ethereumウォレットの接続機能を実装し、トークンの確認・管理。
  - **加速度測定**: スマートフォンの加速度センサーやGPSを利用してドライバーの運転データを取得。
  - **結果画面**: ドライバーのエコドライブ行動に基づいた結果を表示（例えばトークン数、走行データの分析結果）。
  - **トークン→クーポン変換**: 獲得したERC20トークンをクーポンに交換するUIの実装。

### 2. **サーバーサイド**
- **技術スタック**: Go (Gin framework)
- **機能**
  - **REST API**: ユーザーの運転データやトークン情報を管理するためのAPI。
  - **データ分析**: 収集された運転データ（加速度、速度、GPS位置情報など）を分析し、トークン付与基準に適合しているかを判断。
  - **ERC20トークンのミント**: 運転行動に応じたERC20トークンを自動的にミント。
  - **ERC721ドライバーライセンス**: ドライバーライセンスをERC721トークンとして発行し、運転記録に応じて更新。

### 3. **データベース**
- **技術スタック**: PostgreSQL
- **機能**
  - **ユーザーデータ管理**: ユーザーのデータや運転データ、トークン情報、トランザクションデータの保存。
  - **トークン管理**: 獲得したトークン数や利用状況の管理。

### 4. **ブロックチェーン**
- **技術スタック**: Ethereum (ERC20, ERC721)
- **機能**
  - **ERC20トークン**: エコドライブ行動に基づいて付与されるトークンをEthereum上で管理。
  - **ERC721ライセンス**: ドライバーの運転ライセンスをERC721として発行し、運転の進捗に応じて更新。
  - **Infura**: Ethereumブロックチェーンとの接続をInfuraを通じて行い、トークンの管理やトランザクションの処理を行う。
