package handlers

import (
	// "context"
	// "encoding/json"
	// "fmt"
	// "log"
	"net/http"
	"net/http/httptest"
	// "net/url"
	// "reflect"
	// "strings"
	"testing"
	// "time"

	// "github.com/andkolbe/bed-and-breakfast/internal/driver"
	// "github.com/andkolbe/bed-and-breakfast/internal/models"
)

//
type postData struct {
	key   string // name of the form input
	value string // what is typed in the form input
}

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"r1", "/room1", "GET", http.StatusOK},
	{"r2", "/room2", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	{"non-existent", "/this/does/not/exist", "GET", http.StatusNotFound},
	// new routes
	{"login", "/user/login", "GET", http.StatusOK},
	{"logout", "/user/logout", "GET", http.StatusOK},
	{"dashboard", "/admin/dashboard", "GET", http.StatusOK},
	{"new res", "/admin/reservations-new", "GET", http.StatusOK},
	{"all res", "/admin/reservations-all", "GET", http.StatusOK},
	{"show res", "/admin/reservations/new/1/show", "GET", http.StatusOK},
	{"show res cal", "/admin/reservations-calendar", "GET", http.StatusOK},
	{"show res cal with params", "/admin/reservations-calendar?y=2020&m=1", "GET", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes) // create a test server and store it in a variable
	defer ts.Close() // what is written after defer doesn't get executed until after the current function is finished

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url) // tell the client to make a GET request to the url that we want to test
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

// func TestRepository_Reservation(t *testing.T) {
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/make-reservation", nil)
// 	// we have to put our reservation variable into the session of the request. Use context to do this
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	// response recorder simulates the req/res lifecycle
// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	// call our reservation handler
// 	handler := http.HandlerFunc(Repo.Reservation)

// 	// serve http on our handler just like we do on any of our routes
// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}

// 	// test case where reservation is not in session (reset everything)
// 	req, _ = http.NewRequest("GET", "/make-reservation", nil)
// 	// still need to get the context with the session header
// 	// we need to be able to test the situation where we can't find the value in the session because there is no session
// 	ctx = getCtx(req)
// 	// we can put the context back into the request
// 	req = req.WithContext(ctx)
// 	rr = httptest.NewRecorder()

// 	handler.ServeHTTP(rr, req)
// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test with non existent room
// 	req, _ = http.NewRequest("GET", "/make-reservation", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	rr = httptest.NewRecorder()
// 	reservation.RoomID = 100 // override room id with non existent room id for this test
// 	session.Put(ctx, "reservation", reservation)

// 	handler.ServeHTTP(rr, req)
// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// }

// func TestRepository_PostReservation(t *testing.T) {
// 	// build up the request body as a string
// 	// reqBody := "start_date=2050-01-01"
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Andrew")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Kolbe")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=email@andrew.com")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=1111111111")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	// another way to write this
// 	postedData := url.Values{}
// 	postedData.Add("start_date", "2050-01-01")
// 	postedData.Add("end_date", "2050-01-02")
// 	postedData.Add("first_name", "Andrew")
// 	postedData.Add("last_name", "Kolbe")
// 	postedData.Add("email", "email@andrew.com")
// 	postedData.Add("phone", "111-111-111")
// 	postedData.Add("room_id", "1")

// 	// create request
// 	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
// 	// get context with session
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr := httptest.NewRecorder()

// 	handler := http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("PostReservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for missing post body
// 	req, _ = http.NewRequest("POST", "/make-reservation", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for missing post body: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for invalid start date
// 	reqBody := "start_date=invalid"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Andrew")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Kolbe")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=email@andrew.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=1111111111")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid start date: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for invalid end date
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=invalid")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Andrew")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Kolbe")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=email@andrew.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=1111111111")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid end date: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for invalid room id
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Andrew")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Kolbe")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=email@andrew.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=1111111111")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=invalid")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid room id: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for invalid data
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=A") // first name must be 3 characters long
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Kolbe")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=email@andrew.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=1111111111")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid data: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for failure to insert reservation into database
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Andrew")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Kolbe")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=email@andrew.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=1111111111")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=2")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler failed when trying to fail inserting reservation: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for failure to insert restriction into database
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Andrew")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Kolbe")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=email@andrew.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=1111111111")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1000")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler failed when trying to fail inserting reservation: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// }

// func TestNewRepo(t *testing.T) {
// 	var db driver.DB
// 	testRepo := NewRepo(&app, &db)

// 	if reflect.TypeOf(testRepo).String() != "*handlers.Repository" {
// 		t.Errorf("Did not get correct type from NewRepo: got %s, wanted *Repository", reflect.TypeOf(testRepo).String())
// 	}
// }

