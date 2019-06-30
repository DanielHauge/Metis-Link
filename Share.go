package main

import "time"

type Share struct {
	Type string `json:"type"`
	Date time.Time `json:"date"`
	Ip string `json:"ip"`
	element interface{} `json:"element"`
}


func CreateShare (t string, d time.Time, ip string, elem interface{}) *Share{
	return &Share{Type:t, Date:d, element:elem, Ip:ip}
}

func (s *Share) toJson()([]byte, error){
	return json.Marshal(s)
}