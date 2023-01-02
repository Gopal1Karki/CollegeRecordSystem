package main1234

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

// connecting  to DataBase
func connectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	fmt.Println("My sql connected sucessfully....")
	time.Sleep(1 * time.Second)

	dbm = db
}

// Creating a Student Table
func createTable() {
	query := `Create table student(
		sid int auto_increment,
		name text not null,
		address text not null,
		faculty text not null,
		dob text not null, 
		contact text not null,
		email text not null,
		notice text ,
		create_at datetime,
		primary key(sid)
		);`
	_, err := dbm.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}
func addRecord() {
	fmt.Println("Student Registration Page.")
	fmt.Printf("\n\n")
	//sid := 20220000
	var b string
	var e string
	fmt.Printf("Enter the first name: ")
	fmt.Scanln(&b)
	fmt.Printf("Enter the last name: ")
	fmt.Scanln(&e)
	name := fmt.Sprint(b, " ", e)
	var ad string
	fmt.Printf("Enter the Address: ")
	fmt.Scanln(&ad)
	address := ad
	var c string
	fmt.Printf("Enter the Faculty: ")
	fmt.Scanln(&c)
	faculty := c

	var d string
	fmt.Printf("Enter the date of Birth: ")
	fmt.Scanln(&d)
	dob := d
	var co string
	fmt.Printf("Enter the Contact number: ")
	fmt.Scanln(&co)
	var contact string = co

	var em string
	fmt.Printf("Enter you email address: ")
	fmt.Scanln(&em)
	email := em
	create_at := time.Now()

	result, err := dbm.Exec(`insert into student (name, address,faculty,dob,contact,email,create_at) value(?,?,?,?,?,?,?)`,
		name, address, faculty, dob, contact, email, create_at)
	if err != nil {
		log.Fatal(err)
	} else {
		value, _ := result.LastInsertId()
		fmt.Printf("\n")
		fmt.Println("The student record is added sucessfully of sid ", value)
	}
	fmt.Printf("\n\n")
	fmt.Println("+----------------------------+")
	fmt.Println("|       LOGIN CREDENTIAL     |")
	fmt.Println("+----------------------------+")
	fmt.Println("The student ", name, " login details is as follows:")
	fmt.Println("Username = ", email)
	fmt.Println("Password = ", dob)
	fmt.Printf("\n")
EX:
	var ex string
	fmt.Printf("Press 'a' to enter another data or 'e' to exit! ")
	fmt.Scanln(&ex)
	if ex == "e" {
		clearScreen()
		adminInterFace()
	} else if ex == "a" {
		clearScreen()
		addRecord()
	} else {
		fmt.Println("Invalid input")
		goto EX
	}
}

// to see the student record stored in the database
func seeStudentDetail() {
	fmt.Println("Welcome to admin student detail interface")
	fmt.Printf("\n\n")

	type student struct {
		sid       int
		name      string
		address   string
		faculty   string
		dob       int
		contact   string
		email     string
		create_at time.Time
	}
	var s student
	fmt.Println(" Following are the Details of student Recorded till now.")
	row, err := dbm.Query(`select * from student`)
	if err != nil {
		log.Fatal(err)
	} else {
		for row.Next() {
			row.Scan(&s.sid, &s.name, &s.address, &s.faculty, &s.dob, &s.contact, &s.email, &s.create_at)
			fmt.Println("+---------------------------------------+")
			fmt.Println("| Details of student of SID = ", s.sid, "|")
			fmt.Println("+---------------------------------------+")
			fmt.Printf("\n")
			fmt.Println("Full Name: ", s.name)
			fmt.Println("Address: ", s.address)
			fmt.Println("Faculty: ", s.faculty)
			fmt.Println("Date of Birth: ", s.dob)
			fmt.Println("Contact number: ", s.contact)
			fmt.Println("Email: ", s.email)
			fmt.Println("Student Profile Created time: ", s.create_at)
			fmt.Printf("\n")
			fmt.Println("        LOGIN CREDENTIAL      ")
			fmt.Println("        ----------------      ")
			fmt.Println("sid = ", s.sid)
			fmt.Println("Username = ", s.email)
			fmt.Println("Password = ", s.dob)
			fmt.Println("******************************************")
			fmt.Printf("\n\n")
		}
	}
EX:
	var ex string
	fmt.Printf("Press 'e' to exit! ")
	fmt.Scanln(&ex)
	if ex == "e" {
		clearScreen()
		adminInterFace()
	} else {
		fmt.Println("Invalid input")
		goto EX
	}
}

