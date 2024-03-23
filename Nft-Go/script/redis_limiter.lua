local key = KEYS[1]
local maxTokens = tonumber(ARGV[1])
local tokenRate = tonumber(ARGV[2])
local currentTime = tonumber(ARGV[3])
-- 获取key对应的桶
local result = redis.call("hmget", key, "tokens", "timestamp")
local currentTokens = tonumber(result[1])
local lastTimestamp = tonumber(result[2])
-- 如果桶不存在 填满桶
if not currentTokens then
    currentTokens = maxTokens
else
-- 计算桶中的令牌数
    local timePassed = currentTime - lastTimestamp
    local tokensToAdd = timePassed * tokenRate
    currentTokens = math.min(currentTokens + tokensToAdd, maxTokens)
end
-- 如果桶中有令牌 则消耗一个令牌
if currentTokens and currentTokens > 0 then
    redis.call("hset", key, "tokens", currentTokens - 1)
    redis.call("hset", key, "timestamp", currentTime)
    return 1
else
-- 如果桶中没有令牌 则返回0
    return 0
end
