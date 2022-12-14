package utils

import (
	"GoLog/commen"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 构建mysql命名

type ParamPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

type Paging struct {
	Page  int   `json:"page"`  // 页码
	Limit int   `json:"limit"` // 每页条数
	Total int64 `json:"total"` // 总数据条数
}

func (p *Paging) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}
	return offset
}

func (p *Paging) TotalPage() int {
	if p.Total == 0 || p.Limit == 0 {
		return 0
	}
	totalPage := int(p.Total) / p.Limit
	if int(p.Total)%p.Limit > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}

type SqlCnd struct {
	SelectCols []string     // 要查询的字段，如果为空，表示查询所有字段
	Params     []ParamPair  // 参数
	Orders     []OrderByCol // 排序
	Paging     *Paging      // 分页
}

func NewSqlCnd() *SqlCnd {
	return &SqlCnd{}
}

// 排序信息
type OrderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

func (s *SqlCnd) Cols(selectCols ...string) *SqlCnd {
	if len(selectCols) > 0 {
		s.SelectCols = append(s.SelectCols, selectCols...)
	}
	return s
}

func (s *SqlCnd) Eq(column string, args ...interface{}) *SqlCnd {
	s.Where(column+" = ?", args)
	return s
}

func (s *SqlCnd) NotEq(column string, args ...interface{}) *SqlCnd {
	s.Where(column+" <> ?", args)
	return s
}

func (s *SqlCnd) Gt(column string, args ...interface{}) *SqlCnd {
	s.Where(column+" > ?", args)
	return s
}

func (s *SqlCnd) Gte(column string, args ...interface{}) *SqlCnd {
	s.Where(column+" >= ?", args)
	return s
}

func (s *SqlCnd) Lt(column string, args ...interface{}) *SqlCnd {
	s.Where(column+" < ?", args)
	return s
}

func (s *SqlCnd) Lte(column string, args ...interface{}) *SqlCnd {
	s.Where(column+" <= ?", args)
	return s
}

func (s *SqlCnd) Like(column string, str string) *SqlCnd {
	s.Where(column+" LIKE ?", "%"+str+"%")
	return s
}

func (s *SqlCnd) Starting(column string, str string) *SqlCnd {
	s.Where(column+" LIKE ?", str+"%")
	return s
}

func (s *SqlCnd) Ending(column string, str string) *SqlCnd {
	s.Where(column+" LIKE ?", "%"+str)
	return s
}

func (s *SqlCnd) In(column string, params interface{}) *SqlCnd {
	s.Where(column+" in (?) ", params)
	return s
}

func (s *SqlCnd) NotIn(column string, params interface{}) *SqlCnd {
	s.Where(column+" not in (?) ", params)
	return s
}

func (s *SqlCnd) Where(query string, args ...interface{}) *SqlCnd {
	s.Params = append(s.Params, ParamPair{query, args})
	return s
}

func (s *SqlCnd) Asc(column string) *SqlCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: true})
	return s
}

func (s *SqlCnd) Desc(column string) *SqlCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: false})
	return s
}

func (s *SqlCnd) Limit(limit int) *SqlCnd {
	s.Page(1, limit)
	return s
}

func (s *SqlCnd) Page(page, limit int) *SqlCnd {
	if s.Paging == nil {
		s.Paging = &Paging{Page: page, Limit: limit}
	} else {
		s.Paging.Page = page
		s.Paging.Limit = limit
	}
	return s
}

func (s *SqlCnd) Build(db *gorm.DB) *gorm.DB {
	ret := db

	if len(s.SelectCols) > 0 {
		ret = ret.Select(s.SelectCols)
	}

	// where
	if len(s.Params) > 0 {
		for _, param := range s.Params {
			ret = ret.Where(param.Query, param.Args...)
		}
	}

	// order
	if len(s.Orders) > 0 {
		for _, order := range s.Orders {
			if order.Asc {
				ret = ret.Order(order.Column + " ASC")
			} else {
				ret = ret.Order(order.Column + " DESC")
			}
		}
	}

	// limit
	if s.Paging != nil && s.Paging.Limit > 0 {
		ret = ret.Limit(s.Paging.Limit)
	}

	// offset
	if s.Paging != nil && s.Paging.Offset() > 0 {
		ret = ret.Offset(s.Paging.Offset())
	}
	return ret
}

func (s *SqlCnd) Find(db *gorm.DB, out interface{}) {
	if err := s.Build(db).Find(out).Error; err != nil {
		commen.GVA_LOG.Error("未找到对应数据！", zap.Error(err))
	}
}

func (s *SqlCnd) FindOne(db *gorm.DB, out interface{}) error {
	if err := s.Limit(1).Build(db).First(out).Error; err != nil {
		return err
	}
	return nil
}

func (s *SqlCnd) Count(db *gorm.DB, model interface{}) int64 {
	ret := db.Model(model)

	// where
	if len(s.Params) > 0 {
		for _, query := range s.Params {
			ret = ret.Where(query.Query, query.Args...)
		}
	}

	var count int64
	if err := ret.Count(&count).Error; err != nil {
		commen.GVA_LOG.Error("不存在数据或查询错误，请检查！", zap.Error(err))
	}
	return count
}
