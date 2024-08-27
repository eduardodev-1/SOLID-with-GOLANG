// Sem usar (Single Responsibility Principle - SRP)
// package main
//
// import "fmt"
//
//	type User struct {
//		Name  string
//		Email string
//	}
//
//	func (u *User) Save() {
//		fmt.Println("Saving user to database")
//	}
//
//	func (u *User) SendEmail() {
//		fmt.Println("Sending email to:", u.Email)
//	}
//
//	func main() {
//		user := User{"John", "john@example.com"}
//		user.Save()
//		user.SendEmail()
//	}
//

// Usando (Single Responsibility Principle - SRP)
package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type UserRepository struct{}

func (ur *UserRepository) Save(u User) {
	fmt.Println("Saving user to database", u.Name)
}

type EmailService struct{}

func (es *EmailService) SendEmail(u User) {
	fmt.Println("Sending email to:", u.Email)
}

func main() {
	user := User{"John", "john@example.com"}
	userRepository := UserRepository{}
	emailService := EmailService{}

	userRepository.Save(user)
	emailService.SendEmail(user)
}
