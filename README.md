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
