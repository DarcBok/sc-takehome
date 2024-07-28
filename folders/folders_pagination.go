package folders

import (
	"encoding/base64"
	"errors"
	"strconv"

	"github.com/gofrs/uuid"
)

func GetAllFoldersByOrgIdWithPagination(req *FetchFolderWithPaginationRequest) (*FetchFolderWithPaginationResponse, error) {
	var (
		startIdx int
		err error
	)
	
	if (req.Token == "") {
		startIdx = 0
	} else {
		startIdx, err = decodeTokenToIdx(req.Token)
	}
	if err != nil {
		return nil, err
	}

	folders, nextIdx, err := FetchAllFoldersByOrgIDWithPagination(req.OrgId, startIdx, req.PageSize)
	if err != nil {
		return nil, err
	}

	var responseToken string
	if nextIdx == -1 {
		responseToken = ""
	} else {
		responseToken = encodeIdxToToken(nextIdx)
	}
	var response = &FetchFolderWithPaginationResponse{Folders: folders, Token: responseToken}
	return response, nil
}

// Returns the list of folders, and the next index to begin searching from.
// Returns next index of -1 if end of list reached.
func FetchAllFoldersByOrgIDWithPagination(orgID uuid.UUID, startIdx int, pageSize int) ([]*Folder, int, error) {
	folders := FetchData()

	if startIdx > len(folders) || startIdx < 0 {
		return nil, 0, errors.New("startIdx out of bounds")
	}

	resFolder := []*Folder{}
	numMatched := 0
	nextIdx := -1
	for i, folder := range folders[startIdx:] {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
			numMatched++
		}
		if numMatched == pageSize {
			nextIdx = startIdx + i + 1
			break
		}
	}
	return resFolder, nextIdx, nil
}

func encodeIdxToToken(index int) string {
	indexStr := []byte(strconv.Itoa(index))
	return base64.URLEncoding.EncodeToString([]byte(indexStr))
}

func decodeTokenToIdx(token string) (int, error) {
	decodedBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(decodedBytes))
}