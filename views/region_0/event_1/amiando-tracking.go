package event_1

import (
	"bytes"
	"errors"
	"github.com/AlexTi/go-amiando"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-mail"
	"github.com/ungerik/go-start/debug"
	gostart_model "github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
	"log"
	"strconv"
)

func init() {
	debug.Nop()
	Region0_Event1_AmiandoTracking = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				region, event, err := RegionAndEvent(ctx.URLArgs)
				if err != nil {
					return err
				}

				args := ctx.Request.Params

				if !Config.IsProductionServer {
					//get main organiser
					var organiser models.Person
					err := event.Organisers[0].Get(&organiser)
					if err != nil {
						return err
					}
					organiserEmail := organiser.User.Email[0].Address.String()

					//generate password
					password, err := gostart_model.GenerateRandomPassword(10)
					if err != nil {
						return err
					}

					userEmail := "alex@startuplive.in"

					from := event.Start.Format("2.")
					until := event.End.Format("2. Jan 2006")
					Ctx := map[string]string{
						"EventName":            event.Name.Get(),
						"Date":                 from + " - " + until,
						"ParticipantFirstName": "Alexander",
						"ParticipantLastName":  "Tacho",
						"ParticipantEmail":     "alex@startuplive.in",
						"ParticipantPassword":  password,
						"OrganiserName":        organiser.Name.String(),
						"OrganiserEmail":       organiserEmail,
						"PitcherForm":          "http://startuplive.in/" + region.Slug.String() + "/" + event.Number.String() + "/pitcher-registration",
					}

					var m bytes.Buffer

					RenderTemplateString(event.PitcherRegistrationWelcomeText.String(),
						"pitcher-registration", &m, Ctx)

					subject := event.Name.Get() + " - Thanks"
					message := m.String()

					go func() {
						err = email.NewBriefMessageFrom(subject, message, organiserEmail, userEmail).Send()

					}()

					return nil
				}

				// debug.Print("identifier: ", args["eventIdentifier"])

				if args["eventIdentifier"] != event.AmiandoEventIdentifier.String() {
					return errors.New("Error: Trying to sync participant of another event")
				}

				nroftickets, err := strconv.Atoi(args["numberOfTickets"])
				if err != nil {
					return err

				}

				if nroftickets > 0 {
					identifier := event.AmiandoEventIdentifier.Get()

					api := amiando.NewApi(event.AmiandoEventApiKey.Get())
					amiandoEvent, err := amiando.NewEvent(api, identifier)
					if err != nil {
						return err
					}

					paymentId := args["paymentId"]
					var buf bytes.Buffer
					// numParticipants := 0
					p, e := amiandoEvent.EnumParticipantsByPayment(paymentId)
					for participant, ok := <-p; ok; participant, ok = <-p {
						eventparticipant, err := event.UpdateAmiandoParticipant(log.New(&buf, "", 0), participant)
						// _, err := event.UpdateAmiandoParticipant(log.New(&buf, "", 0), participant)
						if err != nil {
							return err
						}

						if eventparticipant.PresentsIdea {
							errchan := sendPitcherConfirmationEmail(ctx, eventparticipant)
							if err, ok := <-errchan; ok {
								return err
							}
						}

						// numParticipants++
					}
					if err, ok := <-e; ok {
						return err
					}

				}

				return nil
			},
		),
	)
}

func sendPitcherConfirmationEmail(ctx *Context, p *models.EventParticipant) <-chan error {
	errChan := make(chan error, 1)

	region, event, err := RegionAndEvent(ctx.URLArgs)
	if err != nil {
		errChan <- err
		return errChan
	}

	//get person 	
	var person models.Person
	err = p.Person.Get(&person)
	if err != nil {
		errChan <- err
		return errChan
	}
	personEmail := person.User.Email[0].Address.String()

	//get main organiser 	
	var organiser models.Person
	err = event.Organisers[0].Get(&organiser)
	if err != nil {
		errChan <- err
		return errChan
	}
	// debug.Print("organiser email: ", organiser.User.Email[0].Address.String())
	organiserEmail := organiser.User.Email[0].Address.String()

	//generate password
	password, err := gostart_model.GenerateRandomPassword(10)
	if err != nil {
		errChan <- err
		return errChan
	}
	// debug.Print(password)
	person.Password.SetHashed(password)
	person.User.ConfirmEmailPassword()
	err = person.Save()
	if err != nil {
		errChan <- err
		return errChan
	}

	from := event.Start.Format("2.")
	until := event.End.Format("2. Jan 2006")
	TemplCtx := map[string]string{
		"EventName":            event.Name.Get(),
		"Date":                 from + " - " + until,
		"ParticipantFirstName": person.Name.First.String(),
		"ParticipantLastName":  person.Name.Last.String(),
		"ParticipantEmail":     personEmail,
		"ParticipantPassword":  password,
		"OrganiserName":        organiser.Name.String(),
		"OrganiserEmail":       organiserEmail,
		"PitcherForm":          "http://startuplive.in/" + region.Slug.String() + "/" + event.Number.String() + "/pitcher-registration",
	}

	pitcherwelcometext := event.PitcherRegistrationWelcomeText.String()
	if event.PitcherRegistrationWelcomeText.String() == "" {
		pitcherwelcometext = event.GetDefaultPitcherRegistrationText()
	}

	var m bytes.Buffer
	RenderTemplateString(pitcherwelcometext,
		"pitcher-registration", &m, TemplCtx)

	subject := event.Name.Get() + " - Thanks for registration"
	message := m.String()

	go func() {
		errChan <- email.NewBriefMessageFrom(subject, message, organiserEmail, personEmail).Send()
		close(errChan)
	}()

	return errChan
}