// deleting a student record from the database
func removeRecord() {
	fmt.Println("Welcome to Delete section interface of student record ")
	fmt.Printf("\n\n")
	var a int
	fmt.Printf("Enter the sid of student which you want to delete from Record: ")
	fmt.Scanln(&a)
	sid := a
A:
	var del string
	fmt.Printf("Are you sure you want to delete the record of SID? (y/n): ")
	fmt.Scanln(&del)
	if del == "y" {
		_, err := dbm.Exec(`delete from student where sid = ?`, sid)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\n\n The data is deleted of sid = ", sid)
		}
		fmt.Printf("\n")
	EX:
		var ex string
		fmt.Printf("Press 'd' to delete another data or 'e' to exit!: ")
		fmt.Scanln(&ex)
		if ex == "e" {
			clearScreen()
			adminInterFace()
		} else if ex == "d" {
			clearScreen()
			removeRecord()
		} else {
			fmt.Println("Invalid input")
			goto EX
		}

	} else if del == "n" {
		clearScreen()
		removeRecord()
	} else {
		fmt.Println("Invalid Input! ")
		goto A
	}
}

// update data in student record
func updateStudentDetail() {
	fmt.Println("Welcome to Update student detail interface.")
	fmt.Printf("\n\n")
	var a int
	fmt.Printf("Enter the Sid: ")
	fmt.Scanln(&a)
	sid := a
	var b string
	var e string
	fmt.Printf("Enter the first name: ")
	fmt.Scanln(&b)
	fmt.Printf("Enter the last name: ")
	fmt.Scanln(&e)
	name := fmt.Sprint(b, " ", e)
	var ad string
	fmt.Printf("Enter the Address: ")
	fmt.Scanln(&ad)
	address := ad
	var c string
	fmt.Printf("Enter the Faculty: ")
	fmt.Scanln(&c)
	faculty := c

	var d int
	fmt.Printf("Enter the date of Birth: ")
	fmt.Scanln(&d)
	dob := d
	var co string
	fmt.Printf("Enter the Contact number: ")
	fmt.Scanln(&co)
	contact := co

	var em string
	fmt.Printf("Enter you email address: ")
	fmt.Scanln(&em)
	email := em

	result, err := dbm.Exec(`update student set name = ?,address= ?,faculty = ?,dob=?,contact = ?, email= ? where sid = ?`,
		name, address, faculty, dob, contact, email, sid)
	if err != nil {
		log.Fatal(err)
	} else {
		value, _ := result.LastInsertId()
		fmt.Println("\n\nRecord update for student id ", value)
	}
EX:
	var ex string
	fmt.Printf("Press 'u' to Update another data or 'e' to exit! ")
	fmt.Scanln(&ex)
	if ex == "e" {
		clearScreen()
		adminInterFace()
	} else if ex == "u" {
		clearScreen()
		updateStudentDetail()
	} else {
		fmt.Println("Invalid input")
		goto EX
	}
}

func dbNotice() {
	fmt.Println("Welcome to Notice Section!")

	fmt.Printf("\n\n")
	a := "Tommorow is a public holiday."
	if a == " " {
		fmt.Println("There is no notice")

	} else {
		fmt.Println("Notice: ", a)
	}
}

// Clearing the Screen
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func firstScreen() {
D:
	var inp string
	fmt.Printf("\n\n\n")
	fmt.Printf("Press c to continue: ")
	fmt.Scanln(&inp)
	switch {
	case inp == "c":
		clearScreen()
		secondScreen()
	default:
		fmt.Println("Invalid Input")
		time.Sleep(1 * time.Second)
		goto D
	}

}

