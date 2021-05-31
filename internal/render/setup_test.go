// whatever is written in here will run before our tests run

package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/andkolbe/bed-and-breakfast/internal/config"
	"github.com/andkolbe/bed-and-breakfast/internal/models"
)


var session *scs.SessionManager
var testApp config.AppConfig // create a copy of the app variable in render

func TestMain(m *testing.M) {

	// what I am going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// print these to the terminal
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // true = cookie persists even if the browser window closes
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp // makes sure that app (defined inside of render.go) is populated with the testApp data

	os.Exit(m.Run())
}

// create an interface that satisfies the requirements for a response writer
type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw  *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}