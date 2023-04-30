package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	pb "github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	core "github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pb.New()

	// Automatic migration
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
		Dir:         "migrations",
	})

	app.OnModelAfterCreate().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == "submissions" {

			// Get the submission record by ID
			record, err := app.Dao().FindRecordById("submissions", e.Model.GetId())
			if err != nil {
				log.Printf("Error fetching submission record: %v\n", err)
				return err
			}

			// Get the forwarders ids from the submission record
			forwarderIds := record.GetStringSlice("forwarders")
			if err != nil {
				log.Printf("Error getting forwarders ids from submission record: %v\n", err)
				return err
			}

			// Get all forwarder records by their ids
			forwarderRecords, err := app.Dao().FindRecordsByIds("forwarders", forwarderIds)
			if err != nil {
				log.Printf("Error fetching forwarder records: %v\n", err)
				return err
			}

			recordData := make(map[string]interface{})
			record.UnmarshalJSONField("data", &recordData)
			normalizedData, err := extractNormalizedData(recordData)
			if err != nil {
				return err
			}

			for _, forwarderRecord := range forwarderRecords {
				res, err := sendEmailIfContactInfoExists(EmailApiData{
					app:            *app,
					data:           normalizedData,
					recipientEmail: forwarderRecord.GetString("recipient"),
				})
				if err != nil {
					createLog(LogsApi{
						dao: app.Dao(),
						data: Log{
							Message:   "Email failed to send",
							Topic:     "email.send.error",
							Level:     2,
							Forwarder: forwarderRecord.GetId(),
						}})
					return err
				} else {
					createLog(LogsApi{
						dao: app.Dao(),
						data: Log{
							Message:   "Email sent",
							Topic:     "email.send.success",
							Level:     1,
							Forwarder: forwarderRecord.GetId(),
							Data:      res.toJSON(),
						}})
				}
			}
		}
		return nil
	})

	app.OnModelAfterUpdate().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == "submissions" {
			record := e.Model.(*models.Record)

			// If submission is flagged as spam, ensure a blocklist entry exists
			userFlaggedAsSpam := record.GetBool("userFlaggedAsSpam")
			if userFlaggedAsSpam {
				sourceIp := record.GetString("sourceIp")
				CreateBlocklistRecord(BlocklistApi{
					dao: app.Dao(),
					data: Blocklist{
						Value: sourceIp,
					},
				})
			}
		}
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// middleware to log every request
		e.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				// Log the request
				log.Printf("%s %s\n", c.Request().Method, c.Request().URL.String())
				// Call the next handler
				return next(c)
			}
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/submit/:forwarders",
			Handler: func(c echo.Context) error {
				sourceIp := c.RealIP()
				forwardersParam := c.PathParam("forwarders")
				forwarderIDs := strings.Split(forwardersParam, ":")

				// Check if the forwarders exist
				records, err := app.Dao().FindRecordsByIds("forwarders", forwarderIDs)
				if err != nil {
					return apis.NewNotFoundError("The forwarder does not exist.", err)
				}

				var forwarders []models.Record
				for _, record := range records {
					forwarder := record.CleanCopy()
					forwarders = append(forwarders, *forwarder)
				}

				// Get the submissions collection
				collection, err := app.Dao().FindCollectionByNameOrId("submissions")
				if err != nil {
					return err
				}

				// Create a new record
				record := models.NewRecord(collection)
				form := forms.NewRecordUpsert(app, record)

				// Load data from the request
				data := make(map[string]interface{})
				for key, values := range c.Request().Form {
					if len(values) > 0 {
						sanitizedValue := sanitizeInput(values[0])
						if !validateInput(sanitizedValue) {
							return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
						}
						data[key] = sanitizedValue
					}
				}

				// Check if the submission is spam
				spamScore := 0
				// (1) Check if we've got a blacklist record for the source IP
				blacklistRecord, err := FindBlocklistRecordByValue(app.Dao(), sourceIp)
				if err != nil {
					log.Println("Error fetching blacklist record: ", err)
				}
				// (2) If we've got a blacklist record, then the submission is spam, otherwise calculate spam score
				if blacklistRecord != nil {
					spamScore += 100
				} else {
					spamScore = calculateSpamScore(data)
				}

				forwarderIDsInterface := make([]interface{}, len(forwarderIDs))
				for i, v := range forwarderIDs {
					forwarderIDsInterface[i] = v
				}

				form.LoadData(map[string]interface{}{
					"data":              data,
					"forwarders":        forwarderIDsInterface,
					"spamScore":         spamScore,
					"userFlaggedAsSpam": false,
					"sourceIp":          sourceIp,
				})

				if err := form.Submit(); err != nil {
					return err
				}

				return c.JSON(http.StatusOK, record)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			}})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
