package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	uuid "github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

type signupData struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Dob         string `json:"dob"`
	Avatar      string `json:"avatar"`
	avatarBytes []byte `sqlite3:"avatar"`
	Nickname    string `json:"nickname"`
	AboutMe     string `json:"aboutMe"`
	Public      bool   `json:"public"`
	Privacy     string `sqlite3:"privacy"`
}

type ProfileData struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Dob         string `json:"dob"`
	Avatar      string `json:"avatar"`
	avatarBytes []byte `sqlite3:"avatar"`
	Nickname    string `json:"nickname"`
	AboutMe     string `json:"aboutMe"`
	Public      bool   `json:"public"`
	Privacy     string `sqlite3:"privacy"`
}

type ProfileDTOtoFrontend struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Dob       string `json:"dob"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	AboutMe   string `json:"aboutMe"`
	Privacy   string `json:"privacy"`
}

type loginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UUIDData struct {
	UUID string `json:"UUID"`
}

func changePrivacyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer recovery(w)

	ID, err := getRequestSenderIDfromcookie(r)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	incomingData := map[string]any{}
	// decode the request body into the DTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&incomingData)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad request")
		return
	}
	wantPublic := incomingData["public"].(bool)

	privacyvalue := map[bool]string{true: "public", false: "private"}[wantPublic]

	_, err = statements["updateUserPrivacy"].Exec(privacyvalue, ID)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusInternalServerError, "updateUserPrivacy query failed")
		return
	}

	jsonResponse(w, http.StatusOK, "Privacy updated")
}

// userProfileHandler the retrieval of a user's profile information and returns
// it in JSON format.
//
// @r.params (optional) { int } id - The id of the user whose profile we want to view.
// If not provided, the profile of the user that is currently logged in is returned.
//
// Outgoing DTO has profile Info + Posts by User + User follows + User followers
//
// Outgoing DTO is called FullProfileDTOtoFrontend
func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer recovery(w)

	myID, err := getRequestSenderIDfromcookie(r)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	seekingOwnProfile := false
	idofUserToSee := 0
	if r.URL.Path == "/profile/" {
		seekingOwnProfile = true
		log.Println("seekingOwnProfile: ", seekingOwnProfile)
	} else {
		idofUserToSee, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/profile/"))
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, "Bad request")
			return
		}
		if idofUserToSee == myID {
			seekingOwnProfile = true
		}
	}
	if seekingOwnProfile {
		jsonResponse(w, 200, giveMeMyFullProfile(myID))
	} else {
		log.Println("idofUserToSee: ", idofUserToSee)
		if doesUserExist(idofUserToSee) {
			if canISeeThisProfile(myID, idofUserToSee) {
				jsonResponse(w, 200, giveMeOtherPersonFullProfile(myID, idofUserToSee))
			} else {
				jsonResponse(w, http.StatusUnauthorized, "You cannot see this profile")
			}
		} else {
			jsonResponse(w, http.StatusNotFound, "User not found")
		}
	}
}

// isFollowing checks if the secondID follows the firstID
func isFollowing(firstID int, secID int) (bool, error) {
	rows, err := statements["doesSecondFollowFirst"].Query(secID, firstID)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func userRegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer recovery(w)

	var data signupData
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusBadRequest, "")
		return
	}

	// if dob is empty, return error
	if data.Dob == "" {
		jsonResponse(w, http.StatusUnprocessableEntity, "Invalid date of birth")
		return
	}

	// check the avatar validity
	if data.Avatar != "" {
		avatarData, err := base64.StdEncoding.DecodeString(data.Avatar)
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
		data.avatarBytes = avatarData
	}

	if data.Avatar == "" {
		rn := randomNum(0, 5)
		defaultAvatar, err := os.Open("./assets/images/profile/defaults/" + strconv.Itoa(rn) + ".jpeg")
		if err != nil {
			log.Println(err.Error())
			jsonResponse(w, http.StatusFailedDependency, "failed to load default avatar")
			return
		}
		defer defaultAvatar.Close()
		data.avatarBytes, err = ioutil.ReadAll(defaultAvatar)
		defaultAvatar.Close()
		if err != nil {
			log.Println(err.Error())
			jsonResponse(w, http.StatusFailedDependency, "failed to load default avatar")
			return
		}
		defaultAvatar.Close()
	}

	data.Privacy = map[bool]string{true: "public", false: "private"}[data.Public]

	onlyEnglishRegex := regexp.MustCompile(`^[a-zA-Z0-9]{2,15}$`)

	if data.Nickname != "" {
		if !onlyEnglishRegex.MatchString(data.Nickname) {
			message := `Invalid nickname: ` + data.Nickname + `
			Nickname must only contain english letters and numbers.
			Nickname must be between 2 and 15 characters long.`
			jsonResponse(w, http.StatusUnprocessableEntity, message)
			return
		}
	}

	if len(data.FirstName) < 1 || len(data.FirstName) > 32 {
		message := `Invalid first name length: ` + data.FirstName + `
		First name must be between 1 and 32 characters long`
		jsonResponse(w, http.StatusUnprocessableEntity, message)
		return
	}

	if len(data.LastName) < 1 || len(data.LastName) > 32 {
		message := `Invalid last name length: ` + data.LastName + `
		Last name must be between 1 and 32 characters long`
		jsonResponse(w, http.StatusUnprocessableEntity, message)
		return
	}

	emailRegex := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,4})+$`)

	if !emailRegex.MatchString(strings.ToLower((data.Email))) {
		message := `Invalid email: ` + data.Email + `
		Email must be a valid email address`
		jsonResponse(w, http.StatusUnprocessableEntity, message)
		return
	}

	if len(data.Password) < 6 || len(data.Password) > 15 {
		message := `Invalid password length: ` + data.Password + `
		Password must be between 6 and 15 characters long`
		jsonResponse(w, http.StatusUnprocessableEntity, message)
		return
	}

	if !onlyEnglishRegex.MatchString(data.Password) {
		message := `Invalid password: ` + data.Password + `
Password must only contain english characters and numbers`
		jsonResponse(w, http.StatusUnprocessableEntity, message)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusUnprocessableEntity, "dependency failure: could not hash password")
		return
	}
	_, err = statements["addUser"].Exec(data.Email, string(hash), data.FirstName, data.LastName, data.Dob, data.avatarBytes, data.Nickname, data.AboutMe, data.Privacy)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			log.Println(err.Error())
			jsonResponse(w, http.StatusConflict, "This email is already taken")
			return
		}

		log.Println(err.Error())
		jsonResponse(w, http.StatusUnprocessableEntity, "database entry for adding user failed")
		return
	}

	UUID, err := createSession(data.Email)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusInternalServerError, " create session failed")
		return
	}

	jsonResponseObj, _ := json.Marshal(map[string]string{
		"UUID":  UUID,
		"email": data.Email,
	})
	_, err = w.Write(jsonResponseObj)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "w.Write(jsonResponseObj)<-UUID,email failed")
		return
	}
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	defer recovery(w)
	var data loginData
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusUnprocessableEntity, "Bad request. The JSON body is not as expected")
		return
	}

	var email, hash string
	rows, err := statements["getUserCredentials"].Query(data.Email)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&email, &hash)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "scan credentials failed")
		return
	}
	rows.Close()

	if email == "" || hash == "" {
		jsonResponse(w, http.StatusUnauthorized, "Invalid credentials. Email or password is incorrect")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusUnauthorized, "Invalid credentials. Forgot password?")
		return
	}

	UUID, err := createSession(email)

	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusInternalServerError, "create session failed")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "user_uuid",
		Value:    UUID,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false,
		SameSite: 4,
		Path:     "/",
	})

	jsonResponseObj, _ := json.Marshal(map[string]string{
		"UUID":  UUID,
		"email": email,
	})
	_, err = w.Write(jsonResponseObj)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "w.Write(jsonResponseObj)<-UUID,email failed")
		return
	}
}

