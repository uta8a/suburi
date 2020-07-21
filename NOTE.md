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
# internal/proto/id
/user
/user/[id]
# internal/proto/scoreboard
/scoreboard
# internal/proto/challenge
/challenge
/challenge#[challenge-name] ?
```
- めんどくさいので、teamでひとつのアカウントにしてもらう(まずは簡単に)
- :ok: check
