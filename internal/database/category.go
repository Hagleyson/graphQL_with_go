package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct{
	db *sql.DB
	ID string
	Name string
	Description string
}

func NewCategory(db *sql.DB)*Category{
	return &Category{db: db}
}

func (c *Category ) Create(name string,description string)(Category, error){
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2,$3)",id,name,description)
	if(err != nil){
		return Category{},err
	}
	return Category{ID:id, Name:name, Description: description},nil
}

func (c *Category ) FindAll()([]Category, error){	
	rows, err := c.db.Query("SELECT * FROM categories")

	if(err != nil){
		return nil,err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next(){
		var id, name,description string
		if err := rows.Scan(&id,&name, &description);err != nil{
			return nil,err
		}
		categories = append(categories, Category{ID:id, Name:name, Description: description})
	}

	return categories,nil
}

func (c *Course) FindByCategoryID(categoryID string)([]Course, error)  {
	rows, err:= c.db.Query("SELECT id,name,description, category_id FROM courses WHERE category_id = $1",categoryID)
	if err != nil{
		return nil, err
	}

	defer rows.Close()
	courses := []Course{}
	for rows.Next(){
		var id, name, description, categoryID string
		if err:= rows.Scan(&id, &name, &description, &categoryID); err != nil{
			return nil, err
		}
		courses = append(courses,Course{ID: id,Name: name,Description: description,CategoryID: categoryID})
	}
	return courses, nil
	
}