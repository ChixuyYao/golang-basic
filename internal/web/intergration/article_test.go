package intergration

import (
	"bytes"
	"encoding/json"
	"golang/internal/repository/dao"
	"golang/internal/web/intergration/startup"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// 测试套件
type ArticleHandlerSuite struct {
	suite.Suite
	db     *gorm.DB
	server *gin.Engine
}

func (s *ArticleHandlerSuite) SetupSuite() {
	s.db = startup.InitDB()
	s.server = startup.InitWebServer()
}

func (s *ArticleHandlerSuite) TearDownTest() {
	s.db.Exec("truncate table `articles`")
}

func (s *ArticleHandlerSuite) TestEdit() {
	t := s.T()
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)
		// 前台提供JSON
		art Article

		wantCode int
		wantRes  Result[string]
	}{
		{
			name:   "新增帖子信息",
			before: func(t *testing.T) {},
			after: func(t *testing.T) {
				// 验证信息是否保存
				var art dao.Article
				err := s.db.Where("author_id=?", "123").
					First(&art).Error
				assert.NoError(t, err)
				assert.True(t, art.CreatedAt > 0)
				assert.True(t, art.UpdatedAt > 0)
				assert.Equal(t, "This is Title", art.Title)
				assert.Equal(t, "This is Content........", art.Content)
				// 防止数据互相污染,truncate table xxx
			},
			art: Article{
				Title:   "This is Title",
				Content: "This is Content........",
			},
			wantCode: http.StatusOK,
			wantRes: Result[string]{
				// 期望Id为1
				Data: "1",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			defer tc.after(t)

			// 准备请求数据
			requestBody, err := json.Marshal(tc.art)
			assert.NoError(t, err)

			// 准备Request请求,记录Recorder
			req, err := http.NewRequest(http.MethodPost,
				"/articles/edit",
				bytes.NewReader(requestBody))

			req.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			recorder := httptest.NewRecorder()

			// 执行代码
			s.server.ServeHTTP(recorder, req)
			// 断言结果
			assert.Equal(t, tc.wantCode, recorder.Code)
			if tc.wantCode != http.StatusOK {
				return
			}
			var res Result[string]
			err = json.NewDecoder(recorder.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

// 提供测试入口文件
func TestArticleHandler(t *testing.T) {
	// 测试套件开启测试
	suite.Run(t, &ArticleHandlerSuite{})
}

type Result[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
	//Meta any    `json:"meta"`
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
