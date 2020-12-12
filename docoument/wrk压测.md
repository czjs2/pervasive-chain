#### 使用wrk进行压测

    wrk -t8 -c200 -d30s --latency  -s test.lua http://www.bing.com
    
    Options:                                           
        -c, --connections <N>  跟服务器建立并保持的TCP连接数量 
        -d, --duration    <T>  压测时间          
        -t, --threads     <N>  使用多少个线程进行压测，压测时，是有一个主线程来控制我们设置的n个子线程间调度  
                                                        
        -s, --script      <S>  指定Lua脚本路径      
        -H, --header      <H>  为每一个HTTP请求添加HTTP头     
            --latency          在压测结束后，打印延迟统计信息  
            --timeout     <T>  超时时间    
        -v, --version          打印正在使用的wrk的详细版本信  

    
    wrk.method = "POST"
    wrk.headers["S-COOKIE2"]="a=2&b=Input&c=10.0&d=20191114***"
    wrk.body = "recent_seven=20191127_32;20191128_111"
    wrk.headers["Host"]="api.shouji.**.com"
    function response(status,headers,body)
            if status ~= 200 then --将服务器返回状态码不是200的请求结果打印出来
                    print(body)
            --      wrk.thread:stop()
            end
    end