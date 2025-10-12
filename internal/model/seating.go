package model

import "database/sql"

type (
	Floor struct {
		ID        int           `json:"id"`
		Name      string        `json:"name"`
		CreatedAt sql.NullInt64 `json:"-"`
		UpdatedAt sql.NullInt64 `json:"-"`

		// Embed
		Tables []*Table `json:"tables,omitempty"`

		// counting
		TotalTables int `json:"total_tables,omitempty"`
	}

	//Table Case Study Restaurant Dine in
	Table struct {
		ID        int           `json:"id"`
		FloorID   int           `json:"floor_id"`
		Name      string        `json:"name"`
		XPos      float64       `json:"x_pos"`
		YPos      float64       `json:"y_pos"`
		WSize     float64       `json:"w_size"`
		HSize     float64       `json:"h_size"`
		DSize     float64       `json:"d_size"`
		Capacity  int           `json:"capacity"`
		Type      string        `json:"type"`
		CreatedAt sql.NullInt64 `json:"-"`
		UpdatedAt sql.NullInt64 `json:"-"`
		// status flag
		// 	•	available
		//	•	occupied
		//	•	reserved
		//	•	disabled (for maintenance) in other words needs-cleaning
		//	•	ordering (in progress)
		//	•	billed (waiting for payment)
		Status string `json:"status"`
		// this will be set if there's an active order linked with current table
		OrderID string `json:"order_id,omitempty"`
	}
)

// Table status flow:
//
//  1. All tables start as `available` (ready for new customers).
//
//  2. A waiter can set a table to `reserved` when a customer books or calls ahead.
//
//  3. When the customer arrives and is seated (but has not ordered yet),
//     the table status should change to `occupied`.
//
//  4. Once an order is placed or sent to the server, the table status should
//     update to `ordering` (active order in progress).
//
//  5. When the customer requests the bill or checkout, the table status
//     changes to `billed` (awaiting payment).
//
//  6. After successful payment, the table status should revert to `available`,
//     or to `disabled` if it requires cleaning or maintenance before reuse.

const TableUpdateEventKey = "update.table"
