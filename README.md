# CodeReview

## Introduction

In the fast-paced world of software development, constructive feedback and peer review are invaluable. **CodeReview** simplifies the code review process, allowing both anonymous and registered users to share, review, and improve code efficiently and effectively.


Key Features:

- **Ease of Use**: Share your code snippets or entire projects effortlessly. Paste your code onto our platform, and with just a few clicks, it's ready for review. No authentication required for posting, making it swift and straightforward for anyone to get feedback.

- **Privacy-Controlled Sharing**: Choose how you share your code. Make it public for the community to see and review, or share a private link with selected peers for targeted feedback. This flexibility ensures that you get the input you need while maintaining control over your work's visibility.

- **Line-by-Line Reviews**: Our platform enables detailed feedback down to specific lines of code. Reviewers can comment directly on lines that need attention, making it easier for the original poster to understand and apply suggestions.

- **Community-Driven Learning**: Whether you're sharing your code for review or contributing by reviewing others' code, CodeReview fosters a learning environment. Gain insights from different perspectives, improve coding practices, and contribute to a culture of knowledge sharing.

- **AI-Generated Code Review** (only for authorized users): Once you submit your code for review, our AI Bot will review your code and submit its feedback. Gain insights in seconds effortlessly.

## Tech-Stack

- GoLang 1.21
- Framework: GoFiber
- ORM: Gorm
- Database: Postgres 16

## Setup

### Docker

To run the app, simply use `docker-compose`:

```
docker-compose up --build
```

### Local

In order to run the app locally, make sure that you have the following:
- Postgres database is up and running
- `.env` file in the following format (please specify credentials of the database you created):
```bash
DB_HOST=localhost # postgres, if you are running it in Docker
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=codereview
```
- GoLang installed (1.21 version is used in the project)


After making sure, please run the following to install the dependencies:
```bash
go mod download
```


And then run the application with:
```bash
go run .
```