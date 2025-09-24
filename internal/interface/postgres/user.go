package postgres

import (
    "database/sql"
    "errors"
    "github.com/DKeshavarz/sinar/internal/dto"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) GetByStudentNumber(number string) (*dto.UserWithUniversity, error) {
    if number == "" {
        return nil, errors.New("student number cannot be empty")
    }

    userWithUniversity := &dto.UserWithUniversity{}
    query := `
        SELECT u.id, u.first_name, u.last_name, u.phone, u.profile_pic, u.student_num, u.sex, 
               uni.id, uni.name, uni.location, uni.logo
        FROM users u
        LEFT JOIN universities uni ON u.university_id = uni.id
        WHERE u.student_num = $1`
    err := r.DB.QueryRow(query, number).Scan(
        &userWithUniversity.User.ID,
        &userWithUniversity.User.FirstName,
        &userWithUniversity.User.LastName,
        &userWithUniversity.User.Phone,
        &userWithUniversity.User.ProfilePic,
        &userWithUniversity.User.StudentNum,
        &userWithUniversity.User.Sex,
        &userWithUniversity.University.ID,
        &userWithUniversity.University.Name,
        &userWithUniversity.University.Location,
        &userWithUniversity.University.Logo,
    )
    if err == sql.ErrNoRows {
        return nil, errors.New("user not found")
    }
    if err != nil {
        return nil, err
    }
    return userWithUniversity, nil
}