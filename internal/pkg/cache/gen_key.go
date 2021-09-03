package cache

import "fmt"

// GenBroadcastCacheKey .
func GenBroadcastCacheKey() string {
	return "Broadcast"
}

// GenBroadcastBonusCacheKey .
func GenBroadcastBonusCacheKey() string {
	return "BroadcastBonus"
}

// GenLastReloadStaticCacheKey .
func GenLastReloadStaticCacheKey() string {
	return "GenLastReloadStaticCacheKey"
}

// GenSessionCacheKey .
func GenSessionCacheKey(key string) string {
	return fmt.Sprintf("Session_%s", key)
}

// GenMessageCacheKey .
func GenMessageCacheKey(key string) string {
	return fmt.Sprintf("Message_%s", key)
}

// GenPlayerInfoCacheKey .
func GenPlayerInfoCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerInfo_%s", playerID)
}

// GenPlayerVipCacheKey .
func GenPlayerVipCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerVip_%s", playerID)
}

// GenPlayerBufferCacheKey .
func GenPlayerBufferCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerBuffer_%s", playerID)
}

// GenPlayerSpecialOfferCacheKey .
func GenPlayerSpecialOfferCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerSpecialOffer_%s", playerID)
}

// GenPlayer7DayRewardCacheKey .
func GenPlayer7DayRewardCacheKey(playerID string) string {
	return fmt.Sprintf("Player7DayReward_%s", playerID)
}

// GenPlayerStatisticsRecordCacheKey .
func GenPlayerStatisticsRecordCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerStatisticsRecord_%s", playerID)
}

// GenPlayerRankInfoCacheKey .
func GenPlayerRankInfoCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerRankInfo_%s", playerID)
}

// GenPlayerHiQuestCacheKey .
func GenPlayerHiQuestCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerHiQuest_%s", playerID)
}

// GenPlayerCashBackCacheKey .
func GenPlayerCashBackCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerCashBack_%s", playerID)
}

// GenPlayerLuckyBingoCacheKey .
func GenPlayerLuckyBingoCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerLuckyBingo_%s", playerID)
}

// GenPlayerMissionCacheKey .
func GenPlayerMissionCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerMission_%s", playerID)
}

// GenPlayerMultiplierCacheKey .
func GenPlayerMultiplierCacheKey(playerID string) string {
	return fmt.Sprintf("PlayerMultiplier_%s", playerID)
}

// GenRankRoomStatusCacheKey .
func GenRankRoomStatusCacheKey(playerID string) string {
	return fmt.Sprintf("RankRoomStatus_%s", playerID)
}

// GenRankRoomPlayerCacheKey .
func GenRankRoomPlayerCacheKey(playerID string) string {
	return fmt.Sprintf("RankRoomPlayer_%s", playerID)
}

// GenSendFriendGiftCoinCacheKey .
func GenSendFriendGiftCoinCacheKey(playerID string) string {
	return fmt.Sprintf("SendFriendGiftCoin_%s", playerID)
}

// GenSendFriendGiftStampCacheKey .
func GenSendFriendGiftStampCacheKey(playerID string) string {
	return fmt.Sprintf("SendFriendGiftStamp_%s", playerID)
}
