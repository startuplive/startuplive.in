package views

import (
	"github.com/ungerik/go-start/media"
	. "github.com/ungerik/go-start/view"
)

var Admin_Auth Authenticator
var SuperAdmin_Auth Authenticator
var Admin_Region0_Auth Authenticator
var Region0_Event1_Admin_Auth Authenticator
var Region0_Event1_Dashboard_Auth Authenticator

var Admin *Page
var Admin_CSS ViewWithURL
var Admin_People *Page
var Admin_ExportEmails ViewWithURL
var Admin_ExportPeople ViewWithURL
var Admin_ExportPitcherEmails ViewWithURL
var Admin_ExportMentorJudgeEmails ViewWithURL
var Admin_ExportContacts ViewWithURL
var Admin_Images *Page
var Admin_Files *Page
var Admin_Person0 *Page
var Admin_Startups *Page
var Admin_Startup0 *Page
var Admin_Teams *Page
var Admin_Events *Page
var Admin_CitySuggestions *Page
var Admin_Wiki *Page
var Admin_WikiEntry0 *Page
var Admin_Regions *Page
var Admin_Region0 *Page
var Admin_Region0_Logo *Page
var Admin_Region0_Logo_Gen *Page
var Admin_Region0_Logo_Gen_SVG *Page
var Admin_Region0_SyncEvent1 *Page
var Region0_Event1_Admin_Mentor2 *Page

var Blog ViewWithURL
var Blog_0 ViewWithURL
var Blog_0_1 ViewWithURL
var Blog_0_1_2 ViewWithURL

var CSS View
var Contact ViewWithURL

var Profile *Page
var MyStartup *Page

var LoginSignup *Page
var ConfirmEmail *Page
var Logout ViewWithURL

var Homepage ViewWithURL
var Imprint ViewWithURL
var Organisers ViewWithURL
var StartupForm ViewWithURL
var StartupFormSuccess *Page

var Region0 ViewWithURL
var Region0_CSS ViewWithURL
var Region0_Event1_FAQ *Page
var Region0_Event1_Extra *Page
var Region0_Event1_Location *Page
var Region0_Event1_Judges *Page
var Region0_Event1_Organisers *Page
var Region0_Event1_Registration *Page
var Region0_Event1_Schedule *Page
var Region0_Event1_Voting *Page
var Region0_Event1_Voting_Dialog *Page
var Region0_Event1_FeedbackParticipants *Page
var Region0_Event1_MentorJudgeFeedback *Page
var Region0_Event1_OrganiserFeedback *Page
var Region0_Event1_Pitcher_Registration *Page
var Region0_Event1_Mentor_Registration *Page
var Region0_Event1_Judge_Registration *Page
var Region0_Event1_Registration_Success *Page
var Region0_Event1_AmiandoTest *Page
var Region0_Event1_AmiandoTracking ViewWithURL
var Region0_Event1 *Page
var Region0_Event1_Admin *Page
var Region0_Event1_Admin_ExportEmails ViewWithURL
var Region0_Event1_Admin_ExportPresentEmails ViewWithURL
var Region0_Event1_Admin_ChooseLogo *Page
var Region0_Event1_Admin_ChooseLogo_SVG *Page
var Region0_Event1_Admin_About *Page
var Region0_Event1_Admin_Judges *Page
var Region0_Event1_Admin_Judges_Export ViewWithURL
var Region0_Event1_Admin_Judge2 *Page
var Region0_Event1_Admin_Location *Page
var Region0_Event1_Admin_Mentors *Page
var Region0_Event1_Admin_Mentors_Export ViewWithURL
var Region0_Event1_Admin_Organisers *Page
var Region0_Event1_Admin_Organiser2 *Page
var Region0_Event1_Admin_Participants *Page
var Region0_Event1_Admin_Participant2 *Page
var Region0_Event1_Admin_Schedule *Page
var Region0_Event1_Admin_Schedule2 *Page
var Region0_Event1_Admin_Teams *Page
var Region0_Event1_Admin_Team2 *Page
var Region0_Event1_Admin_Partners *Page
var Region0_Event1_Admin_Partner2 *Page
var Region0_Event1_Admin_Judgements *Page
var Region0_Event1_Admin_Judgements_Team2 *Page
var Region0_Event1_Admin_Voting *Page
var Region0_Event1_Admin_Amiando *Page
var Region0_Event1_Admin_AmiandoData *Page
var Region0_Event1_Admin_Settings *Page
var Region0_Event1_Admin_FAQ *Page
var Region0_Event1_Admin_DashboardInfo *Page
var Region0_Event1_Admin_Feedback *Page
var Region0_Event1_Admin_Wiki *Page
var Region0_Event1_Admin_WikiEntry0 *Page

