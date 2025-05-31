// Package models provides data models and database access functions for dashboard tables.
package models

import (
	"database/sql"
	"errors"
)

// Component represents a row in the components table.
type Component struct {
	ID    int
	Index int
	Name  string
}

// ComponentChart represents a row in the component_charts table.
type ComponentChart struct {
	ID         int
	ComponentID int
	Type       string
	Color      string
	Unit       string
}

// QueryChart represents a row in the query_charts table.
type QueryChart struct {
	ID           int
	SQL          string
	Description  string
	DataSource   string
	City         string
	UpdateFreq   string
}

// GetAllComponents returns all components from the database.
func GetAllComponents(db *sql.DB) ([]Component, error) {
	rows, err := db.Query("SELECT id, index, name FROM components")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []Component
	for rows.Next() {
		var c Component
		if err := rows.Scan(&c.ID, &c.Index, &c.Name); err != nil {
			return nil, err
		}
		components = append(components, c)
	}
	return components, rows.Err()
}

// GetAllComponentCharts returns all component_charts from the database.
func GetAllComponentCharts(db *sql.DB) ([]ComponentChart, error) {
	rows, err := db.Query("SELECT id, component_id, type, color, unit FROM component_charts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var charts []ComponentChart
	for rows.Next() {
		var cc ComponentChart
		if err := rows.Scan(&cc.ID, &cc.ComponentID, &cc.Type, &cc.Color, &cc.Unit); err != nil {
			return nil, err
		}
		charts = append(charts, cc)
	}
	return charts, rows.Err()
}

// GetAllQueryCharts returns all query_charts from the database.
func GetAllQueryCharts(db *sql.DB) ([]QueryChart, error) {
	rows, err := db.Query("SELECT id, sql, description, data_source, city, update_freq FROM query_charts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var queries []QueryChart
	for rows.Next() {
		var qc QueryChart
		if err := rows.Scan(&qc.ID, &qc.SQL, &qc.Description, &qc.DataSource, &qc.City, &qc.UpdateFreq); err != nil {
			return nil, err
		}
		queries = append(queries, qc)
	}
	return queries, rows.Err()
}

// GetComponentByID returns a component by its ID.
func GetComponentByID(db *sql.DB, id int) (*Component, error) {
	row := db.QueryRow("SELECT id, index, name FROM components WHERE id = ?", id)
	var c Component
	if err := row.Scan(&c.ID, &c.Index, &c.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}
