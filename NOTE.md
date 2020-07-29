# イメージを固める
- gRPC
  - Inceptor: validation, auth
- Microservice?
  - auth serverを作って使いまわしたい(育てていきたい)
- Microserviceは開発の仕方やアーキテクチャを指す
- auth serverを作って使いまわしたい→これgRPCのInterceptorでいいのか？どういう感じになるかわからん
  - https://engineer.retty.me/entry/2019/12/21/171549
  - 今回はInterceptorになるのでAuthサーバはとりあえず作らないかな


# 設計
- URLどうしよう
- これはフロントの方でルーティングできるのであとから変えられる
```
# internal/proto/check
/healthcheck
/tester (tester and admin only)
/secret (admin only checker)
# internal/proto/id # ページに対するAPI UserはLogin関連になる
/team
/team/[id]
# internal/proto/scoreboard
/scoreboard
# internal/proto/challenge
/challenge
/challenge#[challenge-name] ?
```
- めんどくさいので、teamでひとつのアカウントにしてもらう(まずは簡単に)
- :ok: check


# ref
- `https://qiita.com/otanu/items/98d553d4b685a8419952`

```
yay -S protobuf

cd ~/tools
git clone https://github.com/grpc/grpc-web
cd grpc-web
sudo make install-plugin

npx create-react-app client --template typescript
npm install -D grpc-web
```
- YAMLはインデント違いで`has unknown fields`みたいなのだしてくるので、まずインデントを疑う
```
Request header field x-user-agent is not allowed by Access-Control-Allow-Headers in preflight response.
```
- CORS理解するか...普段はnginxでリバースプロキシで解決しているので別のやり方(今回はenvoy)でも通用する方法を知りたい
- Headerが禁止されているので許可してやればよかった
- `:IndentGuidesEnable` でYAMLのときはインデント可視化する

# Auth
- https://qiita.com/arenahito/items/d96e437e5e13ef800ee0 interceptorを使うらしい

# 構想
- Authとか通知とか分けてマイクロサービス風にする(サーバはひとつだけど、分割でもできるようにしてみたい)
- gRPCとMicroserviceを学ぶ

# NOTE
- https://github.com/arenahito/go-grpc-auth-demo をみながらAuth実装
- Authの中身 JWTの勉強
- :ok: go単体で動くようにする
- :ok: serverのファイル分離
- :ok: interceptorの実装
  - https://github.com/arenahito/go-grpc-auth-demo/blob/master/server/authentication.go
- Authをきちんとしたものにする
  - Auth Serverを作る
- https://www.sambaiz.net/article/174/ 参考になりそう
- SSL/TLSにする https://engineering.linecorp.com/ja/blog/combining-slackbots-into-one-with-grpc/
  - envoyでリバースプロキシしているからそこでHTTPS確保すればいいんでは
  - https://kakakakakku.hatenablog.com/entry/2019/12/06/143207
    - あってそう。server側のHTTPS化は必要ないね
- https://github.com/harlow/go-micro-services
  - example
- https://qiita.com/istsh/items/030933950f9fe8c82844
  - これexampleとしてよさそう
- Refactoring 
- https://about.sourcegraph.com/go/grpc-in-production-alan-shreve
  - Interceptor logging
- /secretへのアクセスで、tokenが必要とする。そこはInterceptorで実装
- でもこれって少なくともGetTokenをまともに実装しないとじゃないか
- https://daichan.club/container/78908
  - psqlで参考になった
- volume削除でsql initが正常に動くことがある
- https://qiita.com/Shitimi_613/items/bcd6a7f4134e6a8f0621
  - psql cheetsheet
- SQL Boiler
  - https://ken-aio.github.io/post/2019/04/01/golang-sqlboiler-select/
- Validationはインターフェース部分でかけるのが筋な気がするので、grpc interceptorに任せる
- interceptor
  - auth
  - validate
- https://stackoverflow.com/questions/5258977/are-http-headers-case-sensitive
  - HTTP Headerはcase-Incensitiveらしい。
- frontend iframe postMessage embedded crossorigin  
  - https://developer.mozilla.org/en-US/docs/Web/API/Window/postMessage
  - これやってみたいけど、普通にCookie+hash routerで実装したほうがいいような気がする
- https://daveceddia.com/tailwind-create-react-app/
  - tailwind + react
# 注意
- :ok: Access secret ENVがないときにエラー出す。token返さないようにする
  - エラーは出した。
  - errorを返すようにしたい
- :ok: GetTokenは実装できたので/secretの実装？でもDBをやっておかないとだめだなあ...結局login相当のことを実装する必要がありそう？
  - loginはusername/passwordを投げてTokenを得る。AuthではTokenを投げてinterceptorで許可を受ける
- go は小文字始まりの関数をExportできない。InitDBのように命名する
- Executor erじゃないことに気づかず時間溶かした
- grpcはmain.goのRegisterと、handler/の実際の中身の実装が必要
- HS256なら256bit, 32文字以上でないとだめ
  - ref. https://auth0.com/blog/jp-a-look-at-the-latest-draft-for-jwt-bcp/
- GroupでAdminとTesterを管理する？ホワイトリストでブロックする
- boil.infer()に、boil.Blacklist("id")のように指定すればOK
- 親の方でRouter噛ませていたら、子ではRouterを噛ませてはいけない。
```
// good
<Router>
  <Route>
</Router>
// bad (not rendering)
<Router>
  <Link>
</Router>
```
- https://tailwindcomponents.com/
  - https://tailwindcomponents.com/component/laravel-preset login
  - tailwindcss の例
