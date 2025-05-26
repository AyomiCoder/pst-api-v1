package handlers

import (
	"godb/db"
	"godb/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreatePost creates a new blog post
func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Create post
	query := `
		INSERT INTO posts (title, content, author_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`

	var post models.Post
	err := db.DB.QueryRow(query, input.Title, input.Content, userID).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	post.AuthorID = userID.(int)

	c.JSON(http.StatusCreated, post)
}

// GetPost retrieves a single post by ID
func GetPost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	query := `
		SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at, u.username
		FROM posts p
		JOIN users u ON p.author_id = u.id
		WHERE p.id = $1`

	err = db.DB.QueryRow(query, postID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.AuthorUsername,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPosts retrieves all posts with pagination
func GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query := `
		SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at, u.username
		FROM posts p
		JOIN users u ON p.author_id = u.id
		ORDER BY p.created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.AuthorUsername,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning posts"})
			return
		}
		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}

// UpdatePost updates an existing post
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if post exists and belongs to user
	var authorID int
	err = db.DB.QueryRow("SELECT author_id FROM posts WHERE id = $1", postID).Scan(&authorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if authorID != userID.(int) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this post"})
		return
	}

	// Update post
	query := `
		UPDATE posts 
		SET title = COALESCE($1, title),
			content = COALESCE($2, content),
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING id, title, content, author_id, created_at, updated_at`

	var post models.Post
	err = db.DB.QueryRow(query, input.Title, input.Content, postID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if post exists and belongs to user
	var authorID int
	err = db.DB.QueryRow("SELECT author_id FROM posts WHERE id = $1", postID).Scan(&authorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if authorID != userID.(int) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this post"})
		return
	}

	// Delete post
	_, err = db.DB.Exec("DELETE FROM posts WHERE id = $1", postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
} 