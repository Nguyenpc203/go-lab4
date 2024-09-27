package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// Cấu trúc (struct) sinh viên
type Student struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Major  string `json:"major"`
}

// Tạo danh sách sinh viên lưu trữ
var students = []Student{
    {ID: 1, Name: "Nguyen Van A", Age: 20, Major: "Computer Science"},
    {ID: 2, Name: "Le Thi B", Age: 21, Major: "Mathematics"},
}

func main() {
    r := gin.Default()

    // GET - Lấy danh sách sinh viên
    r.GET("/get-students", getStudents)

    // GET - Lấy chi tiết một sinh viên theo ID
    r.GET("/get-student-detail/:id", getStudentDetail)

    // POST - Thêm sinh viên mới
    r.POST("/add-student", addStudent)

    // PUT - Cập nhật sinh viên
    r.PUT("/update-student/:id", updateStudent)

    // DELETE - Xóa sinh viên
    r.DELETE("/delete-student/:id", deleteStudent)

    // Chạy server tại cổng 8080
    r.Run(":8080")
}

// Lấy danh sách sinh viên
func getStudents(c *gin.Context) {
    c.JSON(http.StatusOK, students)
}

// Lấy chi tiết sinh viên theo ID
func getStudentDetail(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    for _, student := range students {
        if student.ID == id {
            c.JSON(http.StatusOK, student)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy sinh viên"})
}

// Thêm sinh viên mới
func addStudent(c *gin.Context) {
    var newStudent Student
    if err := c.BindJSON(&newStudent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newStudent.ID = len(students) + 1 // Tạo ID mới cho sinh viên
    students = append(students, newStudent)
    c.JSON(http.StatusOK, newStudent)
}

// Cập nhật thông tin sinh viên
func updateStudent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    var updatedStudent Student
    if err := c.BindJSON(&updatedStudent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, student := range students {
        if student.ID == id {
            students[i].Name = updatedStudent.Name
            students[i].Age = updatedStudent.Age
            students[i].Major = updatedStudent.Major
            c.JSON(http.StatusOK, students[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy sinh viên"})
}

// Xóa sinh viên theo ID
func deleteStudent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    for i, student := range students {
        if student.ID == id {
            students = append(students[:i], students[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Xóa thành công"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Không tìm thấy sinh viên"})
}
