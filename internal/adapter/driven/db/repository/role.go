package repository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/nullexp/finman-user-service/internal/domain/model"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (rr RoleRepository) CreateRole(ctx context.Context, role model.Role) (string, error) {
	query := `
        INSERT INTO roles (name, permissions, created_at, updated_at)
        VALUES ($1, $2, NOW(), NOW())
        RETURNING id
    `
	var id string
	err := rr.db.QueryRowContext(ctx, query, role.Name, pq.Array(role.Permissions)).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (rr RoleRepository) GetRoleById(ctx context.Context, id string) (*model.Role, error) {
	query := `SELECT id, name, permissions, created_at, updated_at FROM roles WHERE id = $1`
	row := rr.db.QueryRowContext(ctx, query, id)

	var role model.Role
	if err := row.Scan(&role.Id, &role.Name, pq.Array(&role.Permissions), &role.CreatedAt, &role.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &role, nil
}

func (rr RoleRepository) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	query := `SELECT id, name, permissions, created_at, updated_at FROM roles`
	rows, err := rr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.Id, &role.Name, pq.Array(&role.Permissions), &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (rr RoleRepository) UpdateRole(ctx context.Context, role model.Role) error {
	query := `
        UPDATE roles
        SET name = $1, permissions = $2, updated_at = NOW()
        WHERE id = $3
    `
	_, err := rr.db.ExecContext(ctx, query, role.Name, pq.Array(role.Permissions), role.Id)
	if err != nil {
		return err
	}
	return nil
}

func (rr RoleRepository) DeleteRole(ctx context.Context, id string) error {
	query := `DELETE FROM roles WHERE id = $1`
	_, err := rr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
