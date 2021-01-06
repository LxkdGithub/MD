python2.X 有这些库名可用: urllib, urllib2, urllib3, httplib, httplib2, requests


python3.X 有这些库名可用: urllib, urllib3, httplib2, requests

httplib库比较底层基本不用，所以使用urllib或者为简洁实用requests即可

urllib3和requests不是标准库
urllib3 提供线程安全连接池和文件post支持,与urllib及urllib2的关系不大
所以urllib3也不常用

而python3和2区别就是urllib与urllib2的许多功能到python3的urllib中变成了urllib.xxx.---, 比如quote变成了urllib.request.quote




|Python 2	|Python 3
|--- |--- |
|urllib.urlretrieve()	|urllib.request.urlretrieve()|
|urllib.urlcleanup()	|urllib.request.urlcleanup()
|urllib.quote()	|urllib.parse.quote()
|urllib.quote_plus()	|urllib.parse.quote_plus()
|urllib.unquote()	|urllib.parse.unquote()
|urllib.unquote_plus()	|urllib.parse.unquote_plus()
|urllib.urlencode |urllib.parse.urlencode()
|urllib.pathname2url()	|urllib.request.pathname2url()
|urllib.url2pathname()	|urllib.request.url2pathname()
|urllib.getproxies()	|urllib.request.getproxies()
|urllib.URLopener	|urllib.request.URLopener
|urllib.FancyURLopener	|urllib.request.FancyURLopener
|urllib.ContentTooShortError	|urllib.error.ContentTooShortError
|urllib2.urlopen()	|urllib.request.urlopen()
|urllib2.install_opener()	|urllib.request.install_opener()
|urllib2.build_opener()	|urllib.request.build_opener()
|urllib2.URLError	|urllib.error.URLError
|urllib2.HTTPError	|urllib.error.HTTPError
|urllib2.Request	|urllib.request.Request
|urllib2.OpenerDirector	|urllib.request.OpenerDirector
|urllib2.BaseHandler	|urllib.request.BaseHandler
|urllib2.HTTPDefaultErrorHandler	|urllib.request.HTTPDefaultErrorHandler
|urllib2.HTTPRedirectHandler	|urllib.request.HTTPRedirectHandler
|urllib2.HTTPCookieProcessor	|urllib.request.HTTPCookieProcessor
|urllib2.ProxyHandler	|urllib.request.ProxyHandler
|urllib2.HTTPPasswordMgr	|urllib.request.HTTPPasswordMgr
|urllib2.HTTPPasswordMgrWithDefaultRealm	|urllib.request.HTTPPasswordMgrWithDefaultRealm
|urllib2.AbstractBasicAuthHandler	|urllib.request.AbstractBasicAuthHandler
|urllib2.HTTPBasicAuthHandler	|urllib.request.HTTPBasicAuthHandler
|urllib2.ProxyBasicAuthHandler	|urllib.request.ProxyBasicAuthHandler
|urllib2.AbstractDigestAuthHandler	|urllib.request.AbstractDigestAuthHandler
|urllib2.HTTPDigestAuthHandler	|urllib.request.HTTPDigestAuthHandler
|urllib2.ProxyDigestAuthHandler |	urllib.request.ProxyDigestAuthHandler
|urllib2.HTTPHandler	|urllib.request.HTTPHandler
|urllib2.HTTPSHandler	|urllib.request.HTTPSHandler
|urllib2.FileHandler	|urllib.request.FileHandler
|urllib2.FTPHandler	|urllib.request.FTPHandler
|urllib2.CacheFTPHandler	|urllib.request.CacheFTPHandler
|urllib2.UnknownHandler	| urllib.request.UnknownHandler


[请求范例链接](https://ld246.com/article/1574254515672)

