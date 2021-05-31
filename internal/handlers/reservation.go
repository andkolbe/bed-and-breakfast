package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andkolbe/bed-and-breakfast/internal/forms"
	"github.com/andkolbe/bed-and-breakfast/internal/helpers"
	"github.com/andkolbe/bed-and-breakfast/internal/models"
	"github.com/andkolbe/bed-and-breakfast/internal/render"
)

// reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// get the room id off of the session and check if it is valid
	room, err := m.DB.GetRoomByID(res.RoomID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't find room")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.Room.RoomName = room.RoomName

	m.App.Session.Put(r.Context(), "reservation", res)

	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	// store in a data variable
	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
		// include empty form, data, and string map when the page loads up for the first time
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// POST reservation page handler
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	// we start by pulling our reservation from the session
	// our reservation already has start date, end date, room id, and room name
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("can't get from session"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// update reservation model
	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Phone = r.Form.Get("phone")
	reservation.Email = r.Form.Get("email")

	// create a new form
	form := forms.New(r.PostForm) // PostForm has all of the url values and their associated data

	// check if the incoming form has all of the fields filled out
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	// if one of the fields on the form is not valid, repopulate the form with the data they entered and display the error message where it needs to be
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation // store the reservation in the data map (first name, last name, phone, email)

		// rerender the form with the information the user filled in
		render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// save data to db and get a reservation back
	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't insert reservation into db")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't insert room restriction")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// send notifications - to guest
	htmlMessage := fmt.Sprintf(`
		<strong>Reservation Confirmation</strong><br>
		Dear %s, <br>
		This is to confirm your reservation from %s to %s.
	`, reservation.FirstName, reservation.StartDate.Format("2006-01-02"), reservation.EndDate.Format("2006-01-02"))

	msg := models.MailData{
		To:       reservation.Email,
		From:     "bnbbooking@gmail.com",
		Subject:  "Reservation Confirmation",
		Content:  htmlMessage,
		Template: "basic.html",
	}
	// pass the msg into the app.MailChan channel
	m.App.MailChan <- msg

	// send notifications - to property owner
	htmlMessage = fmt.Sprintf(`
		<strong>Reservation Notification</strong><br>
		A reservation has been made for %s from %s to %s.
	`, reservation.Room.RoomName, reservation.StartDate.Format("2006-01-02"), reservation.EndDate.Format("2006-01-02"))

	msg = models.MailData{
		To:      "bnbbooking@gmail.com",
		From:    "bnbbooking@gmail.com",
		Subject: "Reservation Notification",
		Content: htmlMessage,
	}
	// pass the msg into the app.MailChan channel
	m.App.MailChan <- msg

	// put the reservation data in the session
	m.App.Session.Put(r.Context(), "reservation", reservation)
	// redirect the user to a different page after submitting the form so they can't click the submit button twice
	// StatusSeeOther is response code 303
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// displays the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	// get the reservation data out of the session so we can display it on the screen
	// we need to type assert the reservation data to the type of models.Reservation
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation) 
	// if it finds something called reservation in the session and it manages to assert it to type models.Reservation, ok will be true
	if !ok {
		m.App.ErrorLog.Println("Can't get error from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther) // redirect them to the home page
		return
	}

	// take the reservation out of the session and pass it as template data
	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	// format start and end dates from strings to time
	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	render.Template(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}
