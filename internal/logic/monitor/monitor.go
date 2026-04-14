package monitor

import (
	"context"
	"math"
	"sync"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sMonitor struct {
	logCollections []string
}

func init() {
	service.RegisterMonitor(New())
}

func New() service.IMonitor {
	return &sMonitor{
		logCollections: []string{
			dao.LOG_TEXT,
			dao.LOG_IMAGE,
			dao.LOG_AUDIO,
			dao.LOG_VIDEO,
			dao.LOG_FILE,
			dao.LOG_BATCH,
			dao.LOG_GENERAL,
		},
	}
}

// 监控中心全局实时指标
func (s *sMonitor) Global(ctx context.Context, params model.MonitorGlobalReq) (*model.MonitorGlobalRes, error) {

	nowMs := gtime.TimestampMilli()

	baseMatch := bson.M{
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
	}

	if service.Session().IsResellerRole(ctx) {
		baseMatch["rid"] = service.Session().GetRid(ctx)
	}
	if service.Session().IsUserRole(ctx) {
		baseMatch["user_id"] = service.Session().GetUserId(ctx)
	}

	res := &model.MonitorGlobalRes{}

	// RPS/TPS: 5秒窗口
	secMatch := copyMatch(baseMatch)
	secMatch["created_at"] = bson.M{"$gte": nowMs - 5000, "$lte": nowMs}

	secPipeline := []bson.M{
		{"$match": secMatch},
		{"$group": bson.M{
			"_id":          nil,
			"count":        bson.M{"$sum": 1},
			"total_tokens": bson.M{"$sum": bson.M{"$add": bson.A{"$spend.text.input_tokens", "$spend.text.output_tokens"}}},
		}},
	}

	secResults := s.aggregateAllLogs(ctx, secPipeline)
	var secCount, secTokens float64
	for _, r := range secResults {
		secCount += gconv.Float64(r["count"])
		secTokens += gconv.Float64(r["total_tokens"])
	}
	if secCount > 0 {
		res.RPS = int(math.Ceil(secCount / 5))
		res.TPS = int(math.Round(secTokens / 5))
	}

	// RPM/TPM: 60秒窗口
	minMatch := copyMatch(baseMatch)
	minMatch["created_at"] = bson.M{"$gte": nowMs - 60000, "$lte": nowMs}

	minPipeline := []bson.M{
		{"$match": minMatch},
		{"$group": bson.M{
			"_id":          nil,
			"count":        bson.M{"$sum": 1},
			"total_tokens": bson.M{"$sum": bson.M{"$add": bson.A{"$spend.text.input_tokens", "$spend.text.output_tokens"}}},
		}},
	}

	minResults := s.aggregateAllLogs(ctx, minPipeline)
	for _, r := range minResults {
		res.RPM += gconv.Int(r["count"])
		res.TPM += gconv.Int(r["total_tokens"])
	}

	return res, nil
}

// 监控中心实时性能维度分析
func (s *sMonitor) PerfBreakdown(ctx context.Context, params model.MonitorPerfBreakdownReq) (*model.MonitorPerfBreakdownRes, error) {

	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}

	dimField := dimFieldMapping(params.Dimension)

	baseMatch := bson.M{
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
	}

	if service.Session().IsResellerRole(ctx) {
		baseMatch["rid"] = service.Session().GetRid(ctx)
	}
	if service.Session().IsUserRole(ctx) {
		baseMatch["user_id"] = service.Session().GetUserId(ctx)
	}

	nowMs := gtime.TimestampMilli()

	window := int64(params.Window)
	if window <= 0 {
		window = 10
	}
	windowMs := window * 1000
	minuteMs := int64(60000)

	// 短窗口 -> RPS/TPS + 耗时 + 成功率 + Token
	secMatch := copyMatch(baseMatch)
	secMatch["created_at"] = bson.M{"$gte": nowMs - windowMs, "$lte": nowMs}

	secPipeline := []bson.M{
		{"$match": secMatch},
		{"$group": bson.M{
			"_id":            dimField,
			"count":          bson.M{"$sum": 1},
			"total_tokens":   bson.M{"$sum": bson.M{"$add": bson.A{"$spend.text.input_tokens", "$spend.text.output_tokens"}}},
			"avg_total_time": bson.M{"$avg": "$total_time"},
			"avg_conn_time":  bson.M{"$avg": "$conn_time"},
			"avg_duration":   bson.M{"$avg": "$duration"},
			"avg_internal":   bson.M{"$avg": "$internal_time"},
			"success":        bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", 1}}, 1, 0}}},
			"errors":         bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$ne": bson.A{"$status", 1}}, 1, 0}}},
		}},
	}

	secRawResults := s.aggregateAllLogs(ctx, secPipeline)
	secResult := mergeByDimAvg(secRawResults, []string{"avg_total_time", "avg_conn_time", "avg_duration", "avg_internal"})
	sortByCountDesc(secResult)
	if len(secResult) > limit {
		secResult = secResult[:limit]
	}

	// 分钟窗口 -> RPM/TPM
	minMatch := copyMatch(baseMatch)
	minMatch["created_at"] = bson.M{"$gte": nowMs - minuteMs, "$lte": nowMs}

	minPipeline := []bson.M{
		{"$match": minMatch},
		{"$group": bson.M{
			"_id":          dimField,
			"count":        bson.M{"$sum": 1},
			"total_tokens": bson.M{"$sum": "$spend.total_spend_tokens"},
		}},
	}

	minRawResults := s.aggregateAllLogs(ctx, minPipeline)
	minResult := mergeByDim(minRawResults)
	sortByCountDesc(minResult)
	if len(minResult) > limit {
		minResult = minResult[:limit]
	}

	// 合并结果
	itemMap := make(map[string]*model.MonitorPerfBreakdownItem)
	var orderedNames []string

	for _, r := range secResult {
		name := gconv.String(r["_id"])
		if name == "" {
			name = "-"
		}
		count := gconv.Float64(r["count"])
		totalTokens := gconv.Int(r["total_tokens"])
		tps := 0
		if count >= float64(window) {
			tps = totalTokens / int(window)
		} else if count > 0 {
			tps = totalTokens / int(count)
		}
		successCount := gconv.Float64(r["success"])
		successRate := float64(0)
		if count > 0 {
			successRate = successCount / count * 100
		}
		avgTokensPerReq := 0
		if count > 0 {
			avgTokensPerReq = totalTokens / int(count)
		}
		item := &model.MonitorPerfBreakdownItem{
			Name:            name,
			RPS:             int(math.Ceil(count / float64(window))),
			TPS:             tps,
			AvgTotalTime:    gconv.Int64(r["avg_total_time"]),
			AvgConnTime:     gconv.Int64(r["avg_conn_time"]),
			AvgDuration:     gconv.Int64(r["avg_duration"]),
			AvgInternalTime: gconv.Int64(r["avg_internal"]),
			SuccessRate:     successRate,
			ErrorCount:      gconv.Int(r["errors"]),
			AvgTokensPerReq: avgTokensPerReq,
		}
		itemMap[name] = item
		orderedNames = append(orderedNames, name)
	}

	for _, r := range minResult {
		name := gconv.String(r["_id"])
		if name == "" {
			name = "-"
		}
		rpm := gconv.Int(r["count"])
		tpm := gconv.Int(r["total_tokens"])
		if item, ok := itemMap[name]; ok {
			item.RPM = rpm
			item.TPM = tpm
		} else {
			itemMap[name] = &model.MonitorPerfBreakdownItem{
				Name: name,
				RPM:  rpm,
				TPM:  tpm,
			}
			orderedNames = append(orderedNames, name)
		}
	}

	items := make([]*model.MonitorPerfBreakdownItem, 0, len(orderedNames))
	seen := make(map[string]bool)
	for _, name := range orderedNames {
		if seen[name] {
			continue
		}
		seen[name] = true
		item := itemMap[name]
		item.Name = desensitizeName(params.Dimension, item.Name)
		items = append(items, item)
	}

	return &model.MonitorPerfBreakdownRes{Items: items}, nil
}

