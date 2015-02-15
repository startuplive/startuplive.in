package models

import (
	"strconv"

	"github.com/ungerik/go-start/config"
	// "github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/view"
)

var EventTeams = mongo.NewCollection("eventteams", "Name")

func NewEventTeam() *EventTeam {
	var doc EventTeam
	EventTeams.InitDocument(&doc)
	return &doc
}

type EventTeam struct {
	mongo.DocumentBase `bson:",inline"`
	Event              mongo.Ref    `model:"to=events|required"`
	LeaderNameInput    model.String `view:"label=Your Name|placeholder=firstname lastname"`
	EmailInput         model.Email  `view:"label=Your Email|placeholder=your email"`
	Name               model.String `model:"required" view:"label=Team Name"`
	Tagline            model.String `model:"required" view:"label=Tagline"`
	Leader             mongo.Ref    `model:"to=people"`
	BusinessFields     model.Choice `view:"label=Business Field" model:"options=Arts and Entertainment,Business & Professional Services,Clothing & Accessoires,Community & Government,Construction & Contractors,Education,Food & Dining,Gaming,Health & Medicine,Home & Garden,Industry & Agriculture,Legal & Financial,Media & Communications,Personal Care & Services,Real Estate,Shopping,Sports & Recreation,Travel & Transportation"`
	CurrentStage       model.Choice `view:"label=Current Stage" model:"options=Concept Stage(got an idea),Seed Stage (working on the product),Early Stage (close to market),Growth Stage (we are out there and making some cash),Break Even (we already hit the magic border)"`
	FoundingDate       model.String `model:"required" view:"size=8|label=Since when are you working on your idea? (i.e.: 04/2012)"`

	TeamMembers           []mongo.Ref `model:"to=people" view:"label=Team Members`
	Logo                  media.ImageRef
	LogoURL               model.Url `view:"label=Old Logo URL"`
	Homepage              model.Url
	CancelledDate         model.DateTime
	CancelledBy           mongo.Ref  `model:"to=people"`
	Abstract              model.Text `model:"required"`
	PitchtrainingBooked   model.Bool
	PitchtrainingAttended model.DateTime
	ProblemOpportunity    model.Text   `model:"required" view:"label=Problem / Opportunity"`
	Solution              model.Text   `model:"required"`
	Haves                 model.Text   `model:"required" view:"label=Haves? (Team, Product / Service, Funding, ...) "`
	NeedTechies           model.Bool   `view:"label=I need techies."`
	NeedBizPeople         model.Bool   `view:"label=I need business people."`
	NeedDesigners         model.Bool   `view:"label=I need designers."`
	NeedOther             model.String `view:"label=I need anybody else."`
	Pitching              model.Bool
	BookedMentors         []mongo.Ref `model:"to=people"`
	CrowdVotes            model.Int
	FacebookURL           model.Url `view:"label=Facebook Url"`
	TwitterURL            model.Url `view:"label=Twitter Url"`
	LinkedInURL           model.Url `view:"label=LinkedIn Url"`
	YoutubeURL            model.Url `view:"label=Youtube Url"`
	Judgements            []Judgement

	/*
		RequestedMentors model.Array{Of: &Ref{To: &Person}}},  // own mentorship struct?
		Goals model.Array{
			Of: &Document{
				Nodes: Nodes{
					Description model.String
					Achieved model.DateTime
				},
			},
		}},
		Judgements model.Array{
			Of: &Document{
				Nodes: Nodes{
					Judge model.Ref{To: &Person}},
					TODO categories model.String // todo -> surveys
				},
			},
		}},
	*/
}

type Judgement struct {
	Judge          mongo.Ref    `model:"to=people|required"`
	Presentation   model.Choice `model:"options=k.A,1,2,3,4,5"`
	Innovation     model.Choice `model:"options=k.A,1,2,3,4,5"`
	Traction       model.Choice `model:"options=k.A,1,2,3,4,5"`
	TeamImpression model.Choice `model:"options=k.A,1,2,3,4,5"`
	Market         model.Choice `model:"options=k.A,1,2,3,4,5"`
	Scalability    model.Choice `model:"options=k.A,1,2,3,4,5"`
	Feasability    model.Choice `model:"options=k.A,1,2,3,4,5"`
	Score          model.Float  `view:"hidden"`
	ImageVoting    model.Choice `model:"options=k.A,1,2,3,4,5"`
}

const (
	JudgementPresentation   = "Presentation"
	JudgementInnovation     = "Innovation"
	JudgementTraction       = "Traction"
	JudgementTeamImpression = "TeamImpression"
	JudgementMarket         = "Market"
	JudgementScalability    = "Scalability"
	JudgementFeasability    = "Feasability"
	JudgementImageVoting    = "Image Voting"

	judgementPresentationWeight = 1.1
	judgementInnovationWeight   = 1.2
	judgementTractionWeight     = 1.1
	judgementImpressionWeight   = 1.2
	judgementMarketWeight       = 1.15
	judgementScalabilityWeight  = 1.1
	judgementFeasabilityWeight  = 1.15
	judgementImageVotingWeight  = 0.5

	TeamNotJudgedByJudge int = iota
	TeamJudgingIncompleteByJudge
	TeamJudgedByJudge
)