func secondScreen() {
A:
	clearScreen()
	fmt.Println("Press 1 for Admin Sector Login: ")
	fmt.Println("Press 2 for Student Sector Login: ")
	fmt.Println("Press 3 for Exit: ")
	fmt.Printf("\n\n\n")

	var inp1 int
	fmt.Printf("Enter your choices: ")
	fmt.Scanln(&inp1)
	ch := inp1
	switch ch {
	case 1:
		clearScreen()
		loginScreen1()
	case 2:
		clearScreen()
		loginScreen2()
	case 3:
		var lg string
		fmt.Printf("Are you sure you want to exit? (y/n): ")
		fmt.Scanln(&lg)
		if lg == "y" {
			clearScreen()
			os.Exit(5)
		} else {
			clearScreen()
			secondScreen()
		}

	default:
		fmt.Println("Enter the valid Input")
		goto A
	}
}

func loginScreen1() {
B:
	clearScreen()
	fmt.Println("Admin Login Portal")
	fmt.Printf("\n\n\n")

	fmt.Print("Username: ")
	var adUsername string
	fmt.Scanln(&adUsername)
	fmt.Print("Password: ")
	var adPassword string
	fmt.Scanln(&adPassword)

	if adUsername == "admin" && adPassword == "minad" {
		clearScreen()
		adminInterFace()
	} else {
		fmt.Println("Invalid Username or Password")
		time.Sleep(1 * time.Second)
		goto B
	}

}

func adminInterFace() {
	fmt.Println("Welcome to Admin User InterFace:")
	fmt.Printf("\n\n")
	fmt.Println("Press 1 to Add Student Record")
	fmt.Println("Press 2 to see Student Details")
	fmt.Println("Press 3 to Remove Student Record ")
	fmt.Println("Press 4 to Update Details of Student")
	fmt.Println("Press 5 to Post Notice for student")
	fmt.Println("Press 6 to Logout")
	fmt.Printf("\n\n")
	var inp2 int
	fmt.Printf("Enter the choices: ")
	fmt.Scanln(&inp2)
	switch inp2 {
	case 1:
		clearScreen()
		addRecord()
	case 2:
		clearScreen()
		seeStudentDetail()
	case 3:
		clearScreen()
		removeRecord()
	case 4:
		clearScreen()
		updateStudentDetail()
	case 5:
		clearScreen()
		dbNotice()
		var b string
	X:
		fmt.Printf("Press 'e' to exit !")
		fmt.Scanln(&b)
		if b == "e" {
			clearScreen()
			adminInterFace()
		} else {
			fmt.Println("Invalid Input")
			time.Sleep(1 * time.Second)
			clearScreen()
			goto X
		}
	case 6:
		clearScreen()
		adminLogout()
	default:
		fmt.Println("Enter the valid Input")
	}
}

