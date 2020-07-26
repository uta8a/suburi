# 素振り
- golang+grpc+React+envoy+TypeScript

# TODO
- :ok: check proto file
- :ok: App external file
- :ok: authorizationの実装 checkのみ
- :ok: GetToken 発行をする
- :ok: jwt authenticationをきちんとやる
  - https://github.com/dgrijalva/jwt-go/blob/master/cmd/jwt/app.go
    - 参考になりそう
- :ok: login dbとの接続?
- :ok: Register実装
- logoutの実装
- フロントエンドのjwtの取扱、ログインページの作成
- サーバから情報を読み込むまでローディング待機の仕組みを理解したい
- :arrow_forward: Viewを作り込む


# Bug
- :fixed: user.Login で正しいUsernameとPasswordを入れるとpanic
  - ctxをInterceptorで返していなかった
- :fixed: Authorizationが機能していない。tokenつきでアクセスしても入れない(Invalid signature)