// func TestRepository_PostAvailability(t *testing.T) {
// 	/*****************************************
// 	// first case -- rooms are not available
// 	*****************************************/
// 	// create our request body
// 	reqBody := "start=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2050-01-02")

// 	// create our request
// 	req, _ := http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr := httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler := http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have no rooms available, we expect to get status http.StatusSeeOther
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Post availability when no rooms available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	/*****************************************
// 	// second case -- rooms are available
// 	*****************************************/
// 	// this time, we specify a start date before 2040-01-01, which will give us
// 	// a non-empty slice, indicating that rooms are available
// 	reqBody = "start=2040-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2040-01-02")

// 	// create our request
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusOK
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Post availability when rooms are available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}

// 	/*****************************************
// 	// third case -- empty post body
// 	*****************************************/
// 	// create our request with a nil body, so parsing form fails
// 	req, _ = http.NewRequest("POST", "/search-availability", nil)

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusTemporaryRedirect
// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("Post availability with empty request body (nil) gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	/*****************************************
// 	// fourth case -- start date in wrong format
// 	*****************************************/
// 	// this time, we specify a start date in the wrong format
// 	reqBody = "start=invalid"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2040-01-02")
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusTemporaryRedirect
// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("Post availability with invalid start date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	/*****************************************
// 	// fifth case -- end date in wrong format
// 	*****************************************/
// 	// this time, we specify a start date in the wrong format
// 	reqBody = "start=2040-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "invalid")
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusTemporaryRedirect
// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("Post availability with invalid end date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	/*****************************************
// 	// sixth case -- database query fails
// 	*****************************************/
// 	// this time, we specify a start date of 2060-01-01, which will cause
// 	// our testdb repo to return an error
// 	reqBody = "start=2060-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2060-01-02")
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusTemporaryRedirect
// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("Post availability when database query fails gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}
// }

// func TestRepository_AvailabilityJSON(t *testing.T) {
// 	// first case - rooms are not available
// 	reqBody := "start=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	// tells the server to expect a POST request
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	handler := http.HandlerFunc(Repo.PostReservation)

// 	rr := httptest.NewRecorder()

// 	handler.ServeHTTP(rr, req)

// 	var j jsonResponse
// 	err := json.Unmarshal([]byte(rr.Body.String()), &j)
// 	if err != nil {
// 		t.Error("failed to parse JSON")
// 	}
// }

// func TestRepository_ReservationSummary(t *testing.T) {
// 	/*****************************************
// 	// first case -- reservation in session
// 	*****************************************/
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/reservation-summary", nil)
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.ReservationSummary)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusOK {
// 		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}

// 	/*****************************************
// 	// second case -- reservation not in session
// 	*****************************************/
// 	req, _ = http.NewRequest("GET", "/reservation-summary", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.ReservationSummary)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}
// }

// func TestRepository_ChooseRoom(t *testing.T) {
// 	/*****************************************
// 	// first case -- reservation in session
// 	*****************************************/
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/choose-room/1", nil)
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)
// 	// set the RequestURI on the request so that we can grab the ID
// 	// from the URL
// 	req.RequestURI = "/choose-room/1"

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.ChooseRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	///*****************************************
// 	//// second case -- reservation not in session
// 	//*****************************************/
// 	req, _ = http.NewRequest("GET", "/choose-room/1", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.RequestURI = "/choose-room/1"

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.ChooseRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	///*****************************************
// 	//// third case -- missing url parameter, or malformed parameter
// 	//*****************************************/
// 	req, _ = http.NewRequest("GET", "/choose-room/fish", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.RequestURI = "/choose-room/fish"

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.ChooseRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}
// }

// func TestRepository_BookRoom(t *testing.T) {
// 	/*****************************************
// 	// first case -- database works
// 	*****************************************/
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/book-room?s=2050-01-01&e=2050-01-02&id=1", nil)
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.BookRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	/*****************************************
// 	// second case -- database failed
// 	*****************************************/
// 	req, _ = http.NewRequest("GET", "/book-room?s=2040-01-01&e=2040-01-02&id=4", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.BookRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}
// }

// // loginTests is the data for the Login handler tests
// var loginTests = []struct {
// 	name               string
// 	email              string
// 	expectedStatusCode int
// 	expectedHTML       string
// 	expectedLocation   string
// }{
// 	{
// 		"valid-credentials",
// 		"email@email.com",
// 		http.StatusSeeOther, // status code for a redirect
// 		"", // can't examine html easily on a redirect from a test
// 		"/", // redirect to home page after succesfully logging in
// 	},
// 	{
// 		"invalid-credentials",
// 		"jack@nimble.com", // invalid
// 		http.StatusSeeOther,
// 		"",
// 		"/user/login",
// 	},
// 	{
// 		"invalid-data",
// 		"j",
// 		http.StatusOK,
// 		`action="/user/login"`,
// 		"",
// 	},
// }

