package world

import "encoding/json"

type unitState struct {
	UnitType string `json:"unit_type"`
	XPos     int    `json:"x_pos"`
	ZPos     int    `json:"z_pos"`
}

type BoardState struct {
	XDim  int          `json:"x_dim"`
	ZDim  int          `json:"z_dim"`
	Units []*unitState `json:"units"`
}

func GetDefaultBoardState() string {
	data, err := json.Marshal(defaultBoard.getState())
	if err != nil {
		return ""
	}
	return string(data)
}
