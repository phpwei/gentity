package models

 type Books  struct {
    
        BookId int `gorm:"column:book_id;auto_increment" json:"book_id"`
	
        BookName string `gorm:"column:book_name;" json:"book_name"`
	
        BookIntr string `gorm:"column:book_intr;" json:"book_intr"`
	
        BookPrice1 string `gorm:"column:book_price1;" json:"book_price1"`
	
        BookPrice2 string `gorm:"column:book_price2;" json:"book_price2"`
	
        BookAuthor string `gorm:"column:book_author;" json:"book_author"`
	
        BookPress string `gorm:"column:book_press;" json:"book_press"`
	
        BookDate string `gorm:"column:book_date;" json:"book_date"`
	
        BookKindStr string `gorm:"column:book_kind_str;" json:"book_kind_str"`
	
        BookKind int `gorm:"column:book_kind;" json:"book_kind"`
	
        DdId int `gorm:"column:dd_id;" json:"dd_id"`
	
 }
 
  func(this *Books) WithBookId(BookId int) *Books  {
      this.BookId=BookId
      return this
  }
  
  func(this *Books) WithBookName(BookName string) *Books  {
      this.BookName=BookName
      return this
  }
  
  func(this *Books) WithBookIntr(BookIntr string) *Books  {
      this.BookIntr=BookIntr
      return this
  }
  
  func(this *Books) WithBookPrice1(BookPrice1 string) *Books  {
      this.BookPrice1=BookPrice1
      return this
  }
  
  func(this *Books) WithBookPrice2(BookPrice2 string) *Books  {
      this.BookPrice2=BookPrice2
      return this
  }
  
  func(this *Books) WithBookAuthor(BookAuthor string) *Books  {
      this.BookAuthor=BookAuthor
      return this
  }
  
  func(this *Books) WithBookPress(BookPress string) *Books  {
      this.BookPress=BookPress
      return this
  }
  
  func(this *Books) WithBookDate(BookDate string) *Books  {
      this.BookDate=BookDate
      return this
  }
  
  func(this *Books) WithBookKindStr(BookKindStr string) *Books  {
      this.BookKindStr=BookKindStr
      return this
  }
  
  func(this *Books) WithBookKind(BookKind int) *Books  {
      this.BookKind=BookKind
      return this
  }
  
  func(this *Books) WithDdId(DdId int) *Books  {
      this.DdId=DdId
      return this
  }
  

  
   type BooksOptions struct {
       apply func(*Books)
   }
   func NewBooks(opts...*BooksOptions)  *Books {
       m:= &Books{}
       for _,opt:=range opts{
           opt.apply(m)
       }
       return m
   }
    
   func BooksBookId(book_id int) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookId=book_id
       }}
   }
   
   func BooksBookName(book_name string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookName=book_name
       }}
   }
   
   func BooksBookIntr(book_intr string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookIntr=book_intr
       }}
   }
   
   func BooksBookPrice1(book_price1 string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookPrice1=book_price1
       }}
   }
   
   func BooksBookPrice2(book_price2 string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookPrice2=book_price2
       }}
   }
   
   func BooksBookAuthor(book_author string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookAuthor=book_author
       }}
   }
   
   func BooksBookPress(book_press string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookPress=book_press
       }}
   }
   
   func BooksBookDate(book_date string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookDate=book_date
       }}
   }
   
   func BooksBookKindStr(book_kind_str string) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookKindStr=book_kind_str
       }}
   }
   
   func BooksBookKind(book_kind int) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.BookKind=book_kind
       }}
   }
   
   func BooksDdId(dd_id int) *BooksOptions {
       return &BooksOptions{func(model *Books) {
           model.DdId=dd_id
       }}
   }
   









