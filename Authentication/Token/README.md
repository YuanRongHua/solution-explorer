# Token登录

## 实现流程：
首次登录时：
1.用户访问网站A的页面a，输入密码。
2.服务器验证账号密码，正确则生成token返回给客户端，客户端自行保存。
3.登录成功，允许访问。

后续访问时：
1.用户访问网站A的页面b，请求时传递token，服务端验证token。
2.一致就允许访问，不一致就踢回登录。

## 特点：
服务器端不需要存放 Token，所以不会对服务器端造成压力，即使是服务器集群，也不需要增加维护成本。
Token 可以存放在前端任何地方，可以不用保存在 Cookie 中，提升了页面的安全性。
Token 下发之后，只要在生效时间之内，就一直有效，如果服务器端想收回此 Token 的权限，并不容易。

## Token生成方式：
最常见的 Token 生成方式是使用 JWT（Json Web Token），它是一种简洁的，自包含的方法用于通信双方之间以 JSON 对象的形式安全的传递信息。
使用 Token 后，服务器端并不会存储 Token，那怎么判断客户端发过来的 Token 是合法有效的呢？
答案其实就在 Token 字符串中，其实 Token 并不是一串杂乱无章的字符串，而是通过多种算法拼接组合而成的字符串，我们来具体分析一下。
JWT 算法主要分为 3 个部分：header（头信息），playload（消息体），signature（签名）。  
header 部分指定了该 JWT 使用的签名算法:  
`` header = '{"alg":"HS256","typ":"JWT"}'   // `HS256` 表示使用了 HMAC-SHA256 来生成签名。``  

playload 部分表明了 JWT 的意图：  
`` payload = '{"loggedInAs":"admin","iat":1422779638}'     //iat 表示令牌生成的时间``  
signature 部分为 JWT 的签名，主要为了让 JWT 不能被随意篡改，签名的方法分为两个步骤：
1.输入 base64url 编码的 header 部分 、base64url 编码的 playload 部分，输出 unsignedToken。
2.输入服务器端私钥、unsignedToken，输出 signature 签名。
