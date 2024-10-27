package main

import (
	"database/sql"
	"fmt"
)

type ObjectType struct {
	ID               int
	Name             string
	Color            string
	Monogram         string
	ObjectTypeGroupID sql.NullInt64 // To handle optional ObjectTypeGroupID
	WorkflowID       int
	CreatedAt        string
	UpdatedAt        string
}


func CreateObjectType(db *sql.DB, objType ObjectType) error {
	sqlStatement := `
	INSERT INTO ObjectTypes (name, color, monogram, object_type_group_id, workflow_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, now(), now()) RETURNING id`

	err := db.QueryRow(sqlStatement, objType.Name, objType.Color, objType.Monogram, objType.ObjectTypeGroupID, objType.WorkflowID).Scan(&objType.ID)
	if err != nil {
		return err
	}

	fmt.Printf("New ObjectType ID: %d\n", objType.ID)
	return nil
}

func GetObjectType(db *sql.DB, id int) (ObjectType, error) {
	var objType ObjectType

	sqlStatement := `SELECT id, name, color, monogram, object_type_group_id, workflow_id, created_at, updated_at
	                 FROM ObjectTypes WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&objType.ID, &objType.Name, &objType.Color, &objType.Monogram, &objType.ObjectTypeGroupID, &objType.WorkflowID, &objType.CreatedAt, &objType.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return objType, fmt.Errorf("no ObjectType found with id %d", id)
		}
		return objType, err
	}

	return objType, nil
}

func UpdateObjectType(db *sql.DB, objType ObjectType) error {
	sqlStatement := `
	UPDATE ObjectTypes
	SET name = $2, color = $3, monogram = $4, object_type_group_id = $5, workflow_id = $6, updated_at = now()
	WHERE id = $1`

	_, err := db.Exec(sqlStatement, objType.ID, objType.Name, objType.Color, objType.Monogram, objType.ObjectTypeGroupID, objType.WorkflowID)
	if err != nil {
		return err
	}

	fmt.Printf("ObjectType with ID %d updated\n", objType.ID)
	return nil
}


func DeleteObjectType(db *sql.DB, id int) error {
	sqlStatement := `DELETE FROM ObjectTypes WHERE id = $1`

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	fmt.Printf("ObjectType with ID %d deleted\n", id)
	return nil
}

