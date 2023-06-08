package model

import "time"

type RandomUser struct {
	Results []struct {
		Gender string `json:"gender,omitempty"`
		Name   struct {
			Title string `json:"title,omitempty"`
			First string `json:"first,omitempty"`
			Last  string `json:"last,omitempty"`
		} `json:"name,omitempty"`
		Location struct {
			Street struct {
				Number int    `json:"number,omitempty"`
				Name   string `json:"name,omitempty"`
			} `json:"street,omitempty"`
			City        string `json:"city,omitempty"`
			State       string `json:"state,omitempty"`
			Country     string `json:"country,omitempty"`
			Postcode    int    `json:"postcode,omitempty"`
			Coordinates struct {
				Latitude  string `json:"latitude,omitempty"`
				Longitude string `json:"longitude,omitempty"`
			} `json:"coordinates,omitempty"`
			Timezone struct {
				Offset      string `json:"offset,omitempty"`
				Description string `json:"description,omitempty"`
			} `json:"timezone,omitempty"`
		} `json:"location,omitempty"`
		Email string `json:"email,omitempty"`
		Login struct {
			UUID     string `json:"uuid,omitempty"`
			Username string `json:"username,omitempty"`
			Password string `json:"password,omitempty"`
			Salt     string `json:"salt,omitempty"`
			Md5      string `json:"md5,omitempty"`
			Sha1     string `json:"sha1,omitempty"`
			Sha256   string `json:"sha256,omitempty"`
		} `json:"login,omitempty"`
		Dob struct {
			Date time.Time `json:"date,omitempty"`
			Age  int       `json:"age,omitempty"`
		} `json:"dob,omitempty"`
		Registered struct {
			Date time.Time `json:"date,omitempty"`
			Age  int       `json:"age,omitempty"`
		} `json:"registered,omitempty"`
		Phone string `json:"phone,omitempty"`
		Cell  string `json:"cell,omitempty"`
		ID    struct {
			Name  string `json:"name,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"id,omitempty"`
		Picture struct {
			Large     string `json:"large,omitempty"`
			Medium    string `json:"medium,omitempty"`
			Thumbnail string `json:"thumbnail,omitempty"`
		} `json:"picture,omitempty"`
		Nat string `json:"nat,omitempty"`
	} `json:"results,omitempty"`
	Info struct {
		Seed    string `json:"seed,omitempty"`
		Results int    `json:"results,omitempty"`
		Page    int    `json:"page,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"info,omitempty"`
}
