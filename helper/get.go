package helper

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/makves-test-task/api/models"
	"github.com/spf13/cast"
)

func (f *FileManager) GetAllInfo(req *models.GetAllFileInfoReq) (*models.GetAllFileInfoRes, error) {
	defer f.ReOpenFile()
	reader := csv.NewReader(f.File)
	// Skip the first two lines
	keys := make(map[int]string)
	k, err := reader.Read()
	if err == io.EOF {
		return &models.GetAllFileInfoRes{}, nil
	}
	if err != nil {
		return &models.GetAllFileInfoRes{}, err
	}
	for i, e := range k {
		keys[i] = e
	}
	offset := (int(req.Page)-1)*int(req.Limit) + 1
	response := &models.GetAllFileInfoRes{}
	for i := 1; i < offset; i++ {
		_, err := reader.Read()
		if err == io.EOF {
			return &models.GetAllFileInfoRes{}, nil
		}
		if err != nil {
			return &models.GetAllFileInfoRes{}, err
		}

	}

	// Read the remaining lines
	for i := offset; i < offset+int(req.Limit); i++ {
		temp := make(map[string]interface{})
		record, err := reader.Read()
		if err == io.EOF {
			return &models.GetAllFileInfoRes{}, nil
		}
		if err != nil {
			return &models.GetAllFileInfoRes{}, err
		}
		for j, e := range record {
			temp[keys[j]] = e
		}
		jdata, err := json.Marshal(temp)
		if err != nil {
			return &models.GetAllFileInfoRes{}, err
		}
		response.Data = append(response.Data, string(jdata))
	}
	response.Count = int64(len(response.Data))
	return response, nil
}

func (f *FileManager) GetById(req *models.GetByidReq) (*models.GetAllFileInfoRes, error) {
	defer f.ReOpenFile()
	reader := csv.NewReader(f.File)
	// Skip the first two lines
	keys := make(map[int]string)
	k, err := reader.Read()
	if err == io.EOF {
		return &models.GetAllFileInfoRes{}, nil
	}
	if err != nil {
		return &models.GetAllFileInfoRes{}, err
	}
	for i, e := range k {
		keys[i] = e
	}
	response := &models.GetAllFileInfoRes{}
	// Read the remaining lines
	for {
		temp := make(map[string]interface{})
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return &models.GetAllFileInfoRes{}, err
		}
		for j, e := range record {
			temp[keys[j]] = e
		}
		if v, ok := temp["id"]; ok {
			
			if cast.ToInt(v) != req.Id {
				continue
			}
		}
		jdata, err := json.Marshal(temp)
		if err != nil {
			return &models.GetAllFileInfoRes{}, err
		}
		response.Data = append(response.Data, string(jdata))
	}
	response.Count = int64(len(response.Data))
	return response, nil
}
