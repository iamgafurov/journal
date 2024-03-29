// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/academic_years": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get academic years",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AcademicYearsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AcademicYearsPayload"
                        }
                    }
                }
            }
        },
        "/attendance_journal/get": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get attendance journal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GetJournalRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetAttendanceJournalPayload"
                        }
                    }
                }
            }
        },
        "/attendance_journal/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update attendance journal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateAttendanceJournalRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AttendanceJournalError"
                        }
                    }
                }
            }
        },
        "/courses/at": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get group courses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupCoursesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GroupCoursesPayload"
                        }
                    }
                }
            }
        },
        "/courses/pt": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get group courses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupCoursesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GroupCoursesPayload"
                        }
                    }
                }
            }
        },
        "/faculties": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get user faculties, specialities, years and groups",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.FacultiesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.MainFilterPayload"
                        }
                    }
                }
            }
        },
        "/ping": {
            "post": {
                "description": "ping server",
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/point_journal/get": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get poins journal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GetJournalRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.PointJournal"
                        }
                    }
                }
            }
        },
        "/point_journal/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update poins journal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePointJournalRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/tokenize": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Tokenize user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TokenizeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TokenizePayload"
                        }
                    }
                }
            }
        },
        "/topic/all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all topics",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TopicAllRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetTopicsPayload"
                        }
                    }
                }
            }
        },
        "/topic/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create topic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TopicUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Topic"
                        }
                    }
                }
            }
        },
        "/topic/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete topic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TopicDeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/topic/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update topic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TopicUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/untokenize": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete user token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "Service",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AcademicYearsPayload": {
            "type": "object",
            "properties": {
                "academic_years": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.AcademicYearsRequest": {
            "type": "object",
            "properties": {
                "external_ref": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.Attendance": {
            "type": "object",
            "properties": {
                "topic_number": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.AttendanceJournalError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "student_id": {
                    "type": "integer"
                }
            }
        },
        "dto.Course": {
            "type": "object",
            "properties": {
                "attendance_id": {
                    "type": "integer"
                },
                "course_name": {
                    "type": "string"
                },
                "credits_count": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "is_assistant": {
                    "type": "boolean"
                },
                "point_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "teacher_name": {
                    "type": "string"
                }
            }
        },
        "dto.DeleteTokenRequest": {
            "type": "object",
            "properties": {
                "external_ref": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.FacultiesRequest": {
            "type": "object",
            "properties": {
                "academic_year": {
                    "type": "string"
                },
                "external_ref": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                }
            }
        },
        "dto.Faculty": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "specialties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Speciality"
                    }
                }
            }
        },
        "dto.GetAttendanceJournalPayload": {
            "type": "object",
            "properties": {
                "journal": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.StudentAttendance"
                    }
                }
            }
        },
        "dto.GetJournalRequest": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "external_ref": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.GetTopicsPayload": {
            "type": "object",
            "properties": {
                "topics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Topic"
                    }
                }
            }
        },
        "dto.Group": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.GroupCoursesPayload": {
            "type": "object",
            "properties": {
                "courses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Course"
                    }
                }
            }
        },
        "dto.GroupCoursesRequest": {
            "type": "object",
            "properties": {
                "academic_year": {
                    "type": "string"
                },
                "external_ref": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.LoginPass": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.MainFilterPayload": {
            "type": "object",
            "properties": {
                "faculties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Faculty"
                    }
                }
            }
        },
        "dto.PointJournal": {
            "type": "object",
            "properties": {
                "current_week": {
                    "type": "integer"
                },
                "header": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Week"
                    }
                },
                "max_point": {
                    "type": "integer"
                },
                "students": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.StudentPoint"
                    }
                }
            }
        },
        "dto.PointUpdate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "point": {
                    "type": "number"
                }
            }
        },
        "dto.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.Speciality": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "years": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Year"
                    }
                }
            }
        },
        "dto.StudentAttendance": {
            "type": "object",
            "properties": {
                "attendance": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Attendance"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "record_book": {
                    "type": "string"
                }
            }
        },
        "dto.StudentPoint": {
            "type": "object",
            "properties": {
                "first_rating": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.WeekPoint"
                    }
                },
                "first_rating_sum": {
                    "type": "number"
                },
                "grade": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "points_sum": {
                    "type": "number"
                },
                "record_book": {
                    "type": "string"
                },
                "second_rating": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.WeekPoint"
                    }
                },
                "second_rating_sum": {
                    "type": "number"
                }
            }
        },
        "dto.TokenizePayload": {
            "type": "object",
            "properties": {
                "expire_at": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_code": {
                    "type": "integer"
                }
            }
        },
        "dto.TokenizeRequest": {
            "type": "object",
            "properties": {
                "external_ref": {
                    "type": "string"
                },
                "login_pass": {
                    "$ref": "#/definitions/dto.LoginPass"
                },
                "service_name": {
                    "type": "string"
                }
            }
        },
        "dto.TopicAllRequest": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "external_ref": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.TopicDeleteRequest": {
            "type": "object",
            "properties": {
                "external_ref": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "topicId": {
                    "type": "integer"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.TopicUpdateRequest": {
            "type": "object",
            "properties": {
                "external_ref": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "topic": {
                    "$ref": "#/definitions/models.Topic"
                },
                "topic_id": {
                    "type": "integer"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateAttendanceJournalRequest": {
            "type": "object",
            "properties": {
                "attendance": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.StudentAttendance"
                    }
                },
                "course_id": {
                    "type": "integer"
                },
                "external_ref": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdatePointJournalRequest": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "external_ref": {
                    "type": "string"
                },
                "points": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PointUpdate"
                    }
                },
                "service_name": {
                    "type": "string"
                },
                "userUchprocCode": {
                    "type": "integer"
                }
            }
        },
        "dto.Week": {
            "type": "object",
            "properties": {
                "editable": {
                    "type": "boolean"
                },
                "number": {
                    "type": "integer"
                }
            }
        },
        "dto.WeekPoint": {
            "type": "object",
            "properties": {
                "point": {
                    "type": "number"
                },
                "week_number": {
                    "type": "integer"
                }
            }
        },
        "dto.Year": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Group"
                    }
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.Topic": {
            "type": "object",
            "properties": {
                "cnzap": {
                    "type": "string"
                },
                "dtzap": {
                    "type": "string"
                },
                "editable": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "kol_kmd": {
                    "type": "integer"
                },
                "kol_lab": {
                    "type": "integer"
                },
                "kol_lek": {
                    "type": "integer"
                },
                "kol_obsh": {
                    "type": "integer"
                },
                "kol_prak": {
                    "type": "integer"
                },
                "kol_sem": {
                    "type": "integer"
                },
                "tema": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{"http"},
	Title:       "API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
