package models

type User struct {
  Id int `json:"id"`
  Givey_Tag string `json:"givey_tag"`
  Name string `json:"name"`
  First_Name string `json:"first_name"`
  Last_Name string `json:"last_name"`
  Profile_Url string `json:"profileurl"`
}
