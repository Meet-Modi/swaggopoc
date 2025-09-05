# Swaggo POC - Swagger Documentation Generator with Nested JSON/XML Structures

This is a proof of concept project demonstrating how to use Swaggo to automatically generate Swagger documentation for a Go REST API with complex nested data structures supporting both JSON and XML formats.

## Features

- RESTful API with CRUD operations for users with nested data structures
- **Dual format support**: Both JSON and XML request/response formats
- Automatic Swagger documentation generation for complex nested models
- Interactive Swagger UI
- Proper HTTP status codes and error handling
- Content negotiation based on Accept and Content-Type headers
- JSON and XML request/response validation with nested objects
- Dedicated endpoints for managing nested user data (address, preferences, profile)

## API Endpoints

### Main User Operations
- `GET /api/v1/health` - Health check endpoint
- `GET /api/v1/users` - Get all users with complete nested data
- `GET /api/v1/users/{id}` - Get user by ID with all nested structures
- `POST /api/v1/users` - Create a new user with nested data structures
- `PUT /api/v1/users/{id}` - Update an existing user (supports partial updates of nested data)
- `DELETE /api/v1/users/{id}` - Delete a user

### Nested Data Operations
- `GET /api/v1/users/{id}/address` - Get user's address information
- `PUT /api/v1/users/{id}/address` - Update user's address
- `GET /api/v1/users/{id}/preferences` - Get user's preferences (theme, notifications, privacy)
- `PUT /api/v1/users/{id}/preferences` - Update user's preferences
- `GET /api/v1/users/{id}/profile` - Get user's profile information
- `PUT /api/v1/users/{id}/profile` - Update user's profile

### Documentation
- `GET /swagger/index.html` - Swagger UI documentation

## Installation and Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Install Swag CLI tool (if not already installed):**
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

3. **Generate Swagger documentation:**
   ```bash
   swag init
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

## Usage

### Start the server
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Access Swagger UI
Open your browser and navigate to: `http://localhost:8080/swagger/index.html`

### Example API calls

#### Get all users (JSON format - default)
```bash
curl -X GET http://localhost:8080/api/v1/users
```

#### Get all users (XML format)
```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Accept: application/xml"
```

#### Create a new user with nested structures (JSON)
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "address": {
      "street": "789 Elm Street",
      "city": "Portland",
      "state": "OR",
      "postal_code": "97201",
      "country": "USA"
    },
    "contact": {
      "phone": "+1-555-111-2222",
      "alternate_email": "john.alt@example.com",
      "website": "https://johndoe.dev"
    },
    "profile": {
      "bio": "Frontend developer with React expertise",
      "avatar": "https://example.com/john-avatar.jpg",
      "skills": ["React", "TypeScript", "CSS", "Node.js"],
      "experience": 3
    },
    "preferences": {
      "theme": "dark",
      "language": "en",
      "notifications": {
        "email": true,
        "sms": false,
        "push": true,
        "marketing": false
      },
      "privacy": {
        "profile_visible": true,
        "show_email": false,
        "show_phone": false
      }
    }
  }'
```

#### Create a new user with nested structures (XML)
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/xml" \
  -d '<CreateUserRequest>
    <name>John Doe</name>
    <email>john@example.com</email>
    <address>
      <street>789 Elm Street</street>
      <city>Portland</city>
      <state>OR</state>
      <postal_code>97201</postal_code>
      <country>USA</country>
    </address>
    <contact>
      <phone>+1-555-111-2222</phone>
      <alternate_email>john.alt@example.com</alternate_email>
      <website>https://johndoe.dev</website>
    </contact>
    <profile>
      <bio>Frontend developer with React expertise</bio>
      <avatar>https://example.com/john-avatar.jpg</avatar>
      <skills>
        <skill>React</skill>
        <skill>TypeScript</skill>
        <skill>CSS</skill>
        <skill>Node.js</skill>
      </skills>
      <experience>3</experience>
    </profile>
    <preferences>
      <theme>dark</theme>
      <language>en</language>
      <notifications>
        <email>true</email>
        <sms>false</sms>
        <push>true</push>
        <marketing>false</marketing>
      </notifications>
      <privacy>
        <profile_visible>true</profile_visible>
        <show_email>false</show_email>
        <show_phone>false</show_phone>
      </privacy>
    </preferences>
  </CreateUserRequest>'
```

#### Get a specific user (JSON)
```bash
curl -X GET http://localhost:8080/api/v1/users/1
```

#### Get a specific user (XML)
```bash
curl -X GET http://localhost:8080/api/v1/users/1 \
  -H "Accept: application/xml"
```

#### Update a user (partial nested updates supported - JSON)
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "address": {
      "city": "San Francisco",
      "state": "CA"
    },
    "preferences": {
      "theme": "light",
      "notifications": {
        "email": false,
        "push": false
      }
    }
  }'