func sessionCheckHandler(w http.ResponseWriter, r *http.Request) {
	defer recovery(w)
	var data UUIDData
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusBadRequest, "")
		return
	}
	rows, err := statements["getSession"].Query(data.UUID)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusInternalServerError, "")
		return
	}
	defer rows.Close()
	if !rows.Next() {
		w.WriteHeader(200)
		jsonResponseObj, err := json.Marshal(map[string]bool{
			"Exists": false,
		})
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "failed")
			return
		}
		_, err = w.Write(jsonResponseObj)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, "failed")
		}
		return
	}
	rows.Close()

	jsonResponseObj, _ := json.Marshal(map[string]bool{
		"Exists": true,
	})
	_, err = w.Write(jsonResponseObj)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "failed")
		return
	}

}

func userLogoutHandler(w http.ResponseWriter, r *http.Request) {
	defer recovery(w)

	cookie, err := r.Cookie("user_uuid")
	if err != nil || cookie.Value == "" || cookie == nil {
		jsonResponse(w, http.StatusOK, "You are not logged in")
		return
	}

	uuid := cookie.Value

	_, err = statements["removeSession"].Exec(uuid)
	if err != nil {
		log.Println(err.Error())
		jsonResponse(w, http.StatusInternalServerError, "removeSession query failed")
		return
	}
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")
	jsonResponse(w, http.StatusOK, "Session deleted")
}

