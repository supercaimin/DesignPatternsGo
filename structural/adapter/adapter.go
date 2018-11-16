package adapter

type Booker interface {
	TurnPage()
	Open()
	GetPage() int
}

type Book struct {
	page int
}

func (self *Book)Open()  {
	self.page = 1
}
func (self *Book)TurnPage()  {
	self.page++
}
func (self *Book)GetPage() int {
	return self.page
}

type EBooker interface {
	Unlock()
	PressNext()
	GetPage() (int,int)
}

type Kindle struct {
	page int
	totalPages int
}

func NewKindle() *Kindle{
	return &Kindle{1, 100}
}

func (self *Kindle)Unlock(){

}

func (self *Kindle)PressNext(){
	self.page++
}

func (self *Kindle)GetPage()(int, int){
	return self.page, self.totalPages
}

type EBookAdapter struct {
	ebook EBooker
}

func NewEBookAdapter(ebook EBooker) *EBookAdapter{
	return &EBookAdapter{ebook}
}


func (self *EBookAdapter)Open()  {
	self.ebook.Unlock()
}
func (self *EBookAdapter)TurnPage()  {
	self.ebook.PressNext()
}
func (self *EBookAdapter)GetPage() int {
	page,_ := self.ebook.GetPage()
	return page
}