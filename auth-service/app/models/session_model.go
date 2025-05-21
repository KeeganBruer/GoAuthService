package models

import (
	"errors"
	"fmt"
	"go-auth-service/app/services/jwttokens"
	"math"
	"math/rand"
	"sqlquerybuilder"
	"time"
)

const TOKEN_TIME = 30
const REFRESH_TIME = 60

type SessionModel struct {
	BaseModel
}

func (models *Models) GetSessionModel() *SessionModel {
	return &SessionModel{
		BaseModel{
			models: models,
		},
	}
}

type NewSession struct {
	UserID int
}
type Session struct {
	SessionModel
	ID         int
	RefreshID  int       `json:"refresh_id"`
	UserID     int       `json:"user_id"`
	Experation time.Time `json:"experation"`
}

func (sessionModel *SessionModel) GetSessionByRefreshID(RefreshID int) (*Session, error) {
	builder := sessionModel.GetDBQueryBuilder()
	q := builder.GetTable("sessions").NewSelect()
	q.Where(fmt.Sprintf("refresh_id = %d", RefreshID))
	return sessionModel.GetSessionByQuery(q)
}
func (sessionModel *SessionModel) GetSessionByID(ID int) (*Session, error) {
	builder := sessionModel.GetDBQueryBuilder()
	q := builder.GetTable("sessions").NewSelect()
	q.Where(fmt.Sprintf("id = %d", ID))
	return sessionModel.GetSessionByQuery(q)
}
func (sessionModel *SessionModel) GetSessionByUserID(userID int) (*Session, error) {
	builder := sessionModel.GetDBQueryBuilder()
	q := builder.GetTable("sessions").NewSelect()
	q.Where(fmt.Sprintf("user_id = %d", userID))
	return sessionModel.GetSessionByQuery(q)
}
func (sessionModel *SessionModel) GetSessionByQuery(q *sqlquerybuilder.SQLQuery) (*Session, error) {
	var dateTime string

	existing := &Session{
		SessionModel: SessionModel{
			BaseModel: BaseModel{
				models: sessionModel.models,
			},
		},
	}

	err := q.FindOne(
		&existing.ID,
		&existing.RefreshID,
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
	return existing, err
}

func (sessionModel *SessionModel) CreateSession(data *NewSession) *Session {
	exp := time.Now().Add(time.Duration(REFRESH_TIME) * time.Minute)
	session := &Session{
		SessionModel: SessionModel{
			BaseModel: BaseModel{
				models: sessionModel.models,
			},
		},
		UserID:     data.UserID,
		Experation: exp,
	}
	session.Save()
	return session
}

type SessionTokens struct {
	Token   string
	Refresh string
}

func (session *Session) GetTokens() (*SessionTokens, error) {
	//Create a pair of JWT tokens with different expirations
	token, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		ID:            session.ID,
		Type:          "session",
		MinutesTilExp: TOKEN_TIME,
	})
	if err != nil {
		return nil, err
	}

	//Generate New RefreshID
	RefreshID := rand.Intn(2000000000)
	builder := session.GetDBQueryBuilder()
	sel := builder.GetTable("sessions").NewSelect()
	sel.Where(fmt.Sprintf("refresh_id = %d", RefreshID))
	for sel.Exists() {
		RefreshID = rand.Intn(math.MaxInt)
		sel.Where(fmt.Sprintf("refresh_id = %d", RefreshID))
	}
	session.RefreshID = RefreshID

	refreshToken, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		ID:            session.RefreshID,
		Type:          "refresh_token",
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
	builder := session.GetDBQueryBuilder()

	new := builder.GetTable("sessions").NewInsert()
	new.AddColumn("id", builder.Int2DB(session.ID))
	new.AddColumn("refresh_id", builder.Int2DB(session.RefreshID))
	new.AddColumn("user_id", builder.Int2DB(session.UserID))
	new.AddColumn("experation", builder.Date2DB(session.Experation))
	new.Send()
	if session.ID == 0 {
		//Load the ID for the inserted session
		q := builder.GetTable("sessions").NewSelect()
		q.Where(fmt.Sprintf(
			"user_id = %s AND refresh_id = %s AND experation = %s",
			builder.Int2DB(session.UserID),
			builder.Int2DB(session.RefreshID),
			builder.Date2DB(session.Experation),
		))
		LoadedSess, err := session.GetSessionByQuery(q)
		if err != nil {
			return nil
		}
		session.ID = LoadedSess.ID
	}
	return nil
}

func (sessionModel *SessionModel) GetSessionFromToken(token *jwttokens.TokenData) (*Session, error) {
	if token.Type != "session" {
		return nil, errors.New("token is not a session jwt")
	}
	builder := sessionModel.GetDBQueryBuilder()
	q := builder.GetTable("sessions").NewSelect()
	q.Where(fmt.Sprintf("id = %d", token.ID))
	var dateTime string
	session := &Session{
		SessionModel: SessionModel{
			BaseModel: BaseModel{
				models: sessionModel.models,
			},
		},
	}
	err := q.FindOne(
		&session.ID,
		&session.RefreshID,
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
