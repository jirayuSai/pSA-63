// Code generated by entc, DO NOT EDIT.

package systemmember

const (
	// Label holds the string label denoting the systemmember type in the database.
	Label = "systemmember"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSystemmemberName holds the string denoting the systemmember_name field in the database.
	FieldSystemmemberName = "systemmember_name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgePrescriptions holds the string denoting the prescriptions edge name in mutations.
	EdgePrescriptions = "prescriptions"

	// Table holds the table name of the systemmember in the database.
	Table = "systemmembers"
	// PrescriptionsTable is the table the holds the prescriptions relation/edge.
	PrescriptionsTable = "prescriptions"
	// PrescriptionsInverseTable is the table name for the Prescription entity.
	// It exists in this package in order to avoid circular dependency with the "prescription" package.
	PrescriptionsInverseTable = "prescriptions"
	// PrescriptionsColumn is the table column denoting the prescriptions relation/edge.
	PrescriptionsColumn = "Systemmember_ID"
)

// Columns holds all SQL columns for systemmember fields.
var Columns = []string{
	FieldID,
	FieldSystemmemberName,
	FieldPassword,
}

var (
	// SystemmemberNameValidator is a validator for the "Systemmember_Name" field. It is called by the builders before save.
	SystemmemberNameValidator func(string) error
	// PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
