# CROS

在浏览器中，有一些 HTTP 头是被限制修改的，这些头通常是出于安全和规范的考虑。以下是一些不允许被修改的头部列表：

### 不允许修改的 HTTP 头
1. **以 `Sec-` 开头的头部**：这些头部用于安全相关的功能。
2. **以 `Proxy-` 开头的头部**：这些头部与代理服务器的行为相关。
3. **其他特定头部**：
   - `Accept-Charset`
   - `Accept-Encoding`
   - `Access-Control-Request-Headers`
   - `Access-Control-Request-Method`
   - `Connection`
   - `Content-Length`
   - `Cookie`
   - `Date`
   - `DNT`（Do Not Track）
   - `Expect`
   - `Permissions-Policy`
   - `Host`
   - `Keep-Alive`
   - `Origin`
   - `Referer`
   - `TE`
   - `Trailer`
   - `Transfer-Encoding`
   - `Upgrade`
   - `Via` [[1]](https://developer.mozilla.org/zh-CN/docs/Glossary/Forbidden_header_name)[[2]](https://www.kxblog.com/article-111.html).

### 说明
这些头部的修改受到限制是因为浏览器希望保持对这些关键头部的控制，以确保安全性和一致性。例如，`Cookie` 和 `Authorization` 头部的修改可能会导致安全漏洞，因此浏览器会阻止这些头部的直接修改。

通过了解这些不允许修改的头部，开发者可以更好地设计和实现跨域请求和其他网络交互。

---
Learn more:
1. [禁止修改的标头 - MDN Web 文档术语表：Web 相关术语的定义 | MDN](https://developer.mozilla.org/zh-CN/docs/Glossary/Forbidden_header_name)
2. [禁止修改的Header头（Forbidden header name）- 开心博客](https://www.kxblog.com/article-111.html)
3. [禁止修改的响应标头 - MDN Web 文档术语表：Web 相关术语的定义 | MDN](https://developer.mozilla.org/zh-CN/docs/Glossary/Forbidden_response_header_name)


https://fetch.spec.whatwg.org/#forbidden-header-name
https://httptoolkit.tech/will-it-cors



如果遇到问题 

https://httptoolkit.com/will-it-cors/
https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type
https://fetch.spec.whatwg.org/#cors-safelisted-request-header
https://fetch.spec.whatwg.org/#forbidden-header-name

看这些


简单请求

同时满足以下条件

是POST或者GET
mime-type是 plain/text, application/xxx-urlencode, dataform
不需要candidate