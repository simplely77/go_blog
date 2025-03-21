package flag

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"os"
	"server/global"
	"server/model/elasticsearch"
	"server/model/other"
	"time"
)

// ElasticsearchExport 导出 ES 中的数据到 JSON 文件
func ElasticsearchExport() error {
	// 声明变量用于存储响应结果
	var response other.ESIndexResponse

	// 发起第一次搜索请求，设置查询条件
	res, err := global.ESClient.Search().
		Index(elasticsearch.ArticleIndex()).
		Scroll("1m").
		Size(1000).
		Query(&types.Query{MatchAll: &types.MatchAllQuery{}}).
		Do(context.TODO())
	if err != nil {
		return err
	}

	// 遍历第一次查询结果的文件
	for _, hit := range res.Hits.Hits {
		data := other.Data{
			ID:  hit.Id_,
			Doc: hit.Source_,
		}
		response.Data = append(response.Data, data)
	}

	// 使用 Scroll API 进行后续的滚动查询，直到没有更多数据
	for {
		res, err := global.ESClient.Scroll().ScrollId(*res.ScrollId_).Scroll("1m").Do(context.TODO())
		if err != nil {
			return err
		}
		if len(res.Hits.Hits) == 0 {
			break
		}
		for _, hit := range res.Hits.Hits {
			data := other.Data{
				ID:  hit.Id_,
				Doc: hit.Source_,
			}
			response.Data = append(response.Data, data)
		}
	}

	// 清除滚动查询，释放 Elasticsearch 上的资源
	_, err = global.ESClient.ClearScroll().ScrollId(*res.ScrollId_).Do(context.TODO())
	if err != nil {
		return err
	}

	// 生成文件名，格式为 "es_yyyyMMdd.json"
	fileName := fmt.Sprintf("es_%s.json", time.Now().Format("20060102"))

	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将 response 数据结构转换为 JSON 格式的字节数据
	byteData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// 将 JSON 格式数据写入文件
	_, err = file.Write(byteData)
	if err != nil {
		return err
	}
	return nil
}