```

#### Update a user (XML)
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/xml" \
  -d '<UpdateUserRequest>
    <name>Jane Doe</name>
    <address>
      <city>San Francisco</city>
      <state>CA</state>
    </address>
    <preferences>
      <theme>light</theme>
      <notifications>
        <email>false</email>
        <push>false</push>
      </notifications>
    </preferences>
  </UpdateUserRequest>'
```

#### Get user's address (XML response)
```bash
curl -X GET http://localhost:8080/api/v1/users/1/address \
  -H "Accept: application/xml"
```

#### Update user's address (XML request)
```bash
curl -X PUT http://localhost:8080/api/v1/users/1/address \
  -H "Content-Type: application/xml" \
  -d '<Address>
    <street>456 New Street</street>
    <city>Seattle</city>
    <state>WA</state>
    <postal_code>98101</postal_code>
    <country>USA</country>
  </Address>'
```

#### Get user's preferences
```bash
curl -X GET http://localhost:8080/api/v1/users/1/preferences
```

#### Update user's preferences (JSON)
```bash
curl -X PUT http://localhost:8080/api/v1/users/1/preferences \
  -H "Content-Type: application/json" \
  -d '{
    "theme": "dark",
    "language": "en",
    "notifications": {
      "email": true,
      "sms": true,
      "push": false,
      "marketing": false
    },
    "privacy": {
      "profile_visible": true,
      "show_email": true,
      "show_phone": false
    }
  }'
```

#### Delete a user
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Project Structure

```
.
├── main.go                 # Main application entry point with Swagger config
├── go.mod                  # Go module file
├── handlers/
│   └── users.go           # HTTP handlers with Swagger annotations for nested data
├── models/
│   └── user.go           # Data models with nested structures and Swagger annotations
└── docs/                 # Auto-generated Swagger documentation
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
```

## Data Model Structure

The User model now includes the following nested structures with both JSON and XML tag support:

### User (Root Object)
- **Basic Info**: ID, Name, Email, CreatedAt, UpdatedAt
- **Address**: Street, City, State, PostalCode, Country
- **Contact**: Phone, AlternateEmail, Website
- **Profile**: Bio, Avatar, Skills (array), Experience
- **Preferences**: Theme, Language, Notifications, Privacy

### Nested Objects
- **NotificationSettings**: Email, SMS, Push, Marketing (all boolean)
- **PrivacySettings**: ProfileVisible, ShowEmail, ShowPhone (all boolean)

### Content Negotiation
The API supports both JSON and XML formats:
- **Request Format**: Determined by `Content-Type` header (`application/json` or `application/xml`)
- **Response Format**: Determined by `Accept` header (`application/json` or `application/xml`)
- **Default**: JSON format when no specific headers are provided

## Swagger Annotations for Nested Structures

This project demonstrates the following key Swaggo annotations for nested JSON/XML:

- `@Summary` - Brief description of the endpoint
- `@Description` - Detailed description
- `@Tags` - Groups endpoints in the UI
- `@Accept` - Content-Type that the endpoint accepts (supports `json,xml`)
- `@Produce` - Content-Type that the endpoint produces (supports `json,xml`)
- `@Param` - Parameter description (path, query, body) - supports nested objects
- `@Success` - Success response description with nested object references
- `@Failure` - Error response description
- `@Router` - Route path and HTTP method
- `// @name` - Custom naming for nested structures in Swagger UI

### Key Features for Nested Data:
- **Partial Updates**: The PUT endpoints support partial updates of nested structures
- **Dedicated Endpoints**: Separate endpoints for managing specific nested data
- **Dual Format Support**: All endpoints support both JSON and XML
- **Proper Validation**: JSON and XML binding validation for nested structures
- **Rich Examples**: Comprehensive example data for all nested fields in both formats
- **Content Negotiation**: Automatic format detection based on HTTP headers

## Regenerating Documentation

Whenever you make changes to the Swagger annotations in your code, regenerate the documentation:

```bash
swag init
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [Swaggo](https://github.com/swaggo/swag) - Swagger documentation generator
- [gin-swagger](https://github.com/swaggo/gin-swagger) - Gin middleware for Swagger
- [files](https://github.com/swaggo/files) - Static file serving

## Notes

- The project uses in-memory storage for demonstration purposes
- In a real application, you would replace this with a proper database
- The API includes proper error handling and validation for nested structures in both JSON and XML formats
- All models include comprehensive example values for better Swagger documentation
- Supports partial updates of nested data structures
- Demonstrates best practices for documenting complex JSON/XML APIs with Swagger
- The nested structure showcases real-world scenarios like user profiles, preferences, and contact information
- **Content negotiation**: Automatically detects and responds with the appropriate format (JSON/XML) based on HTTP headers
- **Backward compatibility**: JSON remains the default format for clients that don't specify content type