func createSession(email string) (UUID string, err error) {
	random, _ := uuid.NewV4()
	UUID = random.String()
	ID, err := getIDbyEmail(email)
	if err != nil {
		return "", err
	}
	_, err = statements["addSession"].Exec(UUID, ID)
	if err != nil {
		return "", err
	}
	return UUID, nil
}

func getIDbyEmail(email string) (ID int, err error) {
	rows, err := statements["getUserIDByEmail"].Query(email)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&ID)
	if err != nil {
		return 0, err
	}
	rows.Close()
	return ID, nil
}

func getUserEmailbyID(ID int) (email string, err error) {
	rows, err := statements["getEmailByID"].Query(ID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&email)
	if err != nil {
		return "", err
	}
	rows.Close()
	return email, nil
}

func isImage(data []byte) bool {
	if len(data) < 4 {
		return false
	}

	if data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF {
		return true // JPEG
	}

	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return true // PNG
	}

	if data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x38 {
		return true // GIF
	}

	return false
}

// todo: CHECK IT! , it is refactored to prevent warning
func randomNum(min, max int) int {
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	rng.Seed(time.Now().Unix())
	return rng.Intn(max-min) + min
}

func doesUserExist(ID int) bool {
	rows, err := statements["getUserbyID"].Query(ID)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}

func canISeeThisProfile(myID int, theirID int) bool {
	rows, err := statements["getUserPrivacy"].Query(theirID)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer rows.Close()
	rows.Next()
	var privacy string
	err = rows.Scan(&privacy)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	rows.Close()
	if privacy == "public" {
		return true
	}
	if privacy == "private" {
		iAmFollower, err := isFollowing(theirID, myID)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		if iAmFollower {
			return true
		}
	}
	return false
}

// Min user info for card display
type MinUserInfo struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type FullProfileDTOtoFrontend struct {
	ProfileDTOtoFrontend
	Posts     []Post
	Followers []MinUserInfo
	Following []MinUserInfo
}