// 监控中心历史性能数据
func (s *sMonitor) PerfHistory(ctx context.Context, params model.MonitorPerfHistoryReq) (*model.MonitorPerfHistoryRes, error) {

	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}

	dimField := dimFieldMapping(params.Dimension)

	nowMs := gtime.TimestampMilli()
	var rangeMs int64
	var bucketMs int64

	switch params.Range {
	case "1h":
		rangeMs = 3600 * 1000
		bucketMs = 60 * 1000
	case "3h":
		rangeMs = 3 * 3600 * 1000
		bucketMs = 5 * 60 * 1000
	case "6h":
		rangeMs = 6 * 3600 * 1000
		bucketMs = 10 * 60 * 1000
	case "12h":
		rangeMs = 12 * 3600 * 1000
		bucketMs = 30 * 60 * 1000
	case "2d":
		rangeMs = 2 * 24 * 3600 * 1000
		bucketMs = 2 * 60 * 60 * 1000
	case "3d":
		rangeMs = 3 * 24 * 3600 * 1000
		bucketMs = 3 * 60 * 60 * 1000
	default:
		rangeMs = 24 * 3600 * 1000
		bucketMs = 60 * 60 * 1000
	}

	startMs := nowMs - rangeMs
	bucketSeconds := float64(bucketMs) / 1000

	baseMatch := bson.M{
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
		"created_at":     bson.M{"$gte": startMs, "$lte": nowMs},
	}

	if service.Session().IsResellerRole(ctx) {
		baseMatch["rid"] = service.Session().GetRid(ctx)
	}
	if service.Session().IsUserRole(ctx) {
		baseMatch["user_id"] = service.Session().GetUserId(ctx)
	}

	// 第一步: 找出 top N 维度
	topPipeline := []bson.M{
		{"$match": baseMatch},
		{"$group": bson.M{
			"_id":   dimField,
			"count": bson.M{"$sum": 1},
		}},
	}

	topRawResults := s.aggregateAllLogs(ctx, topPipeline)
	topMerged := mergeByDim(topRawResults)
	sortByCountDesc(topMerged)
	if len(topMerged) > limit {
		topMerged = topMerged[:limit]
	}

	topNames := make([]any, 0, len(topMerged))
	topNameStrs := make([]string, 0, len(topMerged))
	for _, r := range topMerged {
		raw := r["_id"]
		name := gconv.String(raw)
		if name == "" {
			name = "-"
			raw = "-"
		}
		topNames = append(topNames, raw)
		topNameStrs = append(topNameStrs, name)
	}

	if len(topNames) == 0 {
		return &model.MonitorPerfHistoryRes{
			Dates:  make([]string, 0),
			Series: make(map[string][]float64),
		}, nil
	}

	// 第二步: 按时间桶聚合
	histMatch := copyMatch(baseMatch)
	histMatch[dimFieldName(params.Dimension)] = bson.M{"$in": topNames}

	bucketExpr := bson.M{
		"$multiply": bson.A{
			bson.M{"$floor": bson.M{
				"$divide": bson.A{"$created_at", bucketMs},
			}},
			bucketMs,
		},
	}

	pipeline := []bson.M{
		{"$match": histMatch},
		{"$group": bson.M{
			"_id": bson.M{
				"bucket": bucketExpr,
				"dim":    dimField,
			},
			"count":          bson.M{"$sum": 1},
			"total_tokens":   bson.M{"$sum": bson.M{"$add": bson.A{"$spend.text.input_tokens", "$spend.text.output_tokens"}}},
			"avg_total_time": bson.M{"$avg": "$total_time"},
			"avg_conn_time":  bson.M{"$avg": "$conn_time"},
			"avg_duration":   bson.M{"$avg": "$duration"},
			"avg_internal":   bson.M{"$avg": "$internal_time"},
			"success":        bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", 1}}, 1, 0}}},
			"errors":         bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$ne": bson.A{"$status", 1}}, 1, 0}}},
		}},
		{"$sort": bson.M{"_id.bucket": 1}},
	}

	aggResult := s.aggregateAllLogs(ctx, pipeline)
	aggResult = mergeBucketedResults(aggResult)

	// 生成时间桶列表
	bucketCount := int(rangeMs / bucketMs)
	firstBucket := (startMs/bucketMs)*bucketMs + bucketMs
	dates := make([]string, 0, bucketCount)
	bucketIndex := make(map[int64]int)

	for i := 0; i < bucketCount; i++ {
		bkt := firstBucket + int64(i)*bucketMs
		if bkt > nowMs {
			break
		}
		t := gtime.NewFromTimeStamp(bkt / 1000)
		var label string
		if bucketMs >= 3600*1000 {
			label = t.Format("01-02 15:04")
		} else {
			label = t.Format("15:04")
		}
		dates = append(dates, label)
		bucketIndex[bkt] = i
	}

	series := make(map[string][]float64)
	for _, name := range topNameStrs {
		series[name] = make([]float64, len(dates))
	}

	metric := params.Metric
	for _, r := range aggResult {
		bucketTs, dn := parseBucketDimKey(r["_id"])

		idx, ok := bucketIndex[bucketTs]
		if !ok {
			continue
		}
		arr, exists := series[dn]
		if !exists {
			continue
		}

		count := gconv.Float64(r["count"])
		totalTokens := gconv.Float64(r["total_tokens"])
		successCount := gconv.Float64(r["success"])

		var val float64
		switch metric {
		case "rps":
			val = math.Ceil(count / bucketSeconds)
		case "tps":
			if count > 0 {
				val = math.Round(totalTokens / bucketSeconds)
			}
		case "rpm":
			val = math.Round(count / (bucketSeconds / 60))
		case "tpm":
			val = math.Round(totalTokens / (bucketSeconds / 60))
		case "avg_total_time":
			val = math.Round(gconv.Float64(r["avg_total_time"]))
		case "avg_conn_time":
			val = math.Round(gconv.Float64(r["avg_conn_time"]))
		case "avg_duration":
			val = math.Round(gconv.Float64(r["avg_duration"]))
		case "avg_internal_time":
			val = math.Round(gconv.Float64(r["avg_internal"]))
		case "success_rate":
			if count > 0 {
				val = math.Round(successCount/count*10000) / 100
			}
		case "error_count":
			val = gconv.Float64(r["errors"])
		case "input_tokens":
			val = totalTokens
		case "output_tokens":
			val = 0
		case "avg_tokens_per_req":
			if count > 0 {
				val = math.Round(totalTokens / count)
			}
		}

		arr[idx] = val
	}

	desensitizedSeries := make(map[string][]float64, len(series))
	for name, values := range series {
		desensitizedSeries[desensitizeName(params.Dimension, name)] = values
	}

	return &model.MonitorPerfHistoryRes{
		Dates:  dates,
		Series: desensitizedSeries,
	}, nil
}

