# how launcher works
Launcher could run in three different mode and each one is loaded differently and each has different tradeoff.
Launcher => Loader => Env => factory => PlugApp
loading js in origin from random plug would be bummer from "security" wise.(someplugApp would just steam admin token can do anthing)
Security model of web is beshed on the Origin which is basically (domain + protocal + port)
https://evil.example.com cannot access data (cookies,localstorage etc...) from https://admin.exmaple.com.

- Sub origin
    so we can them in seperate suborigin like 
    http://adminconsole.example.com
    http://app1.example.com
    http://app2.example.com
    this means we need way to modify DNS dynamically as new app is installed or lookup apps record answer to wildcard DNS request.
- Iframe
    this a way web gives to run something under seperate origin so we do not need DNS involvement just run run app inside `<iframe>` 
- Native dom
    this mode do not have anysandbox at all, it might be ok if you trust the plug to not be evil or this might be used for mode when you are loading your application in seperate origin than admin console origin or you are using temphia as a heahless CMS and you want to run some whitelisted plugs dynamically run at some block. 

## Sub Origin
    TODO => actual steps
## Iframe
    TODO => actual steps
## Native dom
    TODO => actual steps
