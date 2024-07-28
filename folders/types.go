package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type FetchFolderWithPaginationRequest struct {
	// The OrgId to match to.
	OrgId uuid.UUID
	// The size of the page.
	PageSize int
	// The token passed from the previous page, or an empty string if first request.
	Token string
}

type FetchFolderWithPaginationResponse struct {
	// List of folders fetched.
	Folders []*Folder
	// Can be used in the next request to query the next set of data.
	// Empty string indicates end of data.
	Token string
}