// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-web/lesson/chapter6_1/repository/ent/itemcategory"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ItemCategory is the model entity for the ItemCategory schema.
type ItemCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemCategoryQuery when eager-loading is set.
	Edges ItemCategoryEdges `json:"edges"`
}

// ItemCategoryEdges holds the relations/edges for other nodes in the graph.
type ItemCategoryEdges struct {
	// Items holds the value of the items edge.
	Items []*Item `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e ItemCategoryEdges) ItemsOrErr() ([]*Item, error) {
	if e.loadedTypes[0] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ItemCategory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case itemcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case itemcategory.FieldName:
			values[i] = new(sql.NullString)
		case itemcategory.FieldCreateTime, itemcategory.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ItemCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ItemCategory fields.
func (ic *ItemCategory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case itemcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ic.ID = int(value.Int64)
		case itemcategory.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ic.CreateTime = value.Time
			}
		case itemcategory.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ic.UpdateTime = value.Time
			}
		case itemcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ic.Name = value.String
			}
		}
	}
	return nil
}

// QueryItems queries the "items" edge of the ItemCategory entity.
func (ic *ItemCategory) QueryItems() *ItemQuery {
	return (&ItemCategoryClient{config: ic.config}).QueryItems(ic)
}

// Update returns a builder for updating this ItemCategory.
// Note that you need to call ItemCategory.Unwrap() before calling this method if this ItemCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ic *ItemCategory) Update() *ItemCategoryUpdateOne {
	return (&ItemCategoryClient{config: ic.config}).UpdateOne(ic)
}

// Unwrap unwraps the ItemCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ic *ItemCategory) Unwrap() *ItemCategory {
	tx, ok := ic.config.driver.(*txDriver)
	if !ok {
		panic("ent: ItemCategory is not a transactional entity")
	}
	ic.config.driver = tx.drv
	return ic
}

// String implements the fmt.Stringer.
func (ic *ItemCategory) String() string {
	var builder strings.Builder
	builder.WriteString("ItemCategory(")
	builder.WriteString(fmt.Sprintf("id=%v", ic.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ic.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ic.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(ic.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ItemCategories is a parsable slice of ItemCategory.
type ItemCategories []*ItemCategory

func (ic ItemCategories) config(cfg config) {
	for _i := range ic {
		ic[_i].config = cfg
	}
}