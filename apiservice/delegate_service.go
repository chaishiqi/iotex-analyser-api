package apiservice

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/iotexproject/iotex-analyser-api/common/votings"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type DelegateService struct {
	api.UnimplementedDelegateServiceServer
}

func (s *DelegateService) BucketInfo(ctx context.Context, req *api.BucketInfoRequest) (*api.BucketInfoResponse, error) {
	resp := &api.BucketInfoResponse{
		BucketInfoList: make([]*api.BucketInfoList, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	bucketMap, err := s.getBucketInformation(startEpoch, epochCount, delegateName)
	if err != nil {
		return nil, err
	}
	bucketInfoLists := make([]*api.BucketInfoList, 0)
	for epoch, bucketList := range bucketMap {
		bucketInfoList := &api.BucketInfoList{EpochNumber: epoch, Count: uint64(len(bucketList))}
		bucketInfo := make([]*api.BucketInfo, 0)
		for _, bucket := range bucketList {
			voterIotexAddr := bucket.VoterAddress
			voterAddr, _ := address.FromString(voterIotexAddr)
			bucketInfo = append(bucketInfo, &api.BucketInfo{
				BucketID:          bucket.BucketID,
				VoterEthAddress:   voterAddr.Hex(),
				VoterIotexAddress: voterIotexAddr,
				IsNative:          bucket.IsNative,
				Votes:             bucket.Votes,
				WeightedVotes:     bucket.WeightedVotes,
				RemainingDuration: bucket.RemainingDuration,
				StartTime:         bucket.StartTime,
				Decay:             bucket.Decay,
			})
		}
		page := req.GetPagination()
		var skip, first uint64

		if page != nil {
			skip = page.GetSkip()
			first = page.GetFirst()
		}
		if skip >= uint64(len(bucketInfo)) {
			return nil, errors.New("invalid pagination skip number for bucket info")
		}
		if uint64(len(bucketInfo))-skip < first {
			first = uint64(len(bucketInfo)) - skip
		}
		bucketInfoList.BucketInfo = bucketInfo[skip : skip+first]
		bucketInfoLists = append(bucketInfoLists, bucketInfoList)
	}
	sort.Slice(bucketInfoLists, func(i, j int) bool { return bucketInfoLists[i].EpochNumber < bucketInfoLists[j].EpochNumber })
	resp.Count = uint64(len(bucketInfoLists))
	resp.Exist = resp.Count > 0
	resp.BucketInfoList = bucketInfoLists
	return resp, nil
}

func (s *DelegateService) getBucketInformation(startEpoch, epochCount uint64, delegateName string) (map[uint64][]*votings.VotingInfo, error) {
	currentEpoch, _, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, errors.New("failed to get most recent epoch")
	}
	endEpoch := startEpoch + epochCount - 1
	if endEpoch > currentEpoch {
		endEpoch = currentEpoch
	}
	bucketInfoMap := make(map[uint64][]*votings.VotingInfo)
	for i := startEpoch; i <= endEpoch; i++ {
		voteInfoList, err := votings.GetBucketInfoByEpoch(i, delegateName)
		if err != nil {
			return nil, err
		}
		bucketInfoMap[i] = voteInfoList
	}
	return bucketInfoMap, nil
}

func (s *DelegateService) BookKeeping(ctx context.Context, req *api.BookKeepingRequest) (*api.BookKeepingResponse, error) {
	resp := &api.BookKeepingResponse{
		RewardDistribution: make([]*api.DelegateRewardDistribution, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	includeBlockReward := req.GetIncludeBlockReward()
	includeFoundationBonus := req.GetIncludeFoundationBonus()
	percentAge := req.GetPercentage()

	if percentAge > 100 {
		return nil, errors.New("percentage should be 0-100")
	}
	rewards, err := rewards.GetBookkeeping(ctx, startEpoch, epochCount, delegateName, int(percentAge), includeBlockReward, includeFoundationBonus)
	if err != nil {
		return nil, err
	}
	rds := make([]*api.DelegateRewardDistribution, 0)
	for ioAddr, amount := range rewards {
		voterAddr, _ := address.FromString(ioAddr)
		v := &api.DelegateRewardDistribution{
			VoterEthAddress:   voterAddr.Hex(),
			VoterIotexAddress: ioAddr,
			Amount:            amount.String(),
		}
		rds = append(rds, v)
	}

	sort.Slice(rds, func(i, j int) bool { return rds[i].VoterEthAddress < rds[j].VoterEthAddress })

	resp.Count = uint64(len(rds))
	if resp.Count == 0 {
		return resp, nil
	}
	resp.Exist = resp.Count > 0

	page := req.GetPagination()
	var skip, first uint64

	if page != nil {
		skip = page.GetSkip()
		first = page.GetFirst()
	}
	if skip >= uint64(resp.Count) {
		return nil, errors.New("invalid pagination skip number")
	}
	if resp.Count-skip < first {
		first = resp.Count - skip
	}
	if first == skip && first == 0 {
		resp.RewardDistribution = rds
		return resp, nil
	} else {
		resp.RewardDistribution = rds[skip : skip+first]
	}
	return resp, nil
}

func (s *DelegateService) Productivity(ctx context.Context, req *api.ProductivityRequest) (*api.ProductivityResponse, error) {
	resp := &api.ProductivityResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "select count(case when producer_name=? then block_height end) production, count(case when expected_producer_name=? then block_height end) expected_production from block_meta where epoch_num>=? and epoch_num<=?"
	var result struct {
		Production         uint64
		ExpectedProduction uint64
	}
	err := db.Raw(query, delegateName, delegateName, startEpoch, endEpoch).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	resp.Productivity = &api.Productivity{
		Exist:              result.Production > 0 || result.ExpectedProduction > 0,
		Production:         result.Production,
		ExpectedProduction: result.ExpectedProduction,
	}
	return resp, nil
}

func (s *DelegateService) Reward(ctx context.Context, req *api.RewardRequest) (*api.RewardResponse, error) {
	resp := &api.RewardResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "SELECT SUM(block_reward) block_reward, SUM(epoch_reward) epoch_reward, SUM(foundation_bonus) foundation_bonus FROM hermes_account_rewards WHERE epoch_number >= ?  AND epoch_number <= ? AND candidate_name=?"
	var result struct {
		BlockReward     string
		EpochReward     string
		FoundationBonus string
	}
	err := db.Raw(query, startEpoch, endEpoch, delegateName).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	resp.Reward = &api.Reward{
		Exist:           result.BlockReward != "0" || result.EpochReward != "0" || result.FoundationBonus != "0",
		BlockReward:     result.BlockReward,
		EpochReward:     result.EpochReward,
		FoundationBonus: result.FoundationBonus,
	}
	return resp, nil
}

func (s *DelegateService) HermesByDelegate(ctx context.Context, req *api.HermesByDelegateRequest) (*api.HermesByDelegateResponse, error) {
	resp := &api.HermesByDelegateResponse{
		VoterInfoList: make([]*api.HermesByDelegateVoterInfo, 0),
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	if count, sum, err := rewards.GetTotalHermesByDelegate(ctx, startEpoch, endEpoch, delegateName); err != nil {
		return nil, err
	} else {
		resp.Count = uint64(count)
		resp.TotalRewardsDistributed = sum
	}
	if resp.Count > 0 {
		resp.Exist = true
		skip := common.PageOffset(req.GetPagination())
		first := common.PageSize(req.GetPagination())
		result, err := rewards.GetHermesByDelegate(ctx, startEpoch, endEpoch, delegateName, skip, first)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			resp.VoterInfoList = append(resp.VoterInfoList, &api.HermesByDelegateVoterInfo{
				VoterAddress: v.Recipient,
				FromEpoch:    v.StartEpoch,
				ToEpoch:      v.EndEpoch,
				Amount:       v.Amount,
				ActionHash:   v.ActionHash,
				Timestamp:    time.Unix(v.Timestamp, 0).UTC().String(),
			})
		}
	}
	if results, err := rewards.GetHermesRatioByDelegate(ctx, startEpoch, endEpoch, delegateName); err != nil {
		return nil, err
	} else {
		resp.DistributionRatio = make([]*api.HermesByDelegateDistributionRatio, 0)
		for _, v := range results {
			resp.DistributionRatio = append(resp.DistributionRatio, &api.HermesByDelegateDistributionRatio{
				EpochNumber:          v.EpochNumber,
				BlockRewardRatio:     v.BlockRewardPercentage,
				EpochRewardRatio:     v.EpochRewardPercentage,
				FoundationBonusRatio: v.FoundationBonusPercentage,
			})
		}
	}
	return resp, nil
}

// Staking returns the staking info of the delegate
func (s *DelegateService) Staking(ctx context.Context, req *api.StakingRequest) (*api.StakingResponse, error) {
	resp := &api.StakingResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "SELECT epoch_number,total_weighted_votes,self_staking FROM hermes_voting_results WHERE epoch_number >= ? AND epoch_number <= ? AND delegate_name = ?"
	var result []struct {
		EpochNumber        uint64
		TotalWeightedVotes string
		SelfStaking        string
	}
	if err := db.Raw(query, startEpoch, endEpoch, delegateName).Scan(&result).Error; err != nil {
		return nil, err
	}
	resp.Exist = len(result) > 0
	resp.StakingInfo = make([]*api.StakingResponse_StakingInfo, 0)
	for _, v := range result {
		resp.StakingInfo = append(resp.StakingInfo, &api.StakingResponse_StakingInfo{
			EpochNumber:  v.EpochNumber,
			TotalStaking: v.TotalWeightedVotes,
			SelfStaking:  v.SelfStaking,
		})
	}
	return resp, nil
}

// ProbationHistoricalRate returns the probation historical rate of the delegate
func (s *DelegateService) ProbationHistoricalRate(ctx context.Context, req *api.ProbationHistoricalRateRequest) (*api.ProbationHistoricalRateResponse, error) {
	resp := &api.ProbationHistoricalRateResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	delegateName := req.GetDelegateName()
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "SELECT count(epoch_number) FROM hermes_voting_results WHERE epoch_number >= ? AND epoch_number <= ? AND delegate_name = ?"
	var count int
	if err := db.WithContext(ctx).Raw(query, startEpoch, endEpoch, delegateName).Scan(&count).Error; err != nil {
		return nil, err
	}
	probationExist := func(epochNumber uint64, address string) bool {
		query := "SELECT count(*) FROM probation WHERE epoch_number = ? AND address = ?"
		var count int
		if err := db.WithContext(ctx).Raw(query, epochNumber, address).Scan(&count).Error; err != nil {
			return false
		}
		return count > 0
	}
	probationCount := uint64(0)
	for i := startEpoch; i < startEpoch+epochCount; i++ {
		address, err := votings.GetOperatorAddress(delegateName, i)
		if err != nil {
			return nil, err
		}
		exist := probationExist(i, address)
		if exist {
			probationCount++
		}
	}
	rate := float64(probationCount) / float64(count)
	log.Printf("probationCount: %d, count: %d, rate: %f", probationCount, count, rate)
	resp.ProbationHistoricalRate = fmt.Sprintf("%.2f", rate)
	return resp, nil
}
