package redis

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"math"
	"time"
)

const (
	OneWeekInSeconds         = 7 * 24 * 60 * 60
	VoteScore        float64 = 432
	PostPerAge               = 10
)

func PostVote(postID, userID string, v float64) (err error) {
	postTime := rdb.ZScore(ctx, KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	key := KeyPostVotedZSetPf + postID
	ov := rdb.ZScore(ctx, key, userID).Val()
	diffabs := math.Abs(ov - v)
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(ctx, key, redis.Z{
		Score:  v,
		Member: userID,
	})
	pipeline.ZIncrBy(ctx, KeyPostScoreZSet, VoteScore*diffabs*v, postID)
	zap.L().Debug("votes changed", zap.Float64("old value", ov), zap.Float64("new value", v))
	switch math.Abs(ov) - math.Abs(v) {
	case 1:
		// 取消投票 ov=1/-1 v=0
		// 投票数-1
		pipeline.HIncrBy(ctx, KeyPostInfoHashPf+postID, "votes", -1)
	case 0:
		// 反转投票 ov=-1/1 v=1/-1
		// 投票数不用更新
	case -1:
		// 新增投票 ov=0 v=1/-1
		// 投票数+1
		pipeline.HIncrBy(ctx, KeyPostInfoHashPf+postID, "votes", 1)
	default:
		// 已经投过票了
		return ErrorVoted
	}
	_, err = pipeline.Exec(ctx)
	return
}

func CreatePost(postID, userID, title, summary, communityName string) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyPostVotedZSetPf + postID
	communityKey := KeyCommunityPostSetPf + communityName
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(ctx, votedKey, redis.Z{
		Score:  1,
		Member: userID,
	})
	pipeline.Expire(ctx, votedKey, OneWeekInSeconds*time.Second)

	pipeline.HMSet(ctx, KeyPostInfoHashPf+postID, postInfo)

	pipeline.ZAdd(ctx, KeyPostScoreZSet, redis.Z{ // 添加到分数的ZSet
		Score:  VoteScore,
		Member: postID,
	})
	pipeline.ZAdd(ctx, KeyPostTimeZSet, redis.Z{ // 添加到时间的ZSet
		Score:  now,
		Member: postID,
	})
	pipeline.SAdd(ctx, communityKey, postID) // 添加到对应版块
	_, err = pipeline.Exec(ctx)
	return
}
func GetPost(order string, page int64) []map[string]string {
	key := KeyPostTimeZSet
	if order == "score" {
		key = KeyPostScoreZSet
	}
	start := (page - 1) * PostPerAge
	end := start + PostPerAge - 1
	ids := rdb.ZRevRange(ctx, key, start, end).Val()
	postList := make([]map[string]string, 0, len(ids))
	for _, id := range ids {
		postData := rdb.HGetAll(ctx, KeyPostInfoHashPf+id).Val()
		postData["id"] = id
		postList = append(postList, postData)
	}
	return postList
}
