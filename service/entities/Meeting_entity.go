package entities


//meeting entity
//a meeting map to one meeting entity

//ID    -- the unique identification of meeting
//Owner -- the owner of the meeting
//title -- the theme of the meeting
//members -- the member take part in the meeting
//start_time  -- the time that meeting start
//end_time -- the time that meeting end

type Meeting struct {
    ID        int      `json:id`
    Owner     string   `json:owner`
    Title     string   `json:title`
    Members   []string `json:members`
    Starttime string   `json:starttime`
    Endtime   string   `json:endtime`
}

//NeMeeting returns a new instance of meeting type
func NewMeeting(owner string, title string, members []string, starttime string, endtime string) *Meeting {
    return &Meeting{-1, owner, title, members, starttime, endtime}
}
