package blog

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"strconv"
)

// shit shit shit
var PostType = map[int]model.Post{}

func getId(rc *gin.Context) (int, error) {
	var pt model.PageType
	if err := rc.ShouldBindUri(&pt); err != nil {
		return 0, errors.New("ERROR: Id is not set")
	}

	id, err := strconv.Atoi(pt.Id)
	if err != nil {
		return 0, errors.New("ERROR: Id is not integer")
	}

	return id, nil
}

func ShowPosts(rc *gin.Context) {
	rc.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Home", "data": PostType})
}

func ShowPost(rc *gin.Context) {
	id, err := getId(rc)
	if err != nil {
		rc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(id)
	if id > len(PostType) {
		rc.JSON(http.StatusNotFound, gin.H{"post": "Not Found"})
		return
	}

	rc.HTML(http.StatusOK, "post.tmpl", gin.H{"title": "Show Post", "data": PostType[id]})
}

func AddPost(rc *gin.Context) {
	title := rc.PostForm("title")
	post_message := rc.PostForm("post_message")
	max_len := len(PostType)
	PostType[max_len+1] = model.Post{max_len + 1, title, post_message}
	rc.Redirect(http.StatusFound, "/")
}

func EditPost(rc *gin.Context) {
	id, err := getId(rc)
	if err != nil {
		rc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := rc.PostForm("title")
	post_message := rc.PostForm("post_message")
	max_len := len(PostType)

	// if we try edit unexist post
	if id > max_len {
		rc.JSON(http.StatusNotFound, gin.H{"post": "Not Found"})
		return
	}

	PostType[id] = model.Post{id, title, post_message}
	rc.Redirect(301, "/")
}

func ShowAddPost(rc *gin.Context) {
	rc.HTML(http.StatusOK, "add_post.tmpl", gin.H{"title": "Add Post"})
}

func ShowEditPost(rc *gin.Context) {
	id, err := getId(rc)
	if err != nil {
		rc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id > len(PostType) {
		rc.JSON(http.StatusNotFound, gin.H{"post": "Not Found"})
		return
	}

	rc.HTML(http.StatusOK, "edit.tmpl", gin.H{"title": "Edit Post", "data": PostType[id]})
}
