# HW3：ex-cloudgo-start

---

## 任务要求
1. 编程 web 服务程序 类似 cloudgo 应用。
* 要求有详细的注释
* 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
2. 使用 curl 测试，将测试结果写入 README.md
3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

## 结果
1. 运行后命令行如下：
    >  D:\workspace\gp\src\github.com\c153197984\ex-cloudgo-start> go run main.go
    > [negroni] listening on :8080

2. curl测试命令如下：
    $ curl -v http://localhost:8080/test/use
    *   Trying ::1...
    * Connected to localhost (::1) port 8080 (#0)
    > GET /test/use HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    >
    < HTTP/1.1 200 OK
    < Date: Mon, 15 Jan 2018 07:46:56 GMT
    < Content-Length: 15
    < Content-Type: text/plain; charset=utf-8
    <
    Hello test use!
    * Connection #0 to host localhost left intact

    从上面的结果我们能看出，curl能够收到本地服务器的HTTP响应。
3. ab测试命令如下：
    > PS D:\workspace\gp\src\github.com\c153197984\ex-cloudgo-start> 

    $ ab -n 1000000 -c 1000 http://localhost:8080/test/use
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 100000 requests
    Completed 200000 requests
    Completed 300000 requests
    Completed 400000 requests
    Completed 500000 requests
    Completed 600000 requests
    Completed 700000 requests
    Completed 800000 requests
    Completed 900000 requests
    Completed 1000000 requests
    Finished 1000000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            8080

    Document Path:          /test/use
    Document Length:        15 bytes

    Concurrency Level:      1000
    Time taken for tests:   32.112 seconds
    Complete requests:      1000000
    Failed requests:        0
    Total transferred:      132000000 bytes
    HTML transferred:       15000000 bytes
    Requests per second:    31140.73 [#/sec] (mean)
    Time per request:       32.112 [ms] (mean)
    Time per request:       0.032 [ms] (mean, across all concurrent requests)
    Transfer rate:          4014.23 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0   19  87.6     11    1054
    Processing:     0   13  15.1     12     417
    Waiting:        0    9  14.8      8     415
    Total:          0   32  89.9     23    1437

    Percentage of the requests served within a certain time (ms)
      50%     23
      66%     27
      75%     30
      80%     31
      90%     32
      95%     36
      98%     43
      99%    225
     100%   1437 (longest request)

以上便是ab测试的结果。
