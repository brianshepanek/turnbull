package repository

import (
	"context"
	"database/sql"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	"strings"
	"time"
)

type mysqlPostRepositoryStruct struct {
	db    *sql.DB
	table string
}
type post struct {
	Id       sql.NullInt64
	Title    sql.NullString
	Subtitle sql.NullString
	Views    sql.NullInt32
	Created  sql.NullTime
	Modified sql.NullTime
}

func (r *mysqlPostRepositoryStruct) Browse(ctx context.Context, req entity.Posts) error {

	var sqlStatement string
	sqlStatement += "SELECT id, title, subtitle, views, created, modified FROM " + r.table

	stmt, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var res post

		err := rows.Scan(&res.Id, &res.Title, &res.Subtitle, &res.Views, &res.Created, &res.Modified)
		if err != nil {
			return err
		}

		elem := entity.NewPost()

		if res.Id.Valid {
			value, err := res.Id.Value()
			if err != nil {
				return err
			}
			val := value.(int64)
			elem.SetId(&val)
		}

		if res.Title.Valid {
			value, err := res.Title.Value()
			if err != nil {
				return err
			}
			val := value.(string)
			elem.SetTitle(&val)
		}

		if res.Subtitle.Valid {
			value, err := res.Subtitle.Value()
			if err != nil {
				return err
			}
			val := value.(string)
			elem.SetSubtitle(&val)
		}

		if res.Views.Valid {
			value, err := res.Views.Value()
			if err != nil {
				return err
			}
			val := value.(int)
			elem.SetViews(&val)
		}

		if res.Created.Valid {
			value, err := res.Created.Value()
			if err != nil {
				return err
			}
			val := value.(time.Time)
			elem.SetCreated(&val)
		}

		if res.Modified.Valid {
			value, err := res.Modified.Value()
			if err != nil {
				return err
			}
			val := value.(time.Time)
			elem.SetModified(&val)
		}

		req.Append(elem)

	}

	return nil

}

func (r *mysqlPostRepositoryStruct) Read(ctx context.Context, id int64, req entity.Post) error {

	var sqlStatement string
	sqlStatement += "SELECT id, title, subtitle, views, created, modified FROM " + r.table + " WHERE id = ?"

	stmt, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err != nil {
		return err
	}

	var res post

	err = row.Scan(&res.Id, &res.Title, &res.Subtitle, &res.Views, &res.Created, &res.Modified)
	if err != nil {
		return err
	}

	if res.Id.Valid {
		var val int64
		err := res.Id.Scan(&val)
		if err != nil {
			return err
		}
		req.SetId(&val)
	}

	if res.Title.Valid {
		var val string
		err := res.Title.Scan(&val)
		if err != nil {
			return err
		}
		req.SetTitle(&val)
	}

	if res.Subtitle.Valid {
		var val string
		err := res.Subtitle.Scan(&val)
		if err != nil {
			return err
		}
		req.SetSubtitle(&val)
	}

	if res.Views.Valid {
		var val int
		err := res.Views.Scan(&val)
		if err != nil {
			return err
		}
		req.SetViews(&val)
	}

	if res.Created.Valid {
		var val time.Time
		err := res.Created.Scan(&val)
		if err != nil {
			return err
		}
		req.SetCreated(&val)
	}

	if res.Modified.Valid {
		var val time.Time
		err := res.Modified.Scan(&val)
		if err != nil {
			return err
		}
		req.SetModified(&val)
	}

	return nil

}

func (r *mysqlPostRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Post) error {

	var set []string
	var vals []interface{}
	if req.Id() != nil {
		set = append(set, "id = ?")
		vals = append(vals, req.Id())
	}

	if req.Title() != nil {
		set = append(set, "title = ?")
		vals = append(vals, req.Title())
	}

	if req.Subtitle() != nil {
		set = append(set, "subtitle = ?")
		vals = append(vals, req.Subtitle())
	}

	if req.Views() != nil {
		set = append(set, "views = ?")
		vals = append(vals, req.Views())
	}

	if req.Created() != nil {
		set = append(set, "created = ?")
		vals = append(vals, req.Created())
	}

	if req.Modified() != nil {
		set = append(set, "modified = ?")
		vals = append(vals, req.Modified())
	}

	vals = append(vals, id)

	var sqlStatement string
	sqlStatement += "UPDATE " + r.table + " "
	if len(set) > 0 {
		sqlStatement += "SET " + strings.Join(set, ", ") + " "
	}
	sqlStatement += " WHERE id = ?"

	stmt, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlPostRepositoryStruct) Add(ctx context.Context, req entity.Post) error {

	var set, vars []string
	var vals []interface{}
	if req.Id() != nil {
		set = append(set, "id")
		vars = append(vars, "?")
		vals = append(vals, req.Id())
	}

	if req.Title() != nil {
		set = append(set, "title")
		vars = append(vars, "?")
		vals = append(vals, req.Title())
	}

	if req.Subtitle() != nil {
		set = append(set, "subtitle")
		vars = append(vars, "?")
		vals = append(vals, req.Subtitle())
	}

	if req.Views() != nil {
		set = append(set, "views")
		vars = append(vars, "?")
		vals = append(vals, req.Views())
	}

	if req.Created() != nil {
		set = append(set, "created")
		vars = append(vars, "?")
		vals = append(vals, req.Created())
	}

	if req.Modified() != nil {
		set = append(set, "modified")
		vars = append(vars, "?")
		vals = append(vals, req.Modified())
	}

	var sqlStatement string
	sqlStatement += "INSERT INTO " + r.table + " "
	if len(set) > 0 {
		sqlStatement += "(" + strings.Join(set, ", ") + ") "
	}
	if len(vars) > 0 {
		sqlStatement += "VALUES(" + strings.Join(vars, ", ") + ") "
	}

	stmt, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}

	return nil

}

func (r *mysqlPostRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Post) error {
	var sqlStatement string
	sqlStatement += "DELETE FROM " + r.table + " WHERE id = ?"

	stmt, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil

}
