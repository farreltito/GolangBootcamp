package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

// ---------------------
// Array (data statis)
// ---------------------
var subjects = [3]string{"Math", "Science", "History"}

// ---------------------
// Slice (data dinamis)
// ---------------------
var students []Student

// ---------------------
// Map (akses cepat berdasarkan ID)
// ---------------------
var studentMap = make(map[int]Student)

// ---------------------
// Handlers
// ---------------------

// GET /subjects -> ambil semua mata pelajaran dari array
func getSubjects(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(subjects)
}

// GET /students -> ambil semua siswa dari slice
func getStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(students)
}

// POST /students -> tambahkan siswa ke slice & map
func createStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	students = append(students, newStudent)
	studentMap[newStudent.ID] = newStudent
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)
}

// GET /students/{id} -> ambil siswa dari map berdasarkan ID
func getStudentByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/students/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
		return
	}

	student, exists := studentMap[id]
	if !exists {
		http.Error(w, "Siswa tidak ditemukan", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

// ---------------------
// Main
// ---------------------
func main() {
	http.HandleFunc("/subjects", getSubjects)
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getStudents(w, r)
		} else if r.Method == http.MethodPost {
			createStudent(w, r)
		}
	})
	http.HandleFunc("/students/", getStudentByID) // untuk GET /students/{id}

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
