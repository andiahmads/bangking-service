package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

// return dto.CustomerResponse{
// 	Id:          c.Id,
// 	Name:        c.Name,
// 	City:        c.City,
// 	Zipcode:     c.Zipcode,
// 	DateOfBirth: c.DateOfBirth,
// 	Status:      c.Status,
// }
