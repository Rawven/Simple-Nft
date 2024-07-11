package util

import (
	"Nft-Go/common/db"
	"context"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func DelPageCache(ctx context.Context, prefix string, page int) {
	for i := 0; i < page; i++ {
		err := DelCache(prefix+":"+convertor.ToString(i+1), ctx)
		if err != nil {
			logx.Info(xerror.New("旁路缓存失败--删除步骤", err))
		}
	}
}

func DelCache(key string, ctx context.Context) error {
	err := db.GetRedis().Del(ctx, key).Err()
	if err != nil {
		return xerror.New("删除缓存异常", err)
	}
	return nil
}

func SetCache[T any](key string, ctx context.Context, value T) error {
	data, err := jsonx.MarshalToString(value)
	if err != nil {
		return xerror.New("缓存数据序列化异常", err)
	}
	err = db.GetRedis().Set(ctx, key, data, 0).Err()
	if err != nil {
		return xerror.New("设置缓存异常", err)
	}
	return nil
}

func GetCache[T any](key string, ctx context.Context, value *T) error {
	result, err := db.GetRedis().Get(ctx, key).Result()
	if err != nil {
		return xerror.New("获取缓存数据失败", err)
	}
	err = jsonx.UnmarshalFromString(result, &value)
	if err != nil {
		return xerror.New("缓存数据解析失败", err)
	}
	return nil
}
