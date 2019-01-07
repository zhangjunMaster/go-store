逻辑顺序

data数据服务：
1.q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS_DATA") 向apiServers发送本地地址
2.q.Bind("dataServers") 监听dataServers，接收消息后查询文件是否在本服务器

api服务:
1.q.Bind("apiServers") 接收文件存放地址，并记录
2.q.Publish("dataServes", name) 向dataServes发送要下载的文件，获取文件在哪个服务器上
3.putStream 存文件时，随机选一个文件存放地址存文件


