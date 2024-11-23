基于/64的ipv6地址，请保证你也有这么多ip。ex服务器确实是会ban v6ip的。
仅作为代理

forbidden path:
- /fullimg 
- /archiver.php

特别的：
- GET /s/:key/:id-:page?redirect_to=image
  会302重定向至图片，因为eh的图片显示可能会挂，302并不识别。
- GET /g/:id/:key?redirect_to=cover 
  会302重定向至画廊的第一页的图片连接的/s/:key/:id-:page?redirect_to=image，接着重定向至第一个图片，总之是cover

在运行目录下创建.env，保存这两个键
EXHENTAI_PROXY_PREFIX
EXHENTAI_PROXY_COOKIE="ipb_member_id=【id】; ipb_pass_hash=【hash】; yay=louder; igneous=【igneous】"
从你的游览器得到。