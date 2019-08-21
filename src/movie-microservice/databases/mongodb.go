package databases

import (
	"time"

	"github.com/gillepool/MovieBackend/src/movie-microservice/utils"
	mgo "gopkg.in/mgo.v2"
)

//MongoDB info
type MongoDB struct {
	MgDbSession  *mgo.Session
	Databasename string
}

//Init the MogoDB
func (db *MongoDB) Init() error {
	db.Databasename = utils.Config.MgDbName

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{utils.Config.MgAddrs}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: db.Databasename,           // Database name
		Username: utils.Config.MgDbUsername, // Username
		Password: utils.Config.MgDbPassword, // Password
	}
	var err error
	db.MgDbSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}

	return nil
}

// Close session unless its already to closed
func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Close()
	}
}
