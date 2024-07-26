package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	var (
		// err error
		// f1  Folder
		// fs  []*Folder
	)
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: r}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
