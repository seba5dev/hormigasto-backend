package repository_test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/seba5dev/hormigasto-backend/internal/repository"
	"github.com/seba5dev/hormigasto-backend/models"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// Setup para una DB de test en memoria
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&models.User{})
	require.NoError(t, err)
	return db
}

func TestUserRepository_CreateAndFind(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	user := &models.User{
		Username:     "sebastian",
		Email:        "seba@example.com",
		PasswordHash: "superhash",
		IsActive:     true,
	}

	// Crear usuario
	err := repo.Create(user)
	require.NoError(t, err)
	require.NotZero(t, user.ID)

	// Buscar por ID
	found, err := repo.FindByID(user.ID)
	require.NoError(t, err)
	require.Equal(t, user.Email, found.Email)
	require.Equal(t, user.Username, found.Username)

	// Buscar por email
	found2, err := repo.FindByEmail(user.Email)
	require.NoError(t, err)
	require.Equal(t, user.ID, found2.ID)
}

func TestUserRepository_UpdateAndDelete(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	user := &models.User{
		Username:     "editme",
		Email:        "edit@example.com",
		PasswordHash: "hash",
		IsActive:     true,
	}
	err := repo.Create(user)
	require.NoError(t, err)

	// Update
	user.IsActive = false
	err = repo.Update(user)
	require.NoError(t, err)

	found, err := repo.FindByID(user.ID)
	require.NoError(t, err)
	require.False(t, found.IsActive)

	// Delete
	err = repo.Delete(user.ID)
	require.NoError(t, err)

	_, err = repo.FindByID(user.ID)
	require.Error(t, err)
}

func TestUserRepository_FindAll(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	// Crea varios usuarios
	users := []models.User{
		{Username: "u1", Email: "u1@example.com", PasswordHash: "h1", IsActive: true},
		{Username: "u2", Email: "u2@example.com", PasswordHash: "h2", IsActive: true},
		{Username: "u3", Email: "u3@example.com", PasswordHash: "h3", IsActive: false},
	}
	for _, user := range users {
		err := repo.Create(&user)
		require.NoError(t, err)
	}

	result, err := repo.FindAll()
	require.NoError(t, err)
	require.Len(t, result, 3)

	// Chequea que estén los usuarios correctos
	emails := map[string]bool{}
	for _, u := range result {
		emails[u.Email] = true
	}
	require.True(t, emails["u1@example.com"])
	require.True(t, emails["u2@example.com"])
	require.True(t, emails["u3@example.com"])
}
