package repository

import (
	"context"
	"database/sql"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	"strings"
	"time"
)

type mysqlCommentRepositoryStruct struct {
	db    *sql.DB
	table string
}
type comment struct {
	Id       sql.NullInt64
	PostId   sql.NullInt64
	Title    sql.NullString
	Body     sql.NullString
	Created  sql.NullTime
	Modified sql.NullTime
}

func (r *mysqlCommentRepositoryStruct) Browse(ctx context.Context, req entity.Comments) error {

	var sqlStatement string
	sqlStatement += "SELECT id, post_id, title, body, created, modified FROM " + r.table

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

		var res comment

		err := rows.Scan(&res.Id, &res.PostId, &res.Title, &res.Body, &res.Created, &res.Modified)
		if err != nil {
			return err
		}

		elem := entity.NewComment()

		if res.Id.Valid {
			var val int64
			err := res.Id.Scan(&val)
			if err != nil {
				return err
			}
			elem.SetId(&val)
		}

		if res.PostId.Valid {
			var val int64
			err := res.PostId.Scan(&val)
			if err != nil {
				return err
			}
			elem.SetPostId(&val)
		}

		if res.Title.Valid {
			var val string
			err := res.Title.Scan(&val)
			if err != nil {
				return err
			}
			elem.SetTitle(&val)
		}

		if res.Body.Valid {
			var val string
			err := res.Body.Scan(&val)
			if err != nil {
				return err
			}
			elem.SetBody(&val)
		}

		if res.Created.Valid {
			var val time.Time
			err := res.Created.Scan(&val)
			if err != nil {
				return err
			}
			elem.SetCreated(&val)
		}

		if res.Modified.Valid {
			var val time.Time
			err := res.Modified.Scan(&val)
			if err != nil {
				return err
			}
			elem.SetModified(&val)
		}

		req.Append(elem)

	}

	return nil

}

func (r *mysqlCommentRepositoryStruct) Read(ctx context.Context, id int64, req entity.Comment) error {

	var sqlStatement string
	sqlStatement += "SELECT id, post_id, title, body, created, modified FROM " + r.table + " WHERE id = ?"

	stmt, err := r.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err != nil {
		return err
	}

	var res comment

	err = row.Scan(&res.Id, &res.PostId, &res.Title, &res.Body, &res.Created, &res.Modified)
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

	if res.PostId.Valid {
		var val int64
		err := res.PostId.Scan(&val)
		if err != nil {
			return err
		}
		req.SetPostId(&val)
	}

	if res.Title.Valid {
		var val string
		err := res.Title.Scan(&val)
		if err != nil {
			return err
		}
		req.SetTitle(&val)
	}

	if res.Body.Valid {
		var val string
		err := res.Body.Scan(&val)
		if err != nil {
			return err
		}
		req.SetBody(&val)
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

func (r *mysqlCommentRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Comment) error {

	var set []string
	var vals []interface{}
	if req.Id() != nil {
		set = append(set, "id = ?")
		vals = append(vals, req.Id())
	}

	if req.PostId() != nil {
		set = append(set, "post_id = ?")
		vals = append(vals, req.PostId())
	}

	if req.Title() != nil {
		set = append(set, "title = ?")
		vals = append(vals, req.Title())
	}

	if req.Body() != nil {
		set = append(set, "body = ?")
		vals = append(vals, req.Body())
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

func (r *mysqlCommentRepositoryStruct) Add(ctx context.Context, req entity.Comment) error {

	var set, vars []string
	var vals []interface{}
	if req.Id() != nil {
		set = append(set, "id")
		vars = append(vars, "?")
		vals = append(vals, req.Id())
	}

	if req.PostId() != nil {
		set = append(set, "post_id")
		vars = append(vars, "?")
		vals = append(vals, req.PostId())
	}

	if req.Title() != nil {
		set = append(set, "title")
		vars = append(vars, "?")
		vals = append(vals, req.Title())
	}

	if req.Body() != nil {
		set = append(set, "body")
		vars = append(vars, "?")
		vals = append(vals, req.Body())
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

func (r *mysqlCommentRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Comment) error {
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