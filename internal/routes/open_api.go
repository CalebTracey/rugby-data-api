package routes

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"net/http"
)

//go:generate statik -src=/Users/calebtracey/Desktop/Code/rugby-data-api/swagger-ui
//go:generate go run ../../cmd/openapi-gen/main.go -path ../../swagger-ui
//go:generate oapi-codegen -package openapi3 -generate types  -o ../../pkg/openapi3/types.gen.go ../../swagger-ui/openapi3.yaml
//go:generate oapi-codegen -package openapi3 -generate client -o ../../pkg/openapi3/client.gen.go ../../swagger-ui/openapi3.yaml

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Rugby Data REST API",
			Description: "REST API used for fetching and accessing rugby stats",
			Version:     "0.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/CalebTracey/rugby-data-api",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://0.0.0.0:6080",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"PSQLCompetitionData": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("comp_id", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("compName", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("team_id", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("team_name", openapi3.NewStringSchema().
					WithNullable())),
		"TeamData": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("name", openapi3.NewStringSchema().
					WithNullable())),
		"TeamDataList": openapi3.NewSchemaRef("",
			openapi3.NewArraySchema().
				WithPropertyRef("teamData", &openapi3.SchemaRef{
					Ref: "#/components/schemas/TeamData",
				})),
		"PSQLTeamData": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("team_id", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("name", openapi3.NewStringSchema().
					WithNullable()).
				WithPropertyRef("teams", &openapi3.SchemaRef{
					Ref: "#/components/schemas/TeamDataList",
				}).
				WithPropertyRef("message", &openapi3.SchemaRef{
					Ref: "#/components/schemas/Message",
				})),
		"ErrorLog": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("scope", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("status", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("trace", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("rootCause", openapi3.NewStringSchema().
					WithNullable()).
				WithProperty("query", openapi3.NewStringSchema().
					WithNullable())),
		"ErrorLogs": openapi3.NewSchemaRef("",
			openapi3.NewArraySchema().
				WithPropertyRef("errLog", &openapi3.SchemaRef{
					Ref: "#/components/schemas/ErrorLog",
				})),
		"Message": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithPropertyRef("errorLog", &openapi3.SchemaRef{
					Ref: "#/components/schemas/ErrorLogs",
				}).
				WithProperty("name", openapi3.NewStringSchema().
					WithNullable()).
				WithPropertyRef("teams", &openapi3.SchemaRef{
					Ref: "#/components/schemas/TeamDataList",
				}).
				WithPropertyRef("message", &openapi3.SchemaRef{
					Ref: "#/components/schemas/Message",
				})),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CompetitionRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for fetching competition data").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("source", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("competitionID", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("competitionName", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("table", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("id", openapi3.NewStringSchema().
						WithMinLength(1))),
		},
		"CompetitionCrawlRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for crawling data by date").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("competitionID", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("date", openapi3.NewStringSchema().
						WithFormat("date-time"))),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"CompetitionResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response with competition data").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("id", openapi3.NewStringSchema().
						WithNullable()).
					WithProperty("name", openapi3.NewStringSchema()).
					WithNullable().
					WithPropertyRef("teams", &openapi3.SchemaRef{
						Ref: "#/components/schemas/TeamDataList",
					}).
					WithNullable().
					WithPropertyRef("message", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Message",
					}).
					WithNullable())),
		},
		"CompetitionCrawlResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating tasks.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("id", openapi3.NewStringSchema().
						WithNullable()).
					WithPropertyRef("message", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Message",
					}))),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/competition": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "GetCompetition",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CompetitionRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/CompetitionResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/CompetitionResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CompetitionResponse",
					},
				},
			},
		},
	}

	return swagger
}

func RegisterOpenAPI(r *mux.Router) {
	swagger := NewOpenAPI3()

	r.HandleFunc("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		renderResponse(w, &swagger, http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger)

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}
