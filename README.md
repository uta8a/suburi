# 素振り
- golang+grpc+React+envoy+TypeScript

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

# Todo List
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


# 注意
- Access secret ENVがないときにエラー出す。token返さないようにする