func giveMeMyFullProfile(myID int) FullProfileDTOtoFrontend {
	// ******** ProfileDTOtoFrontend ********

	// statement to be used getUserProfile
	var profile ProfileDTOtoFrontend
	rows, err := statements["getUserProfile"].Query(myID)
	if err != nil {
		log.Println(err.Error())
		return FullProfileDTOtoFrontend{}
	}
	defer rows.Close()
	rows.Next()
	avatarBlob := []byte{}
	err = rows.Scan(&profile.ID, &profile.Email, &profile.FirstName, &profile.LastName,
		&profile.Dob, &avatarBlob, &profile.Nickname, &profile.AboutMe, &profile.Privacy)
	if err != nil {
		log.Println(err.Error())
		return FullProfileDTOtoFrontend{}
	}
	rows.Close()
	profile.Avatar = base64.StdEncoding.EncodeToString(avatarBlob)

	// ******** Posts ********

	// statement to be used getPostsByUserID
	var posts []Post
	rows, err = statements["getPostsByUserID"].Query(myID)
	if err != nil {
		log.Println(err.Error())
		return FullProfileDTOtoFrontend{}
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		var postPicture []byte
		var firstName, lastName, email string
		err = rows.Scan(&post.PostID, &post.Title, &post.Content, &postPicture,
			&firstName, &lastName, &post.Privacy, &email, &post.CreatedAt)
		if err != nil {
			log.Println(err.Error())
			return FullProfileDTOtoFrontend{}
		}
		post.CreatorID = myID
		post.Picture = base64.StdEncoding.EncodeToString(postPicture)
		posts = append(posts, post)
	}

	// ******** Followers ********

	// statement to be used getFollowers
	var followers []MinUserInfo
	rows, err = statements["getFollowers"].Query(myID)
	if err != nil {
		log.Println(err.Error())
		return FullProfileDTOtoFrontend{}
	}
	defer rows.Close()
	for rows.Next() {
		var follower MinUserInfo
		err = rows.Scan(&follower.ID)
		if err != nil {
			log.Println(err.Error())
			return FullProfileDTOtoFrontend{}
		}
		followers = append(followers, follower)
	}
	// for each follower, get their avatar, firstname, lastname, nickname
	for i, follower := range followers {
		var avatarBlob []byte
		rows, err = statements["getUserProfile"].Query(follower.ID)
		if err != nil {
			log.Println(err.Error())
			return FullProfileDTOtoFrontend{}
		}
		defer rows.Close()
		rows.Next()
		var id int
		var email, dob, aboutMe, privacy string
		err = rows.Scan(&id, &email, &follower.FirstName, &follower.LastName, &dob, &avatarBlob, &follower.Nickname, &aboutMe, &privacy)
		if err != nil {
			log.Println(err.Error())
			return FullProfileDTOtoFrontend{}
		}
		follower.Avatar = base64.StdEncoding.EncodeToString(avatarBlob)
		followers[i] = follower
	}

	// ******** Following ********

	// statement to be used getFollowing
	var following []MinUserInfo
	rows, err = statements["getFollowing"].Query(myID)
	if err != nil {
		log.Println(err.Error())
		return FullProfileDTOtoFrontend{}
	}
	defer rows.Close()
	for rows.Next() {
		var follow MinUserInfo
		err = rows.Scan(&follow.ID)
		if err != nil {
			log.Println(err.Error())
			return FullProfileDTOtoFrontend{}
		}
		following = append(following, follow)
	}

	// for each following, get their avatar, firstname, lastname, nickname
	for i, follow := range following {
		var avatarBlob []byte
		rows, err = statements["getUserProfile"].Query(follow.ID)
		if err != nil {
			log.Println(err.Error())
			return FullProfileDTOtoFrontend{}
		}
		defer rows.Close()
		rows.Next()
		var id int
		var email, dob, aboutMe, privacy string
		err = rows.Scan(&id, &email, &follow.FirstName, &follow.LastName, &dob, &avatarBlob, &follow.Nickname, &aboutMe, &privacy)
		if err != nil {
			log.Println(err.Error())
			return FullProfileDTOtoFrontend{}
		}
		follow.Avatar = base64.StdEncoding.EncodeToString(avatarBlob)
		following[i] = follow
	}

	// ******** FullProfileDTOtoFrontend ********

	var fullProfile FullProfileDTOtoFrontend
	fullProfile.ProfileDTOtoFrontend = profile
	fullProfile.Posts = posts
	fullProfile.Followers = followers
	fullProfile.Following = following

	return fullProfile
}

func giveMeOtherPersonFullProfile(myID, idofUserToSee int) FullProfileDTOtoFrontend {

	// ******** GetFullProfileOfOtherPerson ********

	var fullProfile = giveMeMyFullProfile(idofUserToSee)

	// ******** FilterPostsThatICanSee ********

	var filteredPosts []Post
	for i, post := range fullProfile.Posts {
		if post.Privacy == "public" {
			filteredPosts = append(filteredPosts, post)
		} else if post.Privacy == "private" {
			iAmFollower, err := isFollowing(idofUserToSee, myID)
			if err != nil {
				log.Println(err.Error())
				return FullProfileDTOtoFrontend{}
			}
			if iAmFollower {
				filteredPosts = append(filteredPosts, post)
			} else {
				fullProfile.Posts[i].Title = "Private Post"
				fullProfile.Posts[i].Content = "You are not allowed to see this post"
				fullProfile.Posts[i].Picture = ""
			}
		} else if post.Privacy == "almost private" {
			rows, err := statements["getAlmostPrivate"].Query(myID, post.PostID)
			if err != nil {
				log.Println(err.Error())
				return FullProfileDTOtoFrontend{}
			}
			defer rows.Close()
			if rows.Next() {
				filteredPosts = append(filteredPosts, post)
			} else {
				fullProfile.Posts[i].Title = "Private Post"
				fullProfile.Posts[i].Content = "You are not allowed to see this post"
				fullProfile.Posts[i].Picture = ""
			}
		}
	}

	return fullProfile
}
