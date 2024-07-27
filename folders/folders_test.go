package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

type mockTestData struct {
	name string
	orgId uuid.UUID
	mockSampleData []*folders.Folder
	expected []*folders.Folder
}

var testData = []mockTestData {
	{
		name: "Should return empty list when sample data is empty",
		orgId: uuid.FromStringOrNil("99999999-9999-9999-9999-999999999999"),
		mockSampleData: []*folders.Folder{},
		expected: []*folders.Folder{},
	},
	{
		name: "Should return empty list when query matches no items",
		orgId: uuid.FromStringOrNil("99999999-9999-9999-9999-999999999999"),
		mockSampleData: []*folders.Folder{
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
		},
		expected: []*folders.Folder{},
	},
	{
		name: "Should return list of 1 item when query matches single item",
		orgId: uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
		mockSampleData: []*folders.Folder{
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
		},
		expected: []*folders.Folder{
			{
				Id:      uuid.FromStringOrNil("00003d65-d336-485a-8331-7b53f37e8f55"),
				Name:    "Folder 3",
				OrgId:   uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333"),
				Deleted: false,
			},
		},
	},
	{
		name: "Should return list of multiple folders when query matches multiple items",
		orgId: uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222"),
		mockSampleData: []*folders.Folder{
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
		},
		expected: []*folders.Folder{
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
		},
	},
}

func Test_GetAllFoldersByOrgId(t *testing.T) {

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			// Mock fetchData
			folders.FetchData = func() []*folders.Folder {
				return tt.mockSampleData
			}

			req := &folders.FetchFolderRequest{OrgID: tt.orgId}
			resp, _ := folders.GetAllFoldersByOrgId(req)
			assert.Equal(t, tt.expected, resp.Folders)
		})
	}
}