// func TestLogin(t *testing.T) {
// 	// range through all tests
// 	for _, e := range loginTests {
// 		postedData := url.Values{}
// 		postedData.Add("email", e.email)
// 		postedData.Add("password", "password")

// 		// create request
// 		req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(postedData.Encode()))
// 		ctx := getCtx(req)
// 		req = req.WithContext(ctx)

// 		// set the header
// 		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 		rr := httptest.NewRecorder()

// 		// call the handler
// 		handler := http.HandlerFunc(Repo.PostShowLogin)
// 		handler.ServeHTTP(rr, req)

// 		if rr.Code != e.expectedStatusCode {
// 			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
// 		}

// 		if e.expectedLocation != "" {
// 			// get the URL from test
// 			actualLoc, _ := rr.Result().Location()
// 			if actualLoc.String() != e.expectedLocation {
// 				t.Errorf("failed %s: expected location %s, but got location %s", e.name, e.expectedLocation, actualLoc.String())
// 			}
// 		}

// 		// checking for expected values in HTML
// 		if e.expectedHTML != "" {
// 			// read the response body into a string
// 			html := rr.Body.String()
// 			if !strings.Contains(html, e.expectedHTML) {
// 				t.Errorf("failed %s: expected to find %s but did not", e.name, e.expectedHTML)
// 			}
// 		}
// 	}
// }

// var adminPostShowReservationTests = []struct {
// 	name                 string
// 	url                  string
// 	postedData           url.Values
// 	expectedResponseCode int
// 	expectedLocation     string
// 	expectedHTML         string
// }{
// 	{
// 		name: "valid-data-from-new",
// 		url:  "/admin/reservations/new/1/show",
// 		postedData: url.Values{
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedLocation:     "/admin/reservations-new",
// 		expectedHTML:         "",
// 	},
// 	{
// 		name: "valid-data-from-all",
// 		url:  "/admin/reservations/all/1/show",
// 		postedData: url.Values{
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedLocation:     "/admin/reservations-all",
// 		expectedHTML:         "",
// 	},
// 	{
// 		name: "valid-data-from-cal",
// 		url:  "/admin/reservations/cal/1/show",
// 		postedData: url.Values{
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"year":       {"2022"},
// 			"month":      {"01"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedLocation:     "/admin/reservations-calendar?y=2022&m=01",
// 		expectedHTML:         "",
// 	},
// }

// // TestAdminPostShowReservation tests the AdminPostReservation handler
// func TestAdminPostShowReservation(t *testing.T) {
// 	for _, e := range adminPostShowReservationTests {
// 		var req *http.Request
// 		if e.postedData != nil {
// 			req, _ = http.NewRequest("POST", "/user/login", strings.NewReader(e.postedData.Encode()))
// 		} else {
// 			req, _ = http.NewRequest("POST", "/user/login", nil)
// 		}
// 		ctx := getCtx(req)
// 		req = req.WithContext(ctx)
// 		req.RequestURI = e.url

// 		// set the header
// 		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 		rr := httptest.NewRecorder()

// 		// call the handler
// 		handler := http.HandlerFunc(Repo.AdminPostShowReservation)
// 		handler.ServeHTTP(rr, req)

// 		if rr.Code != e.expectedResponseCode {
// 			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedResponseCode, rr.Code)
// 		}

// 		if e.expectedLocation != "" {
// 			// get the URL from test
// 			actualLoc, _ := rr.Result().Location()
// 			if actualLoc.String() != e.expectedLocation {
// 				t.Errorf("failed %s: expected location %s, but got location %s", e.name, e.expectedLocation, actualLoc.String())
// 			}
// 		}

// 		// checking for expected values in HTML
// 		if e.expectedHTML != "" {
// 			// read the response body into a string
// 			html := rr.Body.String()
// 			if !strings.Contains(html, e.expectedHTML) {
// 				t.Errorf("failed %s: expected to find %s but did not", e.name, e.expectedHTML)
// 			}
// 		}
// 	}
// }

