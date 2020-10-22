package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jirayuSai/app/controllers"
	"github.com/jirayuSai/app/ent"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SystemMembers struct {
	SystemMember []SystemMember
}
type SystemMember struct {
	SystemMemberName string
	Password         string
}
type Prescriptions struct {
	Prescription []Prescription
}
type Prescription struct {
	PrescripID string
}

type Patients struct {
	Patient []Patient
}
type Patient struct {
	PatientName   string
	Gender        string
	PatientNumber int
}
type Doctors struct {
	Doctor []Doctor
}
type Doctor struct {
	DoctorName  string
	DoctorEmail string
}
type Mmedicines struct {
	Mmedicine []Mmedicine
}
type Mmedicine struct {
	MmedicineName string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//Set Doctor Data
	doctors := Doctors{
		Doctor: []Doctor{
			Doctor{"Jirayu Rattanawong", "b6113049@g.sut.ac.th"},
			Doctor{"Roberts", "roberts@gmail.com"},
			Doctor{"Taylor", "taylor@gmail.com"},
			Doctor{"Martin", "Martin@g.sut.ac.th"},
			Doctor{"Peretz", "Peretz@gmail.com"},
			Doctor{"Stewart", "Stewart@gmail.com"},
		},
	}
	for _, d := range doctors.Doctor {
		client.Doctor.
			Create().
			SetDoctorName(d.DoctorName).
			SetDoctorEmail(d.DoctorEmail).
			Save(context.Background())

	}

	//Set Mmedicine Data
	mmedicines := Mmedicines{
		Mmedicine: []Mmedicine{
			Mmedicine{"ยาน้ำแก้ไอ"},
			Mmedicine{"ยานอนหลับ"},
			Mmedicine{"ยาขับเสมหะ"},
			Mmedicine{"ยาแก้ปวด"},
			Mmedicine{"ยาปฏิชีวนะ"},
			Mmedicine{"ยาระบาย"},
		},
	}
	for _, me := range mmedicines.Mmedicine {
		client.Mmedicine.
			Create().
			SetMmedicineName(me.MmedicineName).
			Save(context.Background())

	}

	//Set Doctor Data
	patients := Patients{
		Patient: []Patient{
			Patient{"Morgan", "female", 1234567890},
			Patient{"Beckham", "male", 9876787809},
			Patient{"Harris", "female", 5555555555},
			Patient{"Gonzalez", "female", 9998652431},
			Patient{"Jackson", "male", 9999999999},
			Patient{"Nelson", "male", 5423676543},
		},
	}
	for _, pt := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(pt.PatientName).
			SetGender(pt.Gender).
			SetPatientPhone(pt.PatientNumber).
			Save(context.Background())

	}

	//Set Systemmember Data
	systemMembers := SystemMembers{
		SystemMember: []SystemMember{
			SystemMember{"Edwards", "12345678"},
		},
	}
	for _, sm := range systemMembers.SystemMember {
		client.Systemmember.
			Create().
			SetSystemmemberName(sm.SystemMemberName).
			SetPassword(sm.Password).
			Save(context.Background())
	}

	v1 := router.Group("/api/v1")
	controllers.NewDoctorController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewPrescriptionController(v1, client)
	controllers.NewSystemmemberController(v1, client)
	controllers.NewPatiMmedicineController(v1, client)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