func adminLogout() {
	fmt.Println("Welcome to admin logout interface! ")
invalidInput:
	fmt.Printf("\nAre you sure you want to logout? (y/n): ")
	var adLg string
	fmt.Scanln(&adLg)
	if adLg == "y" {
		clearScreen()
		secondScreen()
	} else if adLg == "n" {
		clearScreen()
		adminInterFace()
	} else {
		fmt.Println("Invalid Input")
		goto invalidInput
	}
}
func loginScreen2() {
	counte := 0
A:
	if counte >= 3 {
		fmt.Printf("\n\n")
		fmt.Println("+--------------------------------------------------------------------------------+")
		fmt.Println("| You have been banned for 5 second from login because of multiple incorrect sid |")
		fmt.Println("+--------------------------------------------------------------------------------+")
		time.Sleep(5 * time.Second)
	C:
		fmt.Printf("Press 'e' to go back to menu! ")
		var ab string
		fmt.Scanln(&ab)
		if ab == "e" {
			clearScreen()
			secondScreen()
		} else {
			fmt.Println("Invalid Input!")
			goto C
		}

	}
	fmt.Println("Welcome to Student Login interface")
	fmt.Printf("\n\n")

	type student struct {
		sid       int
		name      string
		address   string
		faculty   string
		dob       int
		contact   string
		email     string
		create_at time.Time
	}
	var s student
	var a int
	count := 0
	fmt.Printf("Enter the Student Id: ")
	fmt.Scanln(&a)
	row, err := dbm.Query(`select * from student`)
	if err != nil {
		log.Fatal(err)
	} else {
		for row.Next() {
			row.Scan(&s.sid, &s.name, &s.address, &s.faculty, &s.dob, &s.contact, &s.email, &s.create_at)
			time.Sleep(1 * time.Second)
			if a == s.sid {
			B:
				if count >= 3 {
					fmt.Printf("\n\n")
					fmt.Println("+------------------------------------------------------------------------------------------+")
					fmt.Println("| You have been banned for 5 second from login because of multiple incorrect login attempt |")
					fmt.Println("+------------------------------------------------------------------------------------------+")
					time.Sleep(5 * time.Second)
					clearScreen()
					secondScreen()
				}
				clearScreen()
				fmt.Println("Student Login Portal! ")
				fmt.Println("---------------------")
				fmt.Printf("\n\n")
				var user string
				fmt.Printf("Username: ")
				fmt.Scanln(&user)

				var pass int
				fmt.Printf("Password: ")
				fmt.Scanln(&pass)
				x := s.sid

				if user == s.email && pass == s.dob {
					fmt.Println("The Login credential is correct!!!!!")
					time.Sleep(1 * time.Second)
					clearScreen()
					stdInterface(x)
				} else {
					count += 1
					fmt.Println("Invalid Username or password")
					time.Sleep(1 * time.Second)
					goto B
				}
			}
		}
	}
	fmt.Println("Invalid SID.")
	counte += 1
	time.Sleep(1 * time.Second)
	clearScreen()
	goto A
}
func stdInterface(a int) {
A:
	type student struct {
		name    string
		address string
		faculty string
		dob     int
		contact string
		email   string
	}
	var g student
	sid := a
	row, err := dbm.Query(`select name,address,faculty,dob,contact,email from student where sid = ?`,
		sid)
	if err != nil {
		log.Fatal(err)
	} else {
		for row.Next() {
			row.Scan(&g.name, &g.address, &g.faculty, &g.dob, &g.contact, &g.email)
			fmt.Printf("\n\n")
			fmt.Println("Welcome, Dear", g.name)
			fmt.Println("---------------------------------")
			fmt.Println("\t\t\t\t\t\t User Id:", sid)
			fmt.Println("\t\t\t\t\t\t ---------------------")
			fmt.Println("Your Profile Details:")
			fmt.Println("\t Name: ", g.name)
			fmt.Println("\t Address: ", g.address)
			fmt.Println("\t Faculty: ", g.faculty)
			fmt.Println("\t Date of Birth: ", g.dob)
			fmt.Println("\t Contact Number: ", g.contact)
			fmt.Println("\t Email: ", g.email)
			fmt.Printf("\n\n")
			dbNotice()
			fmt.Printf("\n\n")
			var inp string
			fmt.Printf("Press 'e' to logout! ")
			fmt.Scanln(&inp)
			if inp == "e" {
				fmt.Printf("Are you sure you want to Logout? (y/n): ")
				var inp1 string
			B:
				fmt.Scanln(&inp1)
				if inp1 == "y" {
					clearScreen()
					secondScreen()
				} else if inp1 == "n" {
					clearScreen()
					goto A
				} else {
					fmt.Println("Invalid Input!!!")
					goto B
				}
			} else {
				fmt.Println("Invalid Input!!!")
			}
		}
	}
}

func main() {
	connectDB()
	//for the first run you have to callthe function create table and then you can delete it
	//createTable()
	clearScreen()
	fmt.Printf("\n\n")
	fmt.Println("###############################")
	fmt.Println("#                             #")
	fmt.Println("#  XYZ COllEGE RECORD SYSTEM  #")
	fmt.Println("#                             #")
	fmt.Println("###############################")
	firstScreen()
}
