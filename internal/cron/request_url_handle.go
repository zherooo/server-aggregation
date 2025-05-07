package cron

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"server-aggregation/config"
	"server-aggregation/pkg/log"
	"sync"
	"time"
)

const (
	namespaceUrl = "/test"
)

// 获取结果
func HandleUrlResult(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	interval := 10 * time.Second

	// 创建 ticker，每隔 interval 发出一个信号
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("定时任务进行中...")
			// 每隔 interval 执行一次任务
			result, err := sendNameSpaceRequest(ctx)
			if err != nil {
				log.New().WithContext(ctx).Named("static_request").Error("namespace URL response信息", zap.Any("static_response", err))
				continue
			}
			//更新数据
			err = updateUrlStatus(ctx, result)
			if err != nil {
				// 更新数据失败，记录错误信息
				log.New().WithContext(ctx).Named("static_request").Error("更新失败", zap.Any("error", err))
			}

		}
	}
}

func sendNameSpaceRequest(ctx context.Context) (result map[string]interface{}, err error) {

	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "namespaces",
		"id":      1,
	}
	jsonData, err := json.Marshal(data)

	//url
	baseUrl := config.GetString("static_analysis.url")
	requestNamespacedUrl := baseUrl + namespaceUrl
	//log.New().WithContext(ctx).Named("static_request").Info("namespace URL 请求", zap.Any("static_response", requestNamespacedUrl))
	// 发送请求
	resp, err := http.Post(requestNamespacedUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Body)
		return result, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody := new(bytes.Buffer)
	if _, err = responseBody.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(responseBody.Bytes(), &result); err != nil {
		return nil, err
	}
	if success, ok := result["success"].(bool); ok && success {
		return result, nil
	} else if errorMsg, ok := result["error"].(string); ok {
		err = fmt.Errorf("error from server: %s", errorMsg)
		return
	}
	return result, nil
}

func updateUrlStatus(ctx context.Context, result map[string]interface{}) (err error) {
	//todo

	return nil
}
