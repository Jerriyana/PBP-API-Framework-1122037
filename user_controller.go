package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	model "martini/models"
	"net/http"

	"github.com/codegangsta/martini"
)

// 1. GET ALL USERS
func GetAllUsers(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect(w)
	defer db.Close()

	query := "SELECT * FROM users WHERE 1=1" // Untuk menambahkan klausa WHERE dengan aman

	// Read from Query Param
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	if name != "" {
		query += " AND name='" + name + "'  "
	}

	if age != "" {
		query += " AND age= '" + age + "' "
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		var user model.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Password, &user.Email, &user.User_type); err != nil {
			log.Println(err)
			sendErrorResponse(w, 500, "internal error")
			return
		} else {
			users = append(users, user)
		}
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	sendGetUsersResponse(w, 200, "berhasil", users)
}

// 2. INSERT USER (POST)
func InsertUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect(w)
	defer db.Close()

	query := "INSERT INTO users (ID, Name, Age, Address, Password, Email, User_type) VALUES (?, ?, ?, ?, ?, ?,?)"
	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	// Mendapatkan data dari request body
	var user model.Users
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	// Menjalankan pernyataan Exec dengan data dari request body
	_, err = stmt.Exec(user.ID, user.Name, user.Age, user.Address)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User berhasil dibuat"))

	// Mengisi nilai objek users
	var users []model.Users
	users = append(users, user)

	var response model.UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users

	// Mengirimkan response JSON
	json.NewEncoder(w).Encode(response)
}

// 3. UPDATE USER (PUT)
func UpdateUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect(w)
	defer db.Close()

	// Mendapatkan ID dari path parameter
	userID := params["id"]

	// Membaca data dari request body
	var updatedUser model.Users
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	// Mengeksekusi pernyataan SQL untuk mengupdate user berdasarkan ID
	query := "UPDATE users SET Name=?, Age=?, Address=? WHERE ID=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	_, err = stmt.Exec(updatedUser.Name, updatedUser.Age, updatedUser.Address, userID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	sendSuccessResponse(w, 200, "berhasil update")
}

// 4. DELETE USER (DELETE)
func DeleteUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := connect(w)
	defer db.Close()

	// Mendapatkan ID dari path parameter
	userID := params["id"]

	// Mengeksekusi pernyataan SQL untuk mendapatkan data user sebelum dihapus
	querySelect := "SELECT Name FROM users WHERE ID=?"
	var userName string
	err := db.QueryRow(querySelect, userID).Scan(&userName)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User dengan ID tersebut tidak ditemukan"))
		return
	}

	// Mengeksekusi pernyataan SQL untuk menghapus user berdasarkan ID
	queryDelete := "DELETE FROM users WHERE ID=?"
	stmt, err := db.Prepare(queryDelete)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	_, err = stmt.Exec(userID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}

	// Membuat response JSON
	message := fmt.Sprintf("Data dengan ID %s dan nama %s dihapus", userID, userName)
	sendSuccessResponse(w, 200, message)
}
