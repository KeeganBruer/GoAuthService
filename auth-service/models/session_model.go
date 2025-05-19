package models

import (
	"fmt"
	"go-auth-service/services/jwttokens"
	"time"
)

const TOKEN_TIME = 30
const REFRESH_TIME = 60

type NewSession struct {
	UserID int
}
type Session struct {
	ID         int
	UserID     int       `json:"user_id"`
	Experation time.Time `json:"experation"`
}

func GetSessionByUserID(userID int) (*Session, error) {
	db_conn := GetDBConnection()
	q := db_conn.GetTable("sessions").NewSelect()
	q.Where(fmt.Sprintf("user_id = %d", userID))

	var dateTime string

	existing := &Session{}

	err := q.FindOne(
		&existing.ID,
		&existing.UserID,
		&dateTime,
	)
	if err != nil {
		return nil, err
	}
	exp, err := time.Parse(time.DateTime, dateTime)
	if err != nil {
		return nil, err
	}
	existing.Experation = exp
	fmt.Println("dateTime: " + existing.Experation.Format(time.Kitchen))
	return existing, err
}
func CreateOrGetSession(data *NewSession) *Session {
	exp := time.Now().Add(time.Duration(REFRESH_TIME) * time.Minute)
	session := &Session{
		UserID:     data.UserID,
		Experation: exp,
	}

	existing, err := GetSessionByUserID(session.UserID)

	if err != nil {
		session.Save()
	} else {
		session = existing
	}
	return session
}

type SessionTokens struct {
	Token   string
	Refresh string
}

func (session *Session) GetTokens() (*SessionTokens, error) {
	//Create a pair of JWT tokens with different expirations
	token, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		SessionID:     session.ID,
		MinutesTilExp: TOKEN_TIME,
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		SessionID:     session.ID,
		MinutesTilExp: REFRESH_TIME,
	})
	if err != nil {
		return nil, err
	}

	//update session experation
	exp := time.Now().Add(time.Duration(REFRESH_TIME) * time.Minute)
	session.Experation = exp
	session.Save()

	return &SessionTokens{
		Token:   token,
		Refresh: refreshToken,
	}, nil
}
func (session *Session) Save() error {
	db_conn := GetDBConnection()
	new := db_conn.GetTable("sessions").NewInsert()
	new.AddIntColumn("id", session.ID)
	new.AddIntColumn("user_id", session.UserID)
	new.AddDateTimeColumn("experation", session.Experation)
	new.Send()
	if session.ID == 0 {
		//Load the ID for the inserted session
		q := db_conn.GetTable("sessions").NewSelect()
		q.Where(fmt.Sprintf("user_id = %d", session.UserID))
		var dateTime string
		err := q.FindOne(
			&session.ID,
			&session.UserID,
			&dateTime,
		)
		if err != nil {
			return err
		}
		exp, err := time.Parse(time.DateTime, dateTime)
		if err != nil {
			return err
		}
		session.Experation = exp
	}
	return nil
}

func GetSessionFromToken(token *jwttokens.TokenData) (*Session, error) {
	db_conn := GetDBConnection()
	q := db_conn.GetTable("sessions").NewSelect()
	q.Where(fmt.Sprintf("id = %d", token.SessionID))
	var dateTime string
	session := &Session{}
	err := q.FindOne(
		&session.ID,
		&session.UserID,
		&dateTime,
	)
	if err != nil {
		return nil, err
	}
	exp, err := time.Parse(time.DateTime, dateTime)
	if err != nil {
		return nil, err
	}
	session.Experation = exp
	return session, nil
}