func (self *EventTeam) String() string {
	return self.Name.Get()
}

func (self *EventTeam) LogoImage(class string, width int) (*view.Image, error) {
	if self.Logo.IsEmpty() {
		return &view.Image{
			Class: class,
			Src:   self.LogoURL.GetOrDefault("/images/dashb/unknown160x160.png"),
			Width: width,
		}, nil
	}
	version, err := self.Logo.VersionWidth(width, false)
	if err != nil {
		config.Logger.Println("Probably an invalid image ref: ", err)
		self.Logo.Set(nil)
		return self.LogoImage(class, width)
	}
	return version.View(class), nil
}

func (self *EventTeam) ParticipantIterator() model.Iterator {
	return EventParticipants.Filter("Team", self.ID).Iterator()
}

func (self *EventTeam) PersonIterator(includeCancelled bool) model.Iterator {
	peopleIDs := make([]interface{}, 0, 128)
	i := EventParticipants.Filter("Team", self.ID).Iterator()
	var participant EventParticipant
	for i.Next(&participant) {
		if includeCancelled || !participant.Cancelled() {
			peopleIDs = append(peopleIDs, participant.Person.ID)
		}
	}
	if i.Err() != nil {
		return model.NewErrorOnlyIterator(i.Err())
	}
	return People.FilterIn("_id", peopleIDs...).Sort("Name.First").Sort("Name.Last").Iterator()
}

func (self *EventTeam) LeaderName() string {
	var leader Person
	found, err := self.Leader.TryGet(&leader)
	if err != nil {
		return err.Error()
	}
	if !found {
		return " warning - name not set "
	}
	return leader.Name.String()
}

func (self *EventTeam) Cancelled() bool {
	return self.CancelledDate.Get() != ""
}

func (self *EventTeam) CancelAndSave(by *Person) error {
	self.CancelledDate.SetNowUTC()
	if by != nil {
		self.CancelledBy.Set(by)
	}
	err := self.Save()
	if err != nil {
		return err
	}

	// Remove reference to team at team members
	i := self.ParticipantIterator()
	var participant EventParticipant
	for i.Next(&participant) {
		participant.Team.Set(nil)
		err = participant.Save()
		if err != nil {
			return err
		}
	}
	return i.Err()
}

func (self *EventTeam) UncancelAndSave() error {
	self.CancelledDate.SetEmpty()
	self.CancelledBy.Set(nil)
	return self.Save()
}

func (self *EventTeam) HasJudged(judge *Person) (key int, status int) {
	incomplete := false

	for i := 0; i < len(self.Judgements); i++ {

		if self.Judgements[i].Judge.ID == judge.ID {

			model.Visit(&self.Judgements[i], model.VisitorFunc(
				func(data *model.MetaData) error {
					if choice, ok := data.Value.Addr().Interface().(*model.Choice); ok {
						//debug.Print("***** walk structure : ", choice.Get())
						if choice.Get() == "k.A" {
							incomplete = true
						}
					}
					return nil
				},
			))
			if incomplete {
				return i, TeamJudgingIncompleteByJudge
			}
			return i, TeamJudgedByJudge
		}
	}

	return len(self.Judgements), TeamNotJudgedByJudge
}

// func (self *EventTeam) ComputeTeamScoreByEvent() float64 {
// 	var overallscore float64
// 	for i := 0; i < len(self.Judgements); i++ {
// 		judgementscore := self.Judgements[i].Score.Get()
// 		overallscore = overallscore + judgementscore
// 		//debug.Print("********** Score: ", overallscore)
// 	}

// 	return overallscore
// }

func (self *EventTeam) ComputeAverageScoreByEvent() float64 {

	var averagescore float64

	averagescore += self.ComputeAverageJudgementCategory(JudgementPresentation)
	averagescore += self.ComputeAverageJudgementCategory(JudgementInnovation)
	averagescore += self.ComputeAverageJudgementCategory(JudgementTraction)
	averagescore += self.ComputeAverageJudgementCategory(JudgementTeamImpression)
	averagescore += self.ComputeAverageJudgementCategory(JudgementMarket)
	averagescore += self.ComputeAverageJudgementCategory(JudgementScalability)
	averagescore += self.ComputeAverageJudgementCategory(JudgementFeasability)
	averagescore += self.ComputeAverageJudgementCategory(JudgementImageVoting)

	return averagescore
}

func (self *EventTeam) ComputeAverageCatScoreByEvent() float64 {

	var averagescore float64
	var filledCats float64

	if self.ComputeAverageJudgementCategory(JudgementPresentation) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementPresentation)
		filledCats++
	}
	if self.ComputeAverageJudgementCategory(JudgementInnovation) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementInnovation)
		filledCats++
	}
	if self.ComputeAverageJudgementCategory(JudgementTraction) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementTraction)
		filledCats++
	}
	if self.ComputeAverageJudgementCategory(JudgementTeamImpression) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementTeamImpression)
		filledCats++
	}
	if self.ComputeAverageJudgementCategory(JudgementMarket) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementMarket)
		filledCats++
	}
	if self.ComputeAverageJudgementCategory(JudgementScalability) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementScalability)
		filledCats++
	}
	if self.ComputeAverageJudgementCategory(JudgementFeasability) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementFeasability)
		filledCats++
	}
	if self.ComputeAverageJudgementCategory(JudgementImageVoting) != 0 {
		averagescore += self.ComputeAverageJudgementCategory(JudgementImageVoting)
		filledCats++
	}

	return averagescore / filledCats
}

