import redis

# 连接到 Redis
r = redis.StrictRedis(host='localhost', port=6379, db=0)

# 初始化游标
cursor = 0

# 循环扫描 Redis 键
while True:
    # 执行 SCAN 命令，获取匹配的键
    cursor, keys = r.scan(cursor=cursor, match="bluebell*", count=100)

    # 如果有匹配的键，删除它们
    if keys:
        r.delete(*keys)

    # 如果游标返回 0，说明扫描完毕，退出循环
    if cursor == 0:
        break

print("所有以 bluebell 开头的键已删除。")