// 并发查询所有日志集合并合并结果
func (s *sMonitor) aggregateAllLogs(ctx context.Context, pipeline []bson.M) []map[string]any {
	var mu sync.Mutex
	var wg sync.WaitGroup
	merged := make([]map[string]any, 0)

	for _, coll := range s.logCollections {
		wg.Add(1)
		go func(collection string) {
			defer wg.Done()
			result := make([]map[string]any, 0)
			if err := dao.Aggregate(ctx, db.DefaultDatabase, collection, pipeline, &result); err != nil {
				logger.Error(ctx, err)
				return
			}
			mu.Lock()
			merged = append(merged, result...)
			mu.Unlock()
		}(coll)
	}
	wg.Wait()
	return merged
}

// 将多集合聚合结果按维度名合并(累加数值字段)
func mergeByDim(results []map[string]any) []map[string]any {
	dimMap := make(map[string]map[string]any)
	var order []string

	for _, r := range results {
		name := gconv.String(r["_id"])
		if name == "" {
			name = "-"
		}
		if existing, ok := dimMap[name]; ok {
			for k, v := range r {
				if k == "_id" {
					continue
				}
				existing[k] = gconv.Float64(existing[k]) + gconv.Float64(v)
			}
		} else {
			entry := make(map[string]any)
			for k, v := range r {
				entry[k] = v
			}
			dimMap[name] = entry
			order = append(order, name)
		}
	}

	merged := make([]map[string]any, 0, len(order))
	for _, name := range order {
		merged = append(merged, dimMap[name])
	}
	return merged
}