func (self *EventTeam) ComputeTeamScoreByJudge(judgement *Judgement) float64 {
	presentation, _ := strconv.ParseFloat(judgement.Presentation.Get(), 64)
	presentation = presentation * judgementPresentationWeight

	innovation, _ := strconv.ParseFloat(judgement.Innovation.Get(), 64)
	innovation = innovation * judgementInnovationWeight

	traction, _ := strconv.ParseFloat(judgement.Traction.Get(), 64)
	traction = traction * judgementTractionWeight

	teamImpression, _ := strconv.ParseFloat(judgement.TeamImpression.Get(), 64)
	teamImpression = teamImpression * judgementImpressionWeight

	market, _ := strconv.ParseFloat(judgement.Market.Get(), 64)
	market = market * judgementMarketWeight

	scalability, _ := strconv.ParseFloat(judgement.Scalability.Get(), 64)
	scalability = scalability * judgementScalabilityWeight

	feasability, _ := strconv.ParseFloat(judgement.Feasability.Get(), 64)
	feasability = feasability * judgementFeasabilityWeight

	imagevoting, _ := strconv.ParseFloat(judgement.ImageVoting.Get(), 64)
	imagevoting = imagevoting * judgementImageVotingWeight

	computedscore := presentation + innovation + traction + teamImpression + market + scalability + feasability + imagevoting

	return computedscore

}

func (self *EventTeam) ComputeAverageTeamScoreByJudge(judgement *Judgement) (score float64) {

	presentation, _ := strconv.ParseFloat(judgement.Presentation.Get(), 64)
	presentation = presentation * judgementPresentationWeight

	innovation, _ := strconv.ParseFloat(judgement.Innovation.Get(), 64)
	innovation = innovation * judgementInnovationWeight

	traction, _ := strconv.ParseFloat(judgement.Traction.Get(), 64)
	traction = traction * judgementTractionWeight

	teamImpression, _ := strconv.ParseFloat(judgement.TeamImpression.Get(), 64)
	teamImpression = teamImpression * judgementImpressionWeight

	market, _ := strconv.ParseFloat(judgement.Market.Get(), 64)
	market = market * judgementMarketWeight

	scalability, _ := strconv.ParseFloat(judgement.Scalability.Get(), 64)
	scalability = scalability * judgementScalabilityWeight

	feasability, _ := strconv.ParseFloat(judgement.Feasability.Get(), 64)
	feasability = feasability * judgementFeasabilityWeight

	imagevoting, _ := strconv.ParseFloat(judgement.ImageVoting.Get(), 64)
	imagevoting = imagevoting * judgementImageVotingWeight

	computedscore := presentation + innovation + traction + teamImpression + market + scalability + feasability + imagevoting

	var valfields float64
	model.Visit(&judgement, model.VisitorFunc(
		func(data *model.MetaData) error {
			if choice, ok := data.Value.Addr().Interface().(*model.Choice); ok {
				//debug.Print("***** walk structure : ", choice.Get())
				if choice.Get() != "k.A" {
					valfields++
				}
			}
			return nil
		},
	))

	score = computedscore / valfields

	return score

}

func (self *EventTeam) ComputeAverageJudgementCategory(category string) (score float64) {
	var overallscore float64
	var judgements float64

	//debug.Print("----> computing average judgement")

	for i := 0; i < len(self.Judgements); i++ {
		switch category {
		case JudgementPresentation:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].Presentation.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementPresentationWeight)
			if err == nil {
				judgements++
			}
			break
		case JudgementInnovation:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].Innovation.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementInnovationWeight)
			if err == nil {
				judgements++
			}
			break
		case JudgementTraction:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].Traction.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementTractionWeight)
			if err == nil {
				judgements++
			}
			break
		case JudgementTeamImpression:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].TeamImpression.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementImpressionWeight)
			if err == nil {
				judgements++
			}
			break
		case JudgementMarket:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].Market.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementMarketWeight)
			if err == nil {
				judgements++
			}
			break
		case JudgementScalability:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].Scalability.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementScalabilityWeight)
			if err == nil {
				judgements++
			}
			break
		case JudgementFeasability:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].Feasability.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementFeasabilityWeight)
			if err == nil {
				judgements++
			}
			break
		case JudgementImageVoting:
			judgementscore, err := strconv.ParseFloat(self.Judgements[i].ImageVoting.Get(), 64)
			overallscore = overallscore + (judgementscore * judgementImageVotingWeight)
			if err == nil {
				judgements++
			}
			break
		}
	}

	if judgements > 0 {
		score = overallscore / judgements
	} else {
		score = 0
	}
	//debug.Print("----> computed score: ", score)

	return score
}
