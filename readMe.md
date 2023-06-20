# Group in GO
- Ta có thể group các request lại với nhau nếu chúng có cùng 1 mục đích cho cùng 1 đối tượng ví dụ như group products vs các chức năng như của một restful API
- Từ đây ta có thể phân ra các tầng của request: ví dụ trong bài này tầng API là cao nhất xong dưới API sẽ có users và products, dưới các users và products là các detail request cho từng tầng

# Middleware in Go
- Ta có thể declare middleware trước các tầng của request. 
- Ví dụ trước khi vào tầng APi ta có thể declare 1 middleware ở đây xong sau khi vào tầng thứ 2 ta có thể declare tầng tiếp theo ở users và products. Thậm chí vô từng detail request ta cũng có thể thêm các middleware vào trong từng cái. VD: 

```
productRouter.GET("/:id", func(ctx *gin.Context) { // Get list pagination
  id := ctx.Param("id")
  ctx.JSON(http.StatusOK, gin.H{
    "message": "get by id: " + id,
  }, hasRole("admin"), middleware2(), middleware3(), ....) // middleware hasRole just appy for only this detail request
})
```

# Model in Go (G-ORM: object relational mapping)
- Lúc khởi tạo model ta cần tạo các giá trị được viết hoa chữ cái đầu
- gorm giúp cho việc tạo và lấy models dễ dàng hơn từ database nó cho phép map với models mình tạo ra
- Các keyword dùng trong gorm có thể refer theo [link này nè](https://gorm.io/docs/models.html#Fields-Tags).
- gorm hỗ trợ cho việc khởi tạo table trong database rất nhiều thông qua việc automigrate
- gorm không trả về dưới dạng các repository của từng bảng mà sẽ trả nguyên cục database khi tương tác ta  chỉ cần truyền vào dạng struct mà ta muốn làm thì nó sẽ tự hiểu ta đang tương tác với bảng cùa struct đó
- Mục đích của việc trả dữ liệu theo gin.H JSON thì để cho việc phân chia dạng dữ liệu của từng object, giúp ta có phân loại các object dễ dàng hơn khi được gửi đến front-end. Ví dụ:
```
userRouter.GET("/", func(ctx *gin.Context) {
  users := &[]models.User{}
  r.DB.Find(users)
  ctx.JSON(http.StatusOK, gin.H{"data": users, "total": 1000})
})
```
- Các công cụ hỗ trợ cho việc làm việc với database trong relationship được hướng dẫn trong mục Associations theo link [này](https://gorm.io/docs/). Như trong bài này ta sử dụng One To Many
VD: 1 User có thể có 1 hoặc nhiều products => One To Many

hoặc thường trong 1 thiết kế ta sẽ có struct là User 1 struct là Profile

```
table User {
  id 
  profileId = Profile.id  ///// Đây là dạng HasOne trong bảng cấu trúc dữ liệu tức là 1 User có 1 Profile có id như trên
}

table Profile {
  id 
  userId = User.id  //// Đây là dạng Belong To tức là Profile này thuộc về User có id như trên 
}
```

- fk trong database = Foreign Key

- Mỗi handler nên cần có 1 file handler cho từng repo riêng không thể tạo 1 file chung do từng loại repo sẽ có 1 cấu trúc dữ liệu khác nhau nên mặc dù file giống nhau nhưng cấu trúc dữ liệu trả về là khác nhau

- Khi tạo object để mapping json từ request lên server thì nếu đã tạo field của struct với dạng data gì thì ở dạng data đó đã có validation cho việc mapping nên dường như không cần phải validation gì nhiều cho vấn đề đó và ông front-end cần đảm bảo tuân theo quy tắc đó khi design hệ thống