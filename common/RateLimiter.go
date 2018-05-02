package common

// 令牌桶算法，系统以恒定等速率往桶里面放令牌，客户端可以获取
type RateLimiter interface{
	TryAcquire() bool
	TryAcquirePermits(permits int) bool
	Acquire() bool
	AcquirePermits(permits int) bool
}

func CreateRateLimiter() RateLimiter{
	return nil
}



// 漏桶算法，系统以恒定等速率从桶里面出来，当放入过大则溢出