// 合并并计算平均值(需要count加权)
func mergeByDimAvg(results []map[string]any, avgFields []string) []map[string]any {
	type dimEntry struct {
		data  map[string]float64
		count float64
	}
	dimMap := make(map[string]*dimEntry)
	var order []string

	avgSet := make(map[string]bool)
	for _, f := range avgFields {
		avgSet[f] = true
	}

	for _, r := range results {
		name := gconv.String(r["_id"])
		if name == "" {
			name = "-"
		}
		count := gconv.Float64(r["count"])
		if existing, ok := dimMap[name]; ok {
			for k, v := range r {
				if k == "_id" {
					continue
				}
				if avgSet[k] {
					totalCount := existing.count + count
					if totalCount > 0 {
						existing.data[k] = (existing.data[k]*existing.count + gconv.Float64(v)*count) / totalCount
					}
				} else {
					existing.data[k] = existing.data[k] + gconv.Float64(v)
				}
			}
			existing.count += count
		} else {
			entry := &dimEntry{data: make(map[string]float64), count: count}
			for k, v := range r {
				if k == "_id" {
					continue
				}
				entry.data[k] = gconv.Float64(v)
			}
			dimMap[name] = entry
			order = append(order, name)
		}
	}

	merged := make([]map[string]any, 0, len(order))
	for _, name := range order {
		entry := dimMap[name]
		m := map[string]any{"_id": name}
		for k, v := range entry.data {
			m[k] = v
		}
		merged = append(merged, m)
	}
	return merged
}

