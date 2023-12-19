package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"google.golang.org/protobuf/types/known/anypb"
	"strings"

	"bblili/service/search/internal/svc"
	"bblili/service/search/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentLogic) GetContent(in *search.GetContentRequest) (*search.GetContentResponse, error) {
	indexes := []string{"videos", "userinfos"}

	for _, indexName := range indexes {
		exists, err := l.svcCtx.ElasClient.IndexExists(indexName).Do(l.ctx)
		if err != nil {
			return nil, err
		}
		if !exists {
			createIndex, err := l.svcCtx.ElasClient.CreateIndex(indexName).Do(l.ctx)
			if err != nil {
				return nil, err
			}
			if !createIndex.Acknowledged {
				return nil, fmt.Errorf("Create index not acknowledged")
			}
		}
	}

	searchSource := elastic.NewSearchSource()
	searchSource.Size(int(in.Size))
	searchSource.From(int((in.Page - 1)))
	multiMatchQuery := elastic.NewMultiMatchQuery(in.Keyword, "title", "nick", "description")
	searchSource.Query(multiMatchQuery)

	// 配置高亮显示
	highlightBuilder := elastic.NewHighlight().
		PreTags("<span style=\"color:red\">").
		PostTags("</span>").
		RequireFieldMatch(false)
	fields := []string{"title", "nick", "description"}
	for _, field := range fields {
		highlightBuilder.Field(field)
	}
	searchSource.Highlight(highlightBuilder)
	// 执行搜索
	searchResult, err := l.svcCtx.ElasClient.Search().Index(indexes...).SearchSource(searchSource).Do(l.ctx)
	if err != nil {
		fmt.Println("Error executing search:", err)
	}

	// 处理搜索结果
	var resultList []map[string]interface{}
	// 根据hit，从map[string]interface，然后转为any
	for _, hit := range searchResult.Hits.Hits {
		var resultMap map[string]interface{}
		err := json.Unmarshal(hit.Source, &resultMap)
		if err != nil {
			return nil, err
		}
		highlightFields := hit.Highlight
		for _, field := range fields {
			highlightField, found := highlightFields[field]
			if found {
				str := strings.Join(highlightField, " ")
				resultMap[field] = str
			}
		}

		resultList = append(resultList, resultMap)
	}

	// todo: 返回带有interface的结果
	return &search.GetContentResponse{Result: []*anypb.Any{}}, nil
}
