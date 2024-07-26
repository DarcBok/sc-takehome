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
	folders, _ := FetchAllFoldersByOrgID(req.OrgID)
	var response = &FetchFolderResponse{Folders: folders}
	return response, nil
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
