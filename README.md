# 日志库
基于zap log封装的一个日志库，存按照个人使用习惯封装的。

1. 方便的打日志， log.Info("msg", "data", "test data") ，第一个参数为日志消息，之后的参数为日志参数，如果日志参数数量为偶数则打印keyValue日志，而奇数则打印数组格式的日志参数
2. 支持模块前缀 支持模块前缀, 方便日志分类
3. 支持控制台和文件形式输出日志
4. 支持日志分割，按天、按小时分割，error和非error日志单独为一个日志文件
5. 支持模块级别的日志开关
```
cLog, err := NewConsoleLog()
cLog.level = LevelInfo
assert.Nil(t, err)
SetModuleSwitchFunc(func(moduleName string) bool {
    if moduleName == "login" {
        return true
    }
    return false
})
cLog.CopyWithModuleName("payment").Debug("pay info", "data", "orderInfo")
cLog.CopyWithModuleName("login").Debug("login process", "data", "loginInfo")
```
会输出
```shell
{"level":"debug","ts":"2023-07-05T15:03:16.305+0800","caller":"llog/log_test.go:91","msg":"login process","data":"loginInfo","module":"login"}
```
开启指定模块的情况下，该模块日志会无视日志级别，打印所有日志，这样做的原因是BUG写多了。新功能上线的时候开启一段时间
观察是否有问题，如果有问题就可以看到更多的日志，一段时间没问题的时候，把开关关掉即可。