func copyMatch(src bson.M) bson.M {
	dst := bson.M{}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// 按 count 字段降序排序
func sortByCountDesc(results []map[string]any) {
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if gconv.Float64(results[j]["count"]) > gconv.Float64(results[i]["count"]) {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}

// 根据维度对密钥类名称脱敏
func desensitizeName(dimension, name string) string {
	if name == "" || name == "-" {
		return name
	}
	if dimension == "key" || dimension == "app_key" {
		return util.Desensitize(name)
	}
	return name
}

// 从复合 _id 中提取 bucket 和 dim
func parseBucketDimKey(idRaw any) (int64, string) {
	var bucketTs int64
	var dimName string
	switch id := idRaw.(type) {
	case bson.D:
		for _, elem := range id {
			if elem.Key == "bucket" {
				bucketTs = gconv.Int64(elem.Value)
			}
			if elem.Key == "dim" {
				dimName = gconv.String(elem.Value)
			}
		}
	case bson.M:
		bucketTs = gconv.Int64(id["bucket"])
		dimName = gconv.String(id["dim"])
	case map[string]any:
		bucketTs = gconv.Int64(id["bucket"])
		dimName = gconv.String(id["dim"])
	}
	if dimName == "" {
		dimName = "-"
	}
	return bucketTs, dimName
}

// 合并多集合桶聚合结果
func mergeBucketedResults(results []map[string]any) []map[string]any {
	type compositeKey struct {
		bucket int64
		dim    string
	}
	merged := make(map[compositeKey]map[string]float64)
	var order []compositeKey

	avgFields := map[string]bool{
		"avg_total_time": true,
		"avg_conn_time":  true,
		"avg_duration":   true,
		"avg_internal":   true,
	}

	for _, r := range results {
		bucketTs, dimName := parseBucketDimKey(r["_id"])
		key := compositeKey{bucket: bucketTs, dim: dimName}

		count := gconv.Float64(r["count"])
		if existing, ok := merged[key]; ok {
			existingCount := existing["count"]
			for k, v := range r {
				if k == "_id" {
					continue
				}
				val := gconv.Float64(v)
				if avgFields[k] {
					totalCount := existingCount + count
					if totalCount > 0 {
						existing[k] = (existing[k]*existingCount + val*count) / totalCount
					}
				} else {
					existing[k] = existing[k] + val
				}
			}
			existing["count"] = existingCount + count
		} else {
			entry := make(map[string]float64)
			for k, v := range r {
				if k == "_id" {
					continue
				}
				entry[k] = gconv.Float64(v)
			}
			merged[key] = entry
			order = append(order, key)
		}
	}

	out := make([]map[string]any, 0, len(order))
	for _, key := range order {
		entry := merged[key]
		m := map[string]any{
			"_id": map[string]any{
				"bucket": key.bucket,
				"dim":    key.dim,
			},
		}
		for k, v := range entry {
			m[k] = v
		}
		out = append(out, m)
	}
	return out
}

// 维度到MongoDB字段映射
func dimFieldMapping(dimension string) string {
	switch dimension {
	case "model_agent":
		return "$model_agent.name"
	case "user":
		return "$user_id"
	case "app":
		return "$app_id"
	case "app_key":
		return "$creator"
	case "provider":
		return "$provider_name"
	case "key":
		return "$key"
	case "group":
		return "$spend.group_name"
	default:
		return "$model"
	}
}

// 维度到MongoDB文档字段名映射(不带$前缀)
func dimFieldName(dimension string) string {
	switch dimension {
	case "model_agent":
		return "model_agent.name"
	case "user":
		return "user_id"
	case "app":
		return "app_id"
	case "app_key":
		return "creator"
	case "provider":
		return "provider_name"
	case "key":
		return "key"
	case "group":
		return "spend.group_name"
	default:
		return "model"
	}
}
