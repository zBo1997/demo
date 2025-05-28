package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

// 模拟服务返回的数据结构
type UserInfo struct {
	ID    int
	Name  string
	Email string
}

type OrderInfo struct {
	OrderID  int
	Amount   float64
	Products []string
}

type Recommendation struct {
	Items []string
}

// 模拟服务调用可能失败的情况
func mockServiceCall() error {
	// 20%概率失败
	if rand.Intn(5) == 0 {
		return errors.New("service unavailable")
	}
	return nil
}

// 模拟用户服务调用
func fetchUserInfo(ctx context.Context, userID int) (*UserInfo, error) {
	if err := mockServiceCall(); err != nil {
		return nil, fmt.Errorf("user service error: %w", err)
	}

	// 模拟处理时间
	select {
	case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return &UserInfo{
		ID:    userID,
		Name:  fmt.Sprintf("User%d", userID),
		Email: fmt.Sprintf("user%d@example.com", userID),
	}, nil
}

// 模拟订单服务调用
func fetchOrderInfo(ctx context.Context, userID int) (*OrderInfo, error) {
	if err := mockServiceCall(); err != nil {
		return nil, fmt.Errorf("order service error: %w", err)
	}

	select {
	case <-time.After(time.Duration(rand.Intn(150)) * time.Millisecond):
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return &OrderInfo{
		OrderID:  rand.Intn(1000),
		Amount:   rand.Float64() * 100,
		Products: []string{"ProductA", "ProductB"},
	}, nil
}

// 模拟推荐服务调用
func fetchRecommendation(ctx context.Context, userID int) (*Recommendation, error) {
	if err := mockServiceCall(); err != nil {
		return nil, fmt.Errorf("recommendation service error: %w", err)
	}

	select {
	case <-time.After(time.Duration(rand.Intn(200)) * time.Millisecond):
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return &Recommendation{
		Items: []string{"RecommendedItem1", "RecommendedItem2"},
	}, nil
}

// 组合结果
type CombinedResult struct {
	User           *UserInfo
	Order          *OrderInfo
	Recommendation *Recommendation
}

// 使用errgroup并发获取所有数据
func fetchAllData(ctx context.Context, userID int) (*CombinedResult, error) {
	var result CombinedResult
	g, ctx := errgroup.WithContext(ctx)

	// 并发获取用户信息
	g.Go(func() error {
		userInfo, err := fetchUserInfo(ctx, userID)
		if err != nil {
			return err
		}
		result.User = userInfo
		return nil
	})

	// 并发获取订单信息
	g.Go(func() error {
		orderInfo, err := fetchOrderInfo(ctx, userID)
		if err != nil {
			return err
		}
		result.Order = orderInfo
		return nil
	})

	// 并发获取推荐信息
	g.Go(func() error {
		recommendation, err := fetchRecommendation(ctx, userID)
		if err != nil {
			return err
		}
		result.Recommendation = recommendation
		return nil
	})

	// 等待所有goroutine完成
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 模拟调用5次，观察不同结果
	for i := 0; i < 5; i++ {
		fmt.Printf("=== Attempt %d ===\n", i+1)
		result, err := fetchAllData(ctx, 123)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			continue
		}

		fmt.Printf("User: %+v\n", result.User)
		fmt.Printf("Order: %+v\n", result.Order)
		fmt.Printf("Recommendation: %+v\n\n", result.Recommendation)
	}
}