// var adminPostReservationCalendarTests = []struct {
// 	name                 string
// 	postedData           url.Values
// 	expectedResponseCode int
// 	expectedLocation     string
// 	expectedHTML         string
// 	blocks               int
// 	reservations         int
// }{
// 	{
// 		name: "cal",
// 		postedData: url.Values{
// 			"year":  {time.Now().Format("2006")},
// 			"month": {time.Now().Format("01")},
// 			fmt.Sprintf("add_block_1_%s", time.Now().AddDate(0, 0, 2).Format("2006-01-2")): {"1"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 	},
// 	{
// 		name:                 "cal-blocks",
// 		postedData:           url.Values{},
// 		expectedResponseCode: http.StatusSeeOther,
// 		blocks:               1,
// 	},
// 	{
// 		name:                 "cal-res",
// 		postedData:           url.Values{},
// 		expectedResponseCode: http.StatusSeeOther,
// 		reservations:         1,
// 	},
// }

// func TestPostReservationCalendar(t *testing.T) {
// 	for _, e := range adminPostReservationCalendarTests {
// 		var req *http.Request
// 		if e.postedData != nil {
// 			req, _ = http.NewRequest("POST", "/admin/reservations-calendar", strings.NewReader(e.postedData.Encode()))
// 		} else {
// 			req, _ = http.NewRequest("POST", "/admin/reservations-calendar", nil)
// 		}
// 		ctx := getCtx(req)
// 		req = req.WithContext(ctx)

// 		now := time.Now()
// 		bm := make(map[string]int)
// 		rm := make(map[string]int)

// 		currentYear, currentMonth, _ := now.Date()
// 		currentLocation := now.Location()

// 		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
// 		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

// 		for d := firstOfMonth; d.After(lastOfMonth) == false; d = d.AddDate(0, 0, 1) {
// 			rm[d.Format("2006-01-2")] = 0
// 			bm[d.Format("2006-01-2")] = 0
// 		}

// 		if e.blocks > 0 {
// 			bm[firstOfMonth.Format("2006-01-2")] = e.blocks
// 		}

// 		if e.reservations > 0 {
// 			rm[lastOfMonth.Format("2006-01-2")] = e.reservations
// 		}

// 		session.Put(ctx, "block_map_1", bm)
// 		session.Put(ctx, "reservation_map_1", rm)

// 		// set the header
// 		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 		rr := httptest.NewRecorder()

// 		// call the handler
// 		handler := http.HandlerFunc(Repo.AdminPostReservationsCalendar)
// 		handler.ServeHTTP(rr, req)

// 		if rr.Code != e.expectedResponseCode {
// 			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedResponseCode, rr.Code)
// 		}

// 	}
// }

// var adminProcessReservationTests = []struct {
// 	name                 string
// 	queryParams          string
// 	expectedResponseCode int
// 	expectedLocation     string
// }{
// 	{
// 		name:                 "process-reservation",
// 		queryParams:          "",
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedLocation:     "",
// 	},
// 	{
// 		name:                 "process-reservation-back-to-cal",
// 		queryParams:          "?y=2021&m=12",
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedLocation:     "",
// 	},
// }

// func TestAdminProcessReservation(t *testing.T) {
// 	for _, e := range adminProcessReservationTests {
// 		req, _ := http.NewRequest("GET", fmt.Sprintf("/admin/process-reservation/cal/1/do%s", e.queryParams), nil)
// 		ctx := getCtx(req)
// 		req = req.WithContext(ctx)

// 		rr := httptest.NewRecorder()

// 		handler := http.HandlerFunc(Repo.AdminProcessReservation)
// 		handler.ServeHTTP(rr, req)

// 		if rr.Code != http.StatusSeeOther {
// 			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedResponseCode, rr.Code)
// 		}
// 	}
// }

// var adminDeleteReservationTests = []struct {
// 	name                 string
// 	queryParams          string
// 	expectedResponseCode int
// 	expectedLocation     string
// }{
// 	{
// 		name:                 "delete-reservation",
// 		queryParams:          "",
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedLocation:     "",
// 	},
// 	{
// 		name:                 "delete-reservation-back-to-cal",
// 		queryParams:          "?y=2021&m=12",
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedLocation:     "",
// 	},
// }

// func TestAdminDeleteReservation(t *testing.T) {
// 	for _, e := range adminDeleteReservationTests {
// 		req, _ := http.NewRequest("GET", fmt.Sprintf("/admin/process-reservation/cal/1/do%s", e.queryParams), nil)
// 		ctx := getCtx(req)
// 		req = req.WithContext(ctx)

// 		rr := httptest.NewRecorder()

// 		handler := http.HandlerFunc(Repo.AdminDeleteReservation)
// 		handler.ServeHTTP(rr, req)

// 		if rr.Code != http.StatusSeeOther {
// 			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedResponseCode, rr.Code)
// 		}
// 	}
// }

// func getCtx(req *http.Request) context.Context {
// 	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return ctx
// }
