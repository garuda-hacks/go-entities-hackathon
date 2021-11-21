package response

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

var paging = Pagination{
	TotalRecords: 30,
	TotalPage:    5,
	LimitPage:    6,
	CurrentPage:  1,
}

func TestPagination_GetPage(t *testing.T) {
	assert := assert2.New(t)

	var getPage = paging.GetPage()
	assert.Equal(1, getPage, "page equals")

	paging.CurrentPage = 2
	getPage = paging.GetPage()
	assert.Equal(2, getPage, "Change Page current")

	paging.CurrentPage = -3
	getPage = paging.GetPage()
	assert.Equal(1, getPage, "always return 0 when negative")
	assert.NotEqual(-3, getPage, "should not be equal")
}

func TestPagination_GetLimit(t *testing.T) {
	assert := assert2.New(t)

	var getLimit = paging.GetLimit()
	assert.Equal(6, getLimit, "Limit equal as defined")

	paging.LimitPage = 0
	getLimit = paging.GetLimit()
	assert.Equal(10, getLimit, "when limit 0 return default value 10")
	assert.NotEqual(0, getLimit, "should be not equal")

	paging.LimitPage = -3
	getLimit = paging.GetLimit()
	assert.Equal(10, getLimit, "when limit negative return default value 10")
	assert.NotEqual(-3, getLimit, "should be not equal")
}

func TestPagination_GetOffset(t *testing.T) {
	assert := assert2.New(t)

	var getOffset = paging.GetOffset()
	assert.Equal(0, getOffset, "Limit should be 0")

	// change current page +
	paging.CurrentPage = 3
	getOffset = paging.GetOffset()
	assert.Equal((paging.CurrentPage-1)*paging.LimitPage, getOffset, "Offset should be equal")

	// change current page -
	paging.CurrentPage = -3
	getOffset = paging.GetOffset()
	assert.Equal((paging.CurrentPage-1)*paging.LimitPage, getOffset, "Offset should be equal")

	// change limit page
	paging.LimitPage = 10
	getOffset = paging.GetOffset()
	assert.Equal((paging.CurrentPage-1)*paging.LimitPage, getOffset, "Offset should be equal")
}
