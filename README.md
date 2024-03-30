# 概要
OpenTelmetry (Otel)を利用した Trace の確認を行う場合の実装の実験です。

サーバにリクエストを行うと１～６のランダムな値を返します。

# 手順
1. rolldice_httpフォルダに移動する。
2. 以下のコマンドを実行し、docker の build と up を行う。
```
$ docker compose build
```
```
$ docker compose up -d
```
4. serverフォルダに移動する。
5. 以下のコマンドを実行し、Roll Diceサーバを立ち上げる。
```
$ make start
```
6. Roll DiceクライアントからRoll Dice サーバにリクエストを行う。
- http接続の場合: [http://localhost:8080](http://localhost:8080) に接続し、リクエストを送る。

- gRPC接続の場合: clientフォルダに移動し、以下のコマンドを実行する。
```
$ make start
```
7. [http://localhost:16686](http://localhost:16686) に接続し、Jaegerのフロントエンドに接続、Traceの確認をする。

# 目次
フォルダごとに構成が異なっています。

- [rolldice_http](https://github.com/kitasan04/otel_go/#rolldice_http)
- [rolldice_grpc](https://github.com/kitasan04/otel_go/#rolldice_grpc)
- [rolldice_grpc_collector](https://github.com/kitasan04/otel_go/#rolldice_grpc_collector)
- [filter](https://github.com/kitasan04/otel_go/#filter)

# 各フォルダーの構成図
## rolldice_http
Roll DiceサーバにRoll Diceクライアント (Webブラウザ) からリクエストを送り、その Trace をJaeger から確認できます。

### 構成図
![rolldice_httpの画像](https://github.com/kitasan04/otel_go/assets/103953052/224f7aef-e888-4593-a01b-41be536359ee)

## rolldice_grpc
Roll DiceサーバにRoll Diceクライアント (サーバ) からリクエストを送り、その Trace をJaeger から確認できます。
### 構成図
![rolldice_grpcの画像](https://github.com/kitasan04/otel_go/assets/103953052/f3996441-7af2-487b-a062-8958428f35fe)

## 3. rolldice_grpc_collector
rolldice_grpcでTraceをRoll DiceサーバからJaegerに送る際、OTel Collectorを介する。
### 構成図
![rolldice_grpc_collectorの画像](https://github.com/kitasan04/otel_go/assets/103953052/3cc247c4-b3c4-44fa-acef-45868d7ac49e)
## filter
rolldice_grpc_collectorにおいて、OTel CollectorでTraceのSpanのAttributeであるroll.valueが2の時はSpanを通さないというfilterをOtel Collectorでかける。
### 構成図
![filter](https://github.com/kitasan04/otel_go/assets/103953052/6e34a93a-d963-4797-9ca3-f379db726d5e)
