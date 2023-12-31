package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type PostRequest struct {
	UserID     int    `json:"user_id"`
	Title      string `json:"title"`
	Categories string `json:"categories"`
	Content    string `json:"content"`
	Privacy    string `json:"privacy"`
	Picture    string `json:"picture"`
	CreatedAt  string `json:"created_at"`
	AbleToSee  string `json:"able_to_see"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	PostID    int    `json:"post_ID"`
	CreatorID int    `json:"creator_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Privacy   string `json:"privacy"`
	Picture   string `json:"picture"`
	CreatedAt string `json:"created_at"`
}

type PostDTOoutElement struct {
	PostID          int    `json:"post_id"`
	CreatorID       int    `json:"creator_id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	Privacy         string `json:"privacy"`
	Picture         string `json:"picture"`
	CreatorFullName string `json:"creatorFullName"`
	CreatorEmail    string `json:"creatorEmail"`
	CreatedAt       string `json:"createdAt"`
}

func postNewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer recovery(w)
	var data PostRequest
	// We are not using the userID from the cookie as we are getting the userID from request body
	var err error
	data.UserID, err = getRequestSenderIDfromcookie(r)
	if err != nil {
		jsonResponse(w, 401, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&data)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusBadRequest, "")
		return
	}
	data.Categories = sanitizeCategories(data.Categories)
	// process the picture
	postPicture := []byte{}
	if data.Picture != "" {
		avatarData, err := base64.StdEncoding.DecodeString(data.Picture)
		if err != nil {
			log.Println(err.Error())
			jsonResponse(w, http.StatusUnprocessableEntity, "Invalid avatar")
			return
		}
		if !isImage(avatarData) {
			log.Println(err.Error())
			jsonResponse(w, http.StatusUnsupportedMediaType, "avatar is not a valid image")
			return
		}
		postPicture = avatarData
	}

	data.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	result, err := statements["addPost"].Exec(data.UserID, data.Title, data.Categories, data.Content, data.Privacy, postPicture, data.CreatedAt)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "addPost query failed")
		return
	}

	postId, err := result.LastInsertId()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "LastInsertId of addPost query failed")
		return
	}
	// privacy check and able to see
	if data.Privacy == "almost private" {
		if data.AbleToSee != "" {
			listOfEmails := strings.Split(data.AbleToSee, " ")
			for _, email := range listOfEmails {
				// get the id of the user from the email
				userID, err := getIDbyEmail(email)
				if err != nil {
					jsonResponse(w, http.StatusUnprocessableEntity, "Invalid email")
					return
				}
				// add the post to the almost_private table (user_id, post_id)
				_, err = statements["addAlmostPrivate"].Exec(userID, postId)
				if err != nil {
					jsonResponse(w, http.StatusInternalServerError, "addAlmostPrivate query failed")
					return
				}
			}
		}
	}

	jsonResponse(w, http.StatusOK, "Post created")

	rows, err := statements["getPosts"].Query(data.UserID)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusInternalServerError, "getPosts query failed")
		return
	}
	var post PostDTOoutElement
	post.CreatorID = data.UserID
	var picBlob []byte
	var fName, lName string
	rows.Next()
	err = rows.Scan(&post.PostID, &post.Title, &post.Content, &picBlob, &fName, &lName, &post.Privacy, &post.CreatorEmail, &post.CreatedAt)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "post scan failed")
		return
	}
	rows.Close()
	post.CreatorFullName = fName + " " + lName
	post.Picture = base64.StdEncoding.EncodeToString(picBlob)
	sendWSPost(post)
}

func postsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer recovery(w)
	rows, err := statements["getPosts"].Query()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "getPosts query failed")
		return
	}
	var posts Posts
	for rows.Next() {
		var post Post
		pictureBlob := []byte{}
		err = rows.Scan(&post.PostID, &post.CreatorID, &post.Title, &post.Content, &pictureBlob)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "post scan failed")
			return
		}
		post.Picture = base64.StdEncoding.EncodeToString(pictureBlob)
		posts.Posts = append(posts.Posts, post)
	}
	rows.Close()

	jsonResponseObj, err := json.Marshal(posts)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "json.Marshal failed")
		return
	}
	_, err = w.Write(jsonResponseObj)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "posts failed")
		return
	}
}

func sanitizeCategories(data string) string {
	newdata := strings.Split(data, ",")
	returndata := ""
	for i, w := range newdata {
		w = strings.TrimSpace(w)
		if w != "" && i > 0 {
			returndata += (", " + w)
		} else {
			returndata += w
		}
		if returndata == "" {
			returndata = generateRandomEmojiSequence()
		}
	}
	return returndata
}

func generateRandomEmojiSequence() string {
	rounds := []string{"🔴", "🟠", "🟡", "🟢", "🔵", "🟣", "🟤", "⚫", "⚪"}
	// Shuffle the rounds using Fisher-Yates algorithm
	for i := len(rounds) - 1; i > 0; i-- {
		bi := big.NewInt(3)
		bj, err := rand.Int(rand.Reader, bi)
		if err != nil {
			log.Fatal(err)
		}
		// convert big.Int to int
		j := int(bj.Int64())
		rounds[i], rounds[j] = rounds[j], rounds[i]
	}

	// Join the shuffled rounds into a single string
	mixedRounds := strings.Join(rounds, " ")

	return mixedRounds
}

func userPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer recovery(w)
	userID, err := getRequestSenderIDfromcookie(r)
	if err != nil {
		jsonResponse(w, 401, err.Error())
		return
	}

	// get the posts of the user
	rows, err := statements["getPostsByUserID"].Query(userID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "getPosts query failed")
		return
	}

	log.Println("starting to iterate over the rows")

	// create a slice of posts
	var posts []Post

	// iterate over the rows and append the posts to the slice
	for rows.Next() {
		var post Post
		var firstName, lastName string
		var pictureBlob []byte
		err = rows.Scan(&post.PostID, &post.Title, &post.Content, &pictureBlob, &firstName, &lastName, &post.CreatedAt)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "post scan failed")
			return
		}
		post.Picture = base64.StdEncoding.EncodeToString(pictureBlob)
		posts = append(posts, post)
	}

	// close the rows
	rows.Close()

	// create a map to store the posts
	// var postsMap map[string][]PostDTOelement

	// add the posts to the map
	// postsMap := map[string][]PostDTOoutElement{
	// 	"posts": posts,
	// }

	// // marshal the map into json
	// jsonResponseObj, err := json.Marshal(postsMap)
	// if err != nil {
	// 	jsonResponse(w, http.StatusInternalServerError, "json.Marshal(postsMap) failed")
	// 	return
	// }

	// // write the response

	// _, err = w.Write(jsonResponseObj)
	// if err != nil {
	// 	jsonResponse(w, http.StatusInternalServerError, "w.Write(jsonResponseObj)<-postsMap failed")
	// 	return
	// }

}
