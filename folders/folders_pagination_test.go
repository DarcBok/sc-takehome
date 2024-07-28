package folders_test

import (
	"encoding/base64"
	"strconv"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

var mockSampleData = []*folders.Folder{
	{
		Id:      uuid.FromStringOrNil("00001d65-d336-485a-8331-7b53f37e8f51"),
		Name:    "Folder 1",
		OrgId:   uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
		Deleted: false,
	},
	{
		Id:      uuid.FromStringOrNil("00002d65-d336-485a-8331-7b53f37e8f53"),
		Name:    "Folder 2",
		OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		Deleted: true,
	},
	{
		Id:      uuid.FromStringOrNil("00003d65-d336-485a-8331-7b53f37e8f55"),
		Name:    "Folder 3",
		OrgId:   uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
		Deleted: false,
	},
	{
		Id:      uuid.FromStringOrNil("00004d65-d336-485a-8331-7b53f37e8f57"),
		Name:    "Folder 4",
		OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		Deleted: true,
	},
	{
		Id:      uuid.FromStringOrNil("00005d65-d336-485a-8331-7b53f37e8f59"),
		Name:    "Folder 5",
		OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		Deleted: true,
	},
	{
		Id:      uuid.FromStringOrNil("00006d65-d336-485a-8331-7b53f37e8f61"),
		Name:    "Folder 6",
		OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		Deleted: true,
	},
	{
		Id:      uuid.FromStringOrNil("00007d65-d336-485a-8331-7b53f37e8f63"),
		Name:    "Folder 7",
		OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		Deleted: true,
	},
	{
		Id:      uuid.FromStringOrNil("00008d65-d336-485a-8331-7b53f37e8f65"),
		Name:    "Folder 8",
		OrgId:   uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
		Deleted: false,
	},
}

var mockSampleDataEmpty = []*folders.Folder{}

func mockSampleDataQuery(mockData []*folders.Folder) {
	folders.FetchData = func() []*folders.Folder {
		return mockData
	}
}

func Test_GetAllFoldersByOrgIdWithPagination(t *testing.T) {
	t.Run("Should return error if token decodes to non-integer", func(t *testing.T) {
		mockSampleDataQuery(mockSampleData)

		indexStr := []byte("abc")
		req := &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
			Token: base64.URLEncoding.EncodeToString(indexStr),
			PageSize: 2,
		}
		_, err := folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.NotNil(t, err)
	})

	t.Run("Should return error if token decodes to index negative", func(t *testing.T) {
		mockSampleDataQuery(mockSampleData)

		indexStr := []byte(strconv.Itoa(-1))
		req := &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
			Token: base64.URLEncoding.EncodeToString(indexStr),
			PageSize: 2,
		}
		_, err := folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.NotNil(t, err)
	})
	
	t.Run("Should return empty list and token when sample data is empty", func(t *testing.T) {
		mockSampleDataQuery(mockSampleDataEmpty)

		req := &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
			Token: "",
			PageSize: 2,
		}
		resp, _ := folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{}, resp.Folders)
		assert.Equal(t, "", resp.Token)
	})

	t.Run("Should return empty list and token when query doesn't match any data orgIds", func(t *testing.T) {
		mockSampleDataQuery(mockSampleData)

		req := &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("99999999-9999-9999-9999-999999999999"),
			Token: "",
			PageSize: 2,
		}
		resp, _ := folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{}, resp.Folders)
		assert.Equal(t, "", resp.Token)
	})

	t.Run("Should return list of 1 match and empty token when query matches 1 data folder and page size is 2", func(t *testing.T) {
		mockSampleDataQuery(mockSampleData)

		req := &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
			Token: "",
			PageSize: 2,
		}
		resp, _ := folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{
			{
				Id:      uuid.FromStringOrNil("00001d65-d336-485a-8331-7b53f37e8f51"),
				Name:    "Folder 1",
				OrgId:   uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"),
				Deleted: false,
			},
		}, resp.Folders)
		assert.Equal(t, "", resp.Token)
	})

	t.Run("Should be able to use token return to fetch more page results", func(t *testing.T) {
		mockSampleDataQuery(mockSampleData)

		req := &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
			Token: "",
			PageSize: 2,
		}
		resp, _ := folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{
			{
				Id:      uuid.FromStringOrNil("00002d65-d336-485a-8331-7b53f37e8f53"),
				Name:    "Folder 2",
				OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
				Deleted: true,
			},
			{
				Id:      uuid.FromStringOrNil("00004d65-d336-485a-8331-7b53f37e8f57"),
				Name:    "Folder 4",
				OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
				Deleted: true,
			},
		}, resp.Folders)
		assert.NotEmpty(t, resp.Token)

		req = &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
			Token: resp.Token,
			PageSize: 2,
		}
		resp, _ = folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{
			{
				Id:      uuid.FromStringOrNil("00005d65-d336-485a-8331-7b53f37e8f59"),
				Name:    "Folder 5",
				OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
				Deleted: true,
			},
			{
				Id:      uuid.FromStringOrNil("00006d65-d336-485a-8331-7b53f37e8f61"),
				Name:    "Folder 6",
				OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
				Deleted: true,
			},
		}, resp.Folders)
		assert.NotEmpty(t, resp.Token)

		req = &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
			Token: resp.Token,
			PageSize: 2,
		}
		resp, _ = folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{
			{
				Id:      uuid.FromStringOrNil("00007d65-d336-485a-8331-7b53f37e8f63"),
				Name:    "Folder 7",
				OrgId:   uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
				Deleted: true,
			},
		}, resp.Folders)
		assert.Equal(t, "", resp.Token)
	})

	t.Run("Should return empty token if end of list reached after 1 extra query", func(t *testing.T) {
		mockSampleDataQuery(mockSampleData)

		req := &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
			Token: "",
			PageSize: 2,
		}
		resp, _ := folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{
			{
				Id:      uuid.FromStringOrNil("00003d65-d336-485a-8331-7b53f37e8f55"),
				Name:    "Folder 3",
				OrgId:   uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
				Deleted: false,
			},
			{
				Id:      uuid.FromStringOrNil("00008d65-d336-485a-8331-7b53f37e8f65"),
				Name:    "Folder 8",
				OrgId:   uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
				Deleted: false,
			},
		}, resp.Folders)
		assert.NotEmpty(t, resp.Token)

		req = &folders.FetchFolderWithPaginationRequest{
			OrgId: uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
			Token: resp.Token,
			PageSize: 2,
		}
		resp, _ = folders.GetAllFoldersByOrgIdWithPagination(req)
		assert.Equal(t, []*folders.Folder{}, resp.Folders)
		assert.Equal(t, "", resp.Token)
	})

}