var Region0_DashboardCSS ViewWithURL
var Region0_DashboardSubmodalCSS ViewWithURL

var Region0_Event1_Dashboard *Page
var Region0_Event1_Dashboard_Info *Page
var Region0_Event1_Dashboard_Judges *Page
var Region0_Event1_Dashboard_Judge2 *Page
var Region0_Event1_Dashboard_Mentors *Page
var Region0_Event1_Dashboard_Mentor2 *Page
var Region0_Event1_Dashboard_Organisers *Page
var Region0_Event1_Dashboard_Participants *Page
var Region0_Event1_Dashboard_Participant2 *Page
var Region0_Event1_Dashboard_Teams *Page
var Region0_Event1_Dashboard_Team2 *Page
var Region0_Event1_Dashboard_VotingResult *Page

var Events_Application *Page
var Events *Page
var Events_FAQ *Page
var Events_Where *Page
var Events_YourCity *Page

var API_Judges ViewWithURL
var API_Mentors ViewWithURL
var API_Organisers ViewWithURL
var API_People ViewWithURL
var API_PartnerOrder ViewWithURL

func Paths() *ViewPath {
	//basicAuth := NewBasicAuth("statuplive.in", "gostart", "gostart")
	return &ViewPath{View: Homepage, Sub: []ViewPath{
		media.ViewPath("media"),
		{Name: "api", Sub: []ViewPath{
			{Name: "people.json", View: API_People},
			{Name: "mentors.json", View: API_Mentors},
			{Name: "judges.json", View: API_Judges},
			{Name: "event-organisers.json", View: API_Organisers},
		}},
		{Name: "style.css", View: CSS},
		{Name: "admin", View: Admin, Auth: Admin_Auth, Sub: []ViewPath{
			{Name: "style.css", View: Admin_CSS},
			{Name: "events", View: Admin_Events, Auth: Admin_Auth},
			{Name: "people", View: Admin_People, Auth: Admin_Auth, Sub: []ViewPath{
				{Name: "export-emails", View: Admin_ExportEmails, Auth: Admin_Auth},
				{Name: "export-people", View: Admin_ExportPeople, Auth: Admin_Auth},
				{Name: "export-pitcher-emails", View: Admin_ExportPitcherEmails, Auth: Admin_Auth},
				{Name: "export-mentor-judge-emails", View: Admin_ExportMentorJudgeEmails, Auth: Admin_Auth},
				{Name: "export-contacts", View: Admin_ExportContacts, Auth: Admin_Auth},
			}},
			{Name: "person", Args: 1, View: Admin_Person0, Auth: Admin_Auth},
			{Name: "startup", Args: 1, View: Admin_Startup0, Auth: Admin_Auth},
			{Name: "teams", View: Admin_Teams, Auth: Admin_Auth},
			{Name: "startups", View: Admin_Startups, Auth: Admin_Auth},
			{Name: "regions", View: Admin_Regions, Auth: Admin_Auth},
			{Name: "citysuggestions", View: Admin_CitySuggestions, Auth: Admin_Auth},
			{Name: "wiki", View: Admin_Wiki, Auth: Admin_Auth},
			{Name: "wikientry", Args: 1, View: Admin_WikiEntry0, Auth: Admin_Auth},
			{Name: "images", View: Admin_Images, Auth: Admin_Auth},
			{Name: "files", View: Admin_Files, Auth: Admin_Auth},
			{Args: 1, View: Admin_Region0, Auth: Admin_Region0_Auth, Sub: []ViewPath{
				{Name: "logo", View: Admin_Region0_Logo, Auth: Admin_Region0_Auth, Sub: []ViewPath{
					{Args: 3, View: Admin_Region0_Logo_Gen, Auth: Admin_Region0_Auth, Sub: []ViewPath{
						{Name: "svg", View: Admin_Region0_Logo_Gen_SVG},
					}},
				}},
				{Name: "sync-event", Args: 1, View: Admin_Region0_SyncEvent1, Auth: Admin_Region0_Auth},
			}},
		}},
		{Name: "login", View: LoginSignup, Sub: []ViewPath{
			{Name: "confirm", View: ConfirmEmail},
		}},
		{Name: "logout", View: Logout},
		{Name: "profile", View: Profile},
		{Name: "my-startup", Args: 1, View: MyStartup},
		{Name: "events", View: Events, Sub: []ViewPath{
			{Name: "where", View: Events_Where},
			{Name: "faq", View: Events_FAQ},
			{Name: "your-city", View: Events_YourCity},
			{Name: "application", View: Events_Application},
		}},
		{Name: "blog", View: Blog, Sub: []ViewPath{
			{Args: 1, View: Blog_0, Sub: []ViewPath{
				{Args: 1, View: Blog_0_1, Sub: []ViewPath{
					{Args: 1, View: Blog_0_1_2},
				}},
			}},
		}},
		{Name: "contact", View: Contact},
		{Name: "imprint", View: Imprint},
		{Name: "organisers", View: Organisers},
		{Name: "startup-form", View: StartupForm},
		{Name: "startup-form-success", View: StartupFormSuccess},
		{Args: 1, View: Region0, Sub: []ViewPath{
			{Name: "region.css", View: Region0_CSS},
			{Name: "dashboard.css", View: Region0_DashboardCSS},
			{Name: "submodal.css", View: Region0_DashboardSubmodalCSS},
			{Args: 1, View: Region0_Event1, Sub: []ViewPath{
				{Name: "location", View: Region0_Event1_Location},
				{Name: "schedule", View: Region0_Event1_Schedule},
				{Name: "judges", View: Region0_Event1_Judges},
				{Name: "organisers", View: Region0_Event1_Organisers},
				{Name: "faq", View: Region0_Event1_FAQ},
				{Name: "extra", View: Region0_Event1_Extra},
				{Name: "participant-feedback", View: Region0_Event1_FeedbackParticipants},
				{Name: "mentorjudge-feedback", View: Region0_Event1_MentorJudgeFeedback},
				{Name: "organiser-feedback", View: Region0_Event1_OrganiserFeedback},
				{Name: "pitcher-registration", View: Region0_Event1_Pitcher_Registration, Auth: Region0_Event1_Dashboard_Auth},
				{Name: "mentor-registration", View: Region0_Event1_Mentor_Registration},
				{Name: "judge-registration", View: Region0_Event1_Judge_Registration},
				{Name: "registration-success", View: Region0_Event1_Registration_Success},
				{Name: "amiando-tracking", View: Region0_Event1_AmiandoTracking},
				{Name: "amiando-test", View: Region0_Event1_AmiandoTest},
				{Name: "voting", View: Region0_Event1_Voting, Sub: []ViewPath{
					{Name: "voted", Args: 1, View: Region0_Event1_Voting_Dialog},
				}},
				{Name: "registration", View: Region0_Event1_Registration},
				{Name: "dashboard", View: Region0_Event1_Dashboard, Auth: Region0_Event1_Dashboard_Auth, Sub: []ViewPath{
					{Name: "info", View: Region0_Event1_Dashboard_Info, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "mentors", View: Region0_Event1_Dashboard_Mentors, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "mentor", Args: 1, View: Region0_Event1_Dashboard_Mentor2, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "judges", View: Region0_Event1_Dashboard_Judges, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "judge", Args: 1, View: Region0_Event1_Dashboard_Judge2, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "participants", View: Region0_Event1_Dashboard_Participants, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "participant", Args: 1, View: Region0_Event1_Dashboard_Participant2, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "teams", View: Region0_Event1_Dashboard_Teams, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "team", Args: 1, View: Region0_Event1_Dashboard_Team2, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "organisers", View: Region0_Event1_Dashboard_Organisers, Auth: Region0_Event1_Dashboard_Auth},
					{Name: "voting-result", View: Region0_Event1_Dashboard_VotingResult, Auth: Region0_Event1_Dashboard_Auth},
				}},
				{Name: "admin", View: Region0_Event1_Admin, Auth: Region0_Event1_Admin_Auth, Sub: []ViewPath{
					{Name: "logo", Args: 3, View: Region0_Event1_Admin_ChooseLogo, Auth: Region0_Event1_Admin_Auth, Sub: []ViewPath{
						{Name: "svg", View: Region0_Event1_Admin_ChooseLogo_SVG},
					}},
					{Name: "about", View: Region0_Event1_Admin_About, Auth: Region0_Event1_Admin_Auth},
					{Name: "location", View: Region0_Event1_Admin_Location, Auth: Region0_Event1_Admin_Auth},
					{Name: "schedule", View: Region0_Event1_Admin_Schedule, Auth: Region0_Event1_Admin_Auth},
					{Name: "schedule", Args: 1, View: Region0_Event1_Admin_Schedule2, Auth: Region0_Event1_Admin_Auth},
					{Name: "organisers", View: Region0_Event1_Admin_Organisers, Auth: Region0_Event1_Admin_Auth},
					{Name: "organiser", Args: 1, View: Region0_Event1_Admin_Organiser2, Auth: Region0_Event1_Admin_Auth},
					{Name: "participants", View: Region0_Event1_Admin_Participants, Auth: Region0_Event1_Admin_Auth},
					{Name: "participant", Args: 1, View: Region0_Event1_Admin_Participant2, Auth: Region0_Event1_Admin_Auth},
					{Name: "teams", View: Region0_Event1_Admin_Teams, Auth: Region0_Event1_Admin_Auth},
					{Name: "team", Args: 1, View: Region0_Event1_Admin_Team2, Auth: Region0_Event1_Admin_Auth},
					{Name: "mentors", View: Region0_Event1_Admin_Mentors, Auth: Region0_Event1_Admin_Auth, Sub: []ViewPath{
						{Name: "export-emails", View: Region0_Event1_Admin_Mentors_Export, Auth: Admin_Auth},
					}},
					{Name: "mentor", Args: 1, View: Region0_Event1_Admin_Mentor2, Auth: Region0_Event1_Admin_Auth},
					{Name: "judges", View: Region0_Event1_Admin_Judges, Auth: Region0_Event1_Admin_Auth, Sub: []ViewPath{
						{Name: "export-emails", View: Region0_Event1_Admin_Judges_Export, Auth: Admin_Auth},
					}},
					{Name: "judge", Args: 1, View: Region0_Event1_Admin_Judge2, Auth: Region0_Event1_Admin_Auth},
					{Name: "partners", View: Region0_Event1_Admin_Partners, Auth: Region0_Event1_Admin_Auth, Sub: []ViewPath{
						{Name: "order", View: API_PartnerOrder, Auth: Region0_Event1_Admin_Auth},
					}},
					{Name: "partner", Args: 1, View: Region0_Event1_Admin_Partner2, Auth: Region0_Event1_Admin_Auth},
					{Name: "judgements", View: Region0_Event1_Admin_Judgements, Auth: Region0_Event1_Admin_Auth, Sub: []ViewPath{
						{Name: "team", Args: 1, View: Region0_Event1_Admin_Judgements_Team2, Auth: Region0_Event1_Admin_Auth},
					}},
					{Name: "voting", View: Region0_Event1_Admin_Voting, Auth: Region0_Event1_Admin_Auth},
					{Name: "export-emails", View: Region0_Event1_Admin_ExportEmails, Auth: Region0_Event1_Admin_Auth},
					{Name: "export-present-emails", View: Region0_Event1_Admin_ExportPresentEmails, Auth: Region0_Event1_Admin_Auth},
					{Name: "amiando-data", View: Region0_Event1_Admin_AmiandoData, Auth: Region0_Event1_Admin_Auth},
					{Name: "amiando-setup", View: Region0_Event1_Admin_Amiando, Auth: Admin_Region0_Auth},
					{Name: "settings", View: Region0_Event1_Admin_Settings, Auth: Region0_Event1_Admin_Auth},
					{Name: "faq", View: Region0_Event1_Admin_FAQ, Auth: Region0_Event1_Admin_Auth},
					{Name: "dashboard-info", View: Region0_Event1_Admin_DashboardInfo, Auth: Region0_Event1_Admin_Auth},
					{Name: "feedback", View: Region0_Event1_Admin_Feedback, Auth: Region0_Event1_Admin_Auth},
					{Name: "compendium", View: Region0_Event1_Admin_Wiki, Auth: Region0_Event1_Admin_Auth, Sub: []ViewPath{
						{Name: "compendium-article", Args: 1, View: Region0_Event1_Admin_WikiEntry0, Auth: Region0_Event1_Admin_Auth},
					}},
				}},
			}},
		}},
	}}
}
