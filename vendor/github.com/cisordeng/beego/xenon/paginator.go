package xenon

import (
	"github.com/cisordeng/beego"
	"github.com/cisordeng/beego/orm"
	"math"
)

type Paginator struct {
	Page int
	CountPerPage int

	totalCount int
}

type PageInfo struct {
	HasPrev bool
	HasNext bool
	Prev int
	Next int
	CurPage int
	MaxPage int
	TotalCount int
}

func NewPaginator(page int, countPerPage int) *Paginator {
	return &Paginator{
		Page: page,
		CountPerPage: countPerPage,

		totalCount: 0,
	}
}

func (page *Paginator) init(qs orm.QuerySeter) {
	totalCount64, err := qs.Count()
	PanicNotNilError(err)
	page.totalCount = int(totalCount64)
	if page.Page <= 0 {
		beego.Info("page must be greater than zero")
		page.Page = 1
	}
	if page.CountPerPage <= 0 {
		beego.Info("count_per_page must be greater than zero")
		page.CountPerPage = 10
	}
	if page.Page > page.maxPage() {
		beego.Info("page maximum exceeded")
		page.Page = page.maxPage()
	}
}

func (page *Paginator) offset() int {
	return (page.Page - 1) * page.CountPerPage
}

func (page *Paginator) curPage() int {
	return page.Page
}

func (page *Paginator) maxPage() int {
	return int(math.Ceil(float64(page.totalCount) / float64(page.CountPerPage)))
}

func (page *Paginator) hasPrev() bool {
	return page.prev() > 0
}

func (page *Paginator) hasNext() bool {
	return page.Page < page.maxPage()
}

func (page *Paginator) prev() int {
	return page.Page - 1
}

func (page *Paginator) next() int {
	if page.hasNext() {
		return page.Page + 1
	} else {
		return 0
	}
}

func (page *PageInfo) ToMap() Map {
	return Map{
		"has_prev": page.HasPrev,
		"has_next": page.HasNext,
		"prev": page.Prev,
		"next": page.Next,
		"cur_page": page.CurPage,
		"max_page": page.MaxPage,
		"total_count": page.TotalCount,
	}
}

func Paginate(qs orm.QuerySeter, page *Paginator, models interface{}) (PageInfo, error)  {
	page.init(qs)
	_, err := qs.Limit(
		page.CountPerPage,
		page.offset(),
	).All(models)

	return PageInfo{
		HasPrev:    page.hasPrev(),
		HasNext:    page.hasNext(),
		Prev:       page.prev(),
		Next:       page.next(),
		CurPage:    page.curPage(),
		MaxPage:    page.maxPage(),
		TotalCount: page.totalCount,
	}, err
}
