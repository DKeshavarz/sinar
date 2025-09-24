package postgres

import (
    "database/sql"
    "github.com/DKeshavarz/sinar/internal/domain"
)

type UniversityRepository struct {
    DB *sql.DB
}

func NewUniversityRepository() *UniversityRepository{
    return &UniversityRepository{
        DB: getDB(),
    }
}

func (r *UniversityRepository) Get(id int) (*domain.University, error) {
    university := &domain.University{}
    query := `
        SELECT id, name, location, logo
        FROM universities WHERE id = $1`
    err := r.DB.QueryRow(query, id).Scan(&university.ID, &university.Name, &university.Location, &university.Logo)
    if err == sql.ErrNoRows {
        return nil, err
    }
    return university, nil
}

func (r *UniversityRepository) Create(university *domain.University) error {
    query := `
        INSERT INTO universities (name, location, logo)
        VALUES ($1, $2, $3) RETURNING id`
    return r.DB.QueryRow(query, university.Name, university.Location, university.Logo).Scan(&university.ID)
}

func (r *UniversityRepository) Update(university *domain.University) error {
    query := `
        UPDATE universities SET name = $1, location = $2, logo = $3
        WHERE id = $4`
    _, err := r.DB.Exec(query, university.Name, university.Location, university.Logo, university.ID)
    return err
}

func (r *UniversityRepository) Delete(id int) error {
    query := `DELETE FROM universities WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}