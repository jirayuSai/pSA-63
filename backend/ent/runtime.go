// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/jirayuSai/app/ent/doctor"
	"github.com/jirayuSai/app/ent/mmedicine"
	"github.com/jirayuSai/app/ent/patient"
	"github.com/jirayuSai/app/ent/schema"
	"github.com/jirayuSai/app/ent/systemmember"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	doctorFields := schema.Doctor{}.Fields()
	_ = doctorFields
	// doctorDescDoctorName is the schema descriptor for Doctor_Name field.
	doctorDescDoctorName := doctorFields[0].Descriptor()
	// doctor.DoctorNameValidator is a validator for the "Doctor_Name" field. It is called by the builders before save.
	doctor.DoctorNameValidator = doctorDescDoctorName.Validators[0].(func(string) error)
	mmedicineFields := schema.Mmedicine{}.Fields()
	_ = mmedicineFields
	// mmedicineDescMmedicineName is the schema descriptor for Mmedicine_Name field.
	mmedicineDescMmedicineName := mmedicineFields[0].Descriptor()
	// mmedicine.MmedicineNameValidator is a validator for the "Mmedicine_Name" field. It is called by the builders before save.
	mmedicine.MmedicineNameValidator = mmedicineDescMmedicineName.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientName is the schema descriptor for Patient_Name field.
	patientDescPatientName := patientFields[0].Descriptor()
	// patient.PatientNameValidator is a validator for the "Patient_Name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	// patientDescPatientPhone is the schema descriptor for Patient_Phone field.
	patientDescPatientPhone := patientFields[2].Descriptor()
	// patient.PatientPhoneValidator is a validator for the "Patient_Phone" field. It is called by the builders before save.
	patient.PatientPhoneValidator = patientDescPatientPhone.Validators[0].(func(int) error)
	systemmemberFields := schema.Systemmember{}.Fields()
	_ = systemmemberFields
	// systemmemberDescSystemmemberName is the schema descriptor for Systemmember_Name field.
	systemmemberDescSystemmemberName := systemmemberFields[0].Descriptor()
	// systemmember.SystemmemberNameValidator is a validator for the "Systemmember_Name" field. It is called by the builders before save.
	systemmember.SystemmemberNameValidator = systemmemberDescSystemmemberName.Validators[0].(func(string) error)
	// systemmemberDescPassword is the schema descriptor for Password field.
	systemmemberDescPassword := systemmemberFields[1].Descriptor()
	// systemmember.PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	systemmember.PasswordValidator = func() func(string) error {
		validators := systemmemberDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Password string) error {
			for _, fn := range fns {
				if err := fn(_Password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
