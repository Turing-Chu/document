# 在 Go 中使用 go-sqlmock 对 GORM 进行单元测试

>[Rosaniline](https://medium.com/@rosaniline?source=post_page-----93cbce1f6b5b--------------------------------) At 2019-03-31
> 
> Translated by Turing Zhu
> 
> Original Article: [Unit testing GORM with go-sqlmock in Go](https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b)

当开始写 go 的时候，我开始想念在 python 中使用 `MagicMock` 的时光。然而在 go 中写测试真的没有那么难。

我我们来讨论下如何用 GORM 测试 go 中的数据库交互。

...

## 先决条件

先以简单的模型 `Persion` 为例。

### 模型
```go
type Persion struct {
	ID uuid.UUID `gorm:"column:id;primary_key" json:"id"`
	Name string  `gorm:"column:name" json:"name"`
}
```

`Repository `

Repository 作为给定模型的包装数据访问层，具有 `GET` 和 `CREATE` 两个函数。

```go
type Repository interface {
   Get(id uuid.UUID) (*model.Person, error)
   Create(id uuid.UUID, name string) error
}

func (p *repo) Create(id uuid.UUID, name string) error {
   person := &model.Person{
      ID:   id,
      Name: name,
   }

   return p.DB.Create(person).Error
}

func (p *repo) Get(id uuid.UUID) (*model.Person, error) {
   person := new(model.Person)

   err := p.DB.Where("id = ?", id).Find(person).Error

   return person, err
}
```

我们的目的是测试在 `Repository ` 中实现的函数，以确保在 GORM 之下发生的事情与我们所期望的一致。

...

## 测试设置

在钻研测试如何实现之前。需要先了解几个组件。

- 来自 `testify` 的 `suite`
- 来自 `DATA-DOG` 的 `sql-mock`

### Suite

我们使用 [testify](https://github.com/stretchr/testify) 的 `suite` 来简化测试设置。如果还不熟悉 `suite` 的话，看一下下面来自 [testify](https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b#:~:text=quote%20from%20the-,testify,-below.) 的引用。

> `suite` 提供给你可能从更常见的面向对象语言中使用的功能。你可以用它构建一个测试单元作为结构体，在结构体上构建 setup/teardown 方法和测试方法，以及是用 ‘go test’ 正常运行它们。

下面是 `suite` 的写法。

```
type Suite struct {
   suite.Suite
   DB   *gorm.DB
   mock sqlmock.Sqlmock

   repository Repository
   person     *model.Person
}
```

### sql-mock

这可能是今天的主要主题。我们也有来自 `DATA-DOG` 的引用来说明 sql-mock 是什么。

> `sqlmock` 是一个 实现 [sql/driver](https://godoc.org/database/sql/driver) 的 mock 库。其有且只有一个目的--在测试中模拟任何 sql 驱动器行为，而无需真正的数据库连接。它帮助维护正确的 TDD 工作流。

...

## 测试
最后这是我们的主题。我们来看看测试是如何写的，以及一步一步的测试 GORM 操作。

- 设置 `suite`
- 使用 `sql-mock` 设置一些 sql 语句 的 `Expects`
- 调用函数来测试
- 断言函数返回是正确的
- 检查 `sql-mock` 的`预期`**是否**满足

### 设置套件
在这个阶段我们会准备好我们的模拟数据库与存储库。这与面向对象的设置过程极其相似但是是用 `sql-mock` 作为 sql 驱动。

```go
func (s *Suite) SetupSuite() {
   var (
      db  *sql.DB
      err error
   )

   db, s.mock, err = sqlmock.New()
   require.NoError(s.T(), err)

   s.DB, err = gorm.Open("postgres", db)
   require.NoError(s.T(), err)

   s.DB.LogMode(true)

   s.repository = CreateRepository(s.DB)
}
```

### 测试 SELECT 语句

还记得在 `Repository` 中我们有一个 `GET` 方法么？为了在 `persion`中根据给定的 id 检索一行。我们来看下如何测试它。

```go
func (s *Suite) Test_repository_Get() {
   var (
      id   = uuid.NewV4()
      name = "test-name"
   )

   s.mock.ExpectQuery(regexp.QuoteMeta(
      `SELECT * FROM "person" WHERE (id = $1)`)).
      WithArgs(id.String()).
      WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
         AddRow(id.String(), name))

   res, err := s.repository.Get(id)

   require.NoError(s.T(), err)
   require.Nil(s.T(), deep.Equal(&model.Person{ID: id, Name: name}, res))
}
``` 

这里我们利用 `sql-mock` 来为我们做这些事情

- 期望执行 `SELECT * FROM "person" WHERE (id = $1)`
- 带上 `id` 参数
- 返回 `id` 和 `name` 做为 persion 记录的存根


### 测试 INSERT 语句

在 `Repository ` 除了 `GET` 还有另一个 `CREATE` 函数。

```go
func (s *Suite) Test_repository_Create() {
   var (
      id   = uuid.NewV4()
      name = "test-name"
   )

   s.mock.ExpectQuery(regexp.QuoteMeta(
      `INSERT INTO "person" ("id","name") 
       VALUES ($1,$2) RETURNING "person"."id"`)).
      WithArgs(id, name).
      WillReturnRows(
         sqlmock.NewRows([]string{"id"}).AddRow(id.String()))

   err := s.repository.Create(id, name)

   require.NoError(s.T(), err)
}
```

这里我们利用 `sql` `sql-mock` 来为我们做这些事情

- 期望执行 `INSERT` 语句
- 带上 `id` 和`name` 参数
- 返回创建行的 `id`

### 检查 `sql-mock` 的`预期`**是否**满足

该检查是放在 `AfterTest` 部分以确保在每个测试用例之后执行。

```
func (s *Suite) AfterTest(_, _ string) {
   require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
``` 

对于上面提到的代码，如果你发现分开的片段难以阅读。代码仓库在这里 [repository](https://github.com/Rosaniline/gorm-ut) 。

在 Go 中测试 GORM 并不难，对吧？编码愉快 🐤

