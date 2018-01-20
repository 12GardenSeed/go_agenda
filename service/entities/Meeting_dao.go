package entities

import (
    "crypto/md5"
    "fmt"
    "io"
    "strings"
)

type meetingDao DataAccessObject

func (dao *meetingDao) Insert(meeting *Meeting) error {
    sqlStmt := `
        INSERT INTO　meetings（owner, title, members, starttime, endtime) VALUES(
            '` + meeting.Owner + `',
            '` + meeting.Title + `',
            '` + strings.Join(meeting.Members, "&") + `',
            '` + meeting.Starttime + `',
            '` + meeting.Endtime + `'
    );`
    result, err := db.Exec(sqlStmt)
    CheckErr(err)
    id, err := result.LastInsertId()
    CheckErr(err)
    user.ID =int(id)
    return nil
}

func (dao *meetingDao) FindAll()([]Meeting, error) {
    sqlStmt := `SELECT * FROM meetings`
    rows, err := dao.Query(sqlStmt)
    defer rows.Close()
    CheckErr(err)
    meeting_list := make(Meeting, 0, 0)
    for rows.Next(){
        meetings = Meeting{}
        err := rows.Scan(&meetings.ID, &meeting.Owner, &meetings.Title, &memberList, &meetings.Starttime, &meetings.Endtime)
        CheckErr(err)
        meetings.Members = strings.Split(memberList, "&")
        meeting_list = append(meeting_list, meetings)
    }
    return meeting_list, nil
}

func(dao *meetingDao) FindBy(col string, value) (Meeting, error) {
    sqlStmt := `SELECT * FROM meetings WHERE` + col + ` = '` + value + `';`
    rows, err := dao.Query(sqlStmt)
    defer rows.Close()
    CheckErr(err)
    meeting := Meeting{}
    if rows.Next(){
        err := rows.Scan(&meeting.ID, &meeting.Owner,&meeting.Title,  &memberList, &meeting.Starttime, &meeting.Endtime)
        CheckErr(err)
        meeting.Members = strings.Split(memberList, "&")
    }
    return meeting, nil
}

func(dao *MeetingDao) DeleteMeetingByOwner(user string) error {
    sqlStmt :=sqlStmt :=  `DELETE FROM meetings WHERE owner = '` + key + `';`
    _, err := dao.Exec(sqlStmt)
    CheckErr(err)
    return nil
}
