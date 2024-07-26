package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFoldersByOrgId(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

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
