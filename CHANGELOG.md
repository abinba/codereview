# CHANGELOG

## Version 0.0.1

#### NOTE: the code snippet schema changed.

For creating the code snippet, there should be two requests: 

1) POST /api/v1/code_snippet -> Title, UserID (optional)
2) POST /api/v1/code_snippet_version -> CodeSnippetID, ProgramLanguageID, Text

This is done for being able to have versioning for the code snippets.

Features List:

- Code Snippets: Added GET /api/v1/user_code_snippets to retrieve user code snippets
- Code Snippets: Added POST /api/v1/review_comment to create reviews for code snippets
- Code Snippets: Added POST /api/v1/code_snippet_version to create a new version of the code snippet
- Authorization: Added POST /api/v1/register and POST /api/v1/login
- Authorization: Added JWT token generation and check for login / signup
- Authorization: Added validation of username and password
- Security: Added security middleware
- Database: Programming languages are inserted with the first migration
