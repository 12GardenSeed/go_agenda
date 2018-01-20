package entities

// MeetingServiceProvider provide series of op to Entity meetings
type MeetingServiceProvider struct {
}

var UserService = UserServiceProvider{}

func (*MeetingServiceProvider) Insert(meeting *Meeting) error {
	dao := meetingDao{mydb}
	err := dao.Insert(meeting)
	CheckErr(err)
	return nil
}

func (*MeetingServiceProvider) FindAll() ([]Meeting, error) {
	dao := meetingDao{mydb}
	meetings, err := dao.FindAll()
	CheckErr(err)
	return meetings, nil
}

func (*MeetingServiceProvider) DeleteMeetingByOwner(user string) error {
	dao := meetingDao{mydb}
	err := dao, DeleteMeetingByOwner(user)
	CheckErr(err)
	return nil
}
