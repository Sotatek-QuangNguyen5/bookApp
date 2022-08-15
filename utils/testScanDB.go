package utils


// func (b DefaultBookRepository) List() ([]*models.Book, *errs.AppError) {

// 	res, err := b.db.Query("SELECT * FROM book")

// 	if err != nil {

// 		return nil, errs.ErrorGetData()
// 	}

// 	cols, error := res.Columns()

// 	if error != nil {

// 		return nil, errs.ErrorGetData()
// 	}

// 	count := len(cols)
// 	values := make([]interface{}, count)
// 	valuePtr := make([]interface{}, count)
// 	fmt.Println(cols)
// 	var books []*models.Book

// 	for res.Next() {

// 		//var book = new(models.Book)
// 		//err := res.Scan(&book.Book_id, &book.Name, &book.Description, &book.Authors)
		
// 		for index, _ := range cols {

// 			valuePtr[index] = &values[index]
// 		}
// 		//fmt.Println(valuePtr)
// 		err := res.Scan(valuePtr...)
// 		if err != nil {

// 			return nil, errs.ErrorReadData()
// 		}
// 		//fmt.Println(values)
// 		for index, col := range cols {

// 			var v interface{}
// 			val := values[index]
// 			fmt.Println(val)
// 			b, ok := val.([]byte)
// 			if ok {

// 				v = string(b)
// 			} else {

// 				v = val
// 				fmt.Println(v, "    11111")
// 			}
// 			fmt.Printf("%s,   Type of v : %T", col, v)
// 		}

// 		//fmt.Println(valuePtr)
// 		//books = append(books, book)
// 	}

// 	return books, nil

// }
