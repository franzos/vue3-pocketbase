package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	pb "github.com/pocketbase/pocketbase"
	core "github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pb.New()

	// Automatic migration
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
		Dir:         "migrations",
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

        app.OnModelAfterCreate().Add(func(e *core.ModelEvent) error {
            if e.Model.TableName() == "submissions" {
        
                // Get the submission record by ID
                submissionRecord, err := app.Dao().FindRecordById("submissions", e.Model.GetId())
                if err != nil {
                    log.Printf("Error fetching submission record: %v\n", err)
                    return err
                }
        
                // Get the forwarders ids from the submission record
                forwarderIds := submissionRecord.GetStringSlice("forwarders")
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

                submissionRecordData := make(map[string]interface{})
                submissionRecord.UnmarshalJSONField("data", &submissionRecordData)
                normalizedData, err := extractNormalizedData(submissionRecordData)
                if err != nil {
                    return err
                }
        
                for _, forwarderRecord := range forwarderRecords {
                    err := sendEmailIfContactInfoExists(normalizedData, forwarderRecord.GetString("email"))
                    if err != nil {
                        log.Printf("Error sending email to forwarder: %v\n", err)
                        return err
                    }
                }
            }
            return nil
        })

		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/submit/:forwarders",
			Handler: func(c echo.Context) error {
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

                // Calculate the spam score
                spamScore := calculateSpamScore(data)
                // if spamScore >= 50 { // You can adjust the threshold as needed
                //     return c.JSON(http.StatusForbidden, map[string]string{"error": "Submission rejected as spam"})
                // }

                forwarderIDsInterface := make([]interface{}, len(forwarderIDs))
                for i, v := range forwarderIDs {
                    forwarderIDsInterface[i] = v
                }

                form.LoadData(map[string]interface{}{
                    "data":       data,
                    "forwarders": forwarderIDsInterface,
                    "spamScore":  spamScore,
                    "userFlaggedAsSpam": false,
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
    
			
