package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mmvergara/golang-gala/backend/model"
	"github.com/mmvergara/golang-gala/backend/repository/post"
)

type Post struct {
	Repo *post.RedisRepo
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		AuthorID        uuid.UUID `json:"author_id"`
		PostTitle       string    `json:"post_title"`
		PostDescription string    `json:"post_description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("Error decoding body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()
	post := model.Post{
		PostID:          uuid.Must(uuid.NewRandom()),
		AuthorID:        body.AuthorID,
		PostTitle:       body.PostTitle,
		PostDescription: body.PostDescription,
		CreatedAt:       &now,
	}
	fmt.Printf("Post: %+v\n", post)

	if err := p.Repo.Create(r.Context(), post); err != nil {
		fmt.Println("Error inserting post", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error marshalling post", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func (p *Post) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LIST POST +======= 1")
	cursor := r.URL.Query().Get("cursor")
	if cursor == "" {
		cursor = "0"
	}

	const decimal = 10
	const bitSize = 64

	cursorInt, err := strconv.ParseUint(cursor, decimal, bitSize)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("LIST POST +======= 2")

	const size = 1

	res, err := p.Repo.FindAll(r.Context(), post.FindAllPage{
		Size:   size,
		Offset: cursorInt,
	})

	if err != nil {
		fmt.Println("Error getting posts", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("LIST POST +======= 3")

	var response struct {
		Posts      []model.Post `json:"posts"`
		NextCursor uint64       `json:"next_cursor,omitempty"`
	}

	response.Posts = res.Posts
	response.NextCursor = res.Cursor

	resJSON, err := json.Marshal(response)
	fmt.Println("LIST POST +======= 4")

	if err != nil {
		fmt.Println("Error marshalling posts", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("LIST POST +======= 61")


	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	pID, err := uuid.Parse(postID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := p.Repo.FindById(r.Context(), pID)
	if errors.Is(err, post.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Println("failed to marshall", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (p *Post) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	postID, err := uuid.Parse(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.Repo.DeleteByID(r.Context(),postID)
	if errors.Is(err, post.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
