package helper

import (
	"time"

	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/config"
)

type  InventoryServiceHelper struct {
	cfg config.Config
}


func NewHelper(config config.Config) *InventoryServiceHelper {
	return &InventoryServiceHelper{
		cfg: config,
	}
}
func (h *InventoryServiceHelper)StringToTime(timeStr string) (time.Time, error) {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func (h *InventoryServiceHelper) TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}



