package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (img *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal Json value:", value))
	}

	var im_ge Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*img = im_ge
	return nil
}

func (img *Image) Value() (driver.Value, error) {
	if img == nil {
		return nil, nil
	}
	return json.Marshal(img)
}

type Images []Image

func (img *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal Json value", value))
	}

	var img_new []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}
	*img = img_new
	return nil
}

func (imgs *Images) Value() (driver.Value, error) {
	if imgs == nil {
		return nil, nil
	}
	return json.Marshal(imgs)
}
