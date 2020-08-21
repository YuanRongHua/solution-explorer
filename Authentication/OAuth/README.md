
# OAuth
以微信为例：  
1.首先，a.com 的运营者需要在微信开放平台注册账号，并向微信申请使用微信登录功能。  
2.申请成功后，得到申请的 appid、appsecret。  
3.用户在 a.com 上选择使用微信登录。  
4.这时会跳转微信的 OAuth 授权登录，并带上 a.com 的回调地址。  
5.用户输入微信账号和密码，登录成功后，需要选择具体的授权范围，如：授权用户的头像、昵称等。  
6.授权之后，微信会根据拉起 a.com?code=123 ，这时带上了一个临时票据 code。  
7.获取 code 之后， a.com 会拿着 code 、appid、appsecret，向微信服务器申请 token，验证成功后，微信会下发一个 token。 
8.有了 token 之后， a.com 就可以凭借 token 拿到对应的微信用户头像，用户昵称等信息了。  
9.a.com 提示用户登录成功，并将登录状态写入 Cooke，以作为后续访问的凭证。  


# 参考资料  
- [理解OAuth 2.0](http://www.ruanyifeng.com/blog/2014/05/oauth_2_0.html)
- [OAuth,Token和JWT](https://www.jianshu.com/p/9f80be6ba2e9)
- [go-oauth2](https://github.com/go-oauth2/oauth2)
- [go-oauth2案例](https://github.com/go-oauth2/oauth2/tree/master/example)
- [authboss](https://github.com/volatiletech/authboss)
