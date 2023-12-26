package do

type SortDO struct {
	Id       *uint     `json:"id"`
	Sort     *uint     `json:"sort"`
	UserAble bool      `json:"user_able"`
	Title    string    `json:"title"`
	Desc     *string   `json:"desc"`
	Blogs    *[]BlogDO `json:"blogs"`
}
