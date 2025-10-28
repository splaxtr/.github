<div align="right">

[![English](https://img.shields.io/badge/Language-English-blue?style=flat-square)](CONTRIBUTING.md)
[![Türkçe](https://img.shields.io/badge/Dil-Türkçe-red?style=flat-square)](CONTRIBUTING.tr.md)

</div>

# Contributing

First off, thank you for considering contributing to my projects! 🎉

Your contributions help make these projects better for everyone.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
- [Development Setup](#development-setup)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Commit Guidelines](#commit-guidelines)
- [Testing Guidelines](#testing-guidelines)

---

## 📜 Code of Conduct

This project and everyone participating in it is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

---

## 🤝 How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates.

**How to write a good bug report:**
- Use the [bug report template](.github/ISSUE_TEMPLATE/bug_report.yml)
- Provide a clear title and description
- Include steps to reproduce
- Add screenshots/logs if applicable
- Specify your environment (OS, browser, version)

### Suggesting Features

We love feature suggestions!

**How to write a good feature request:**
- Use the [feature request template](.github/ISSUE_TEMPLATE/feature_request.yml)
- Explain the problem you're trying to solve
- Describe your proposed solution
- Consider alternatives
- Assess impact on existing users

### Code Contributions

**Good first issues:**
Look for [`good first issue`](https://github.com/search?q=user%3Asplaxtr+label%3A%22good+first+issue%22+state%3Aopen&type=Issues) label - great for newcomers!

**Contribution steps:**
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Write/update tests
5. Run tests (`npm test` or equivalent)
6. Commit your changes (follow [commit guidelines](#commit-guidelines))
7. Push to your fork (`git push origin feature/amazing-feature`)
8. Open a Pull Request

---

## 🛠️ Development Setup

### Prerequisites

Depending on the project:
- **Node.js**: 18+ (for JavaScript/TypeScript projects)
- **Python**: 3.11+ (for Python projects)
- **Go**: 1.21+ (for Go projects)
- **Flutter**: 3.13+ (for Flutter projects)
- **Docker**: 20+ (optional, for containerized development)
- **Git**: 2.30+

### Getting Started

```bash
# Clone the repository
git clone https://github.com/splaxtr/[project-name].git
cd [project-name]

# Install dependencies
npm install  # or: pip install -r requirements.txt, go mod download, flutter pub get

# Run tests
npm test     # or: pytest, go test ./..., flutter test

# Start development
npm run dev  # or project-specific command
```

---

## 🔄 Pull Request Process

### Before Submitting

1. **Update documentation** - README, API docs, etc.
2. **Write tests** - Maintain 80%+ code coverage
3. **Run linters** - `npm run lint` or equivalent
4. **Run all tests** - Ensure everything passes
5. **Update CHANGELOG** - Add entry under "Unreleased"
6. **Self-review** - Review your own code first

### PR Requirements

- [ ] Clear title (following conventional commits)
- [ ] Description explains what and why
- [ ] All tests passing
- [ ] Code coverage maintained/improved
- [ ] Documentation updated
- [ ] No merge conflicts
- [ ] Approved by at least 1 reviewer

### PR Title Convention

```
<type>(<scope>): <subject>

Examples:
feat(auth): add OAuth2 support
fix(api): resolve memory leak in WebSocket
docs(readme): update installation instructions
refactor(core): improve performance
test(utils): add unit tests for helpers
```

**Types:**
- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation only
- `style` - Formatting, missing semicolons, etc.
- `refactor` - Code restructuring
- `perf` - Performance improvement
- `test` - Adding/updating tests
- `chore` - Maintenance tasks

---

## 💻 Coding Standards

### General Principles

- **Keep it simple** - Prefer clarity over cleverness
- **DRY** - Don't Repeat Yourself
- **YAGNI** - You Aren't Gonna Need It
- **Write tests** - Test your code thoroughly
- **Document** - Explain the "why", not just the "what"

### JavaScript/TypeScript

**Style Guide:** Follow [Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript)

```typescript
// ✅ Good
interface User {
  id: string;
  username: string;
  email: string;
}

async function fetchUser(userId: string): Promise<User> {
  const response = await api.get(`/users/${userId}`);
  return response.data;
}

// ❌ Bad
function getUser(id) {
  return api.get('/users/' + id).then(r => r.data);
}
```

**Rules:**
- Use TypeScript (no `any` type)
- Use `async/await` over `.then()`
- Use `const` over `let`, never `var`
- Use arrow functions for callbacks
- Run ESLint and Prettier

### Python

**Style Guide:** Follow [PEP 8](https://peps.python.org/pep-0008/)

```python
# ✅ Good
def fetch_user(user_id: str) -> User:
    """Fetch user by ID.

    Args:
        user_id: The user's ID

    Returns:
        User: The user object

    Raises:
        UserNotFound: If user doesn't exist
    """
    response = api.get(f'/users/{user_id}')
    return User.from_dict(response.json())

# ❌ Bad
def getuser(id):
    return User.from_dict(api.get('/users/'+id).json())
```

### Go

**Style Guide:** Follow [Effective Go](https://golang.org/doc/effective_go.html)

```go
// ✅ Good
func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
    if err := s.validator.Validate(req); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    // ... implementation
}

// ❌ Bad
func createUser(r *CreateUserRequest) *User {
    // Missing context, error handling, validation
}
```

---

## 📝 Commit Guidelines

### Conventional Commits

We follow [Conventional Commits](https://www.conventionalcommits.org/).

**Format:**
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Example:**
```
feat(auth): add OAuth2 login flow

Implement OAuth2 authorization code flow with
Google, GitHub and Microsoft providers.

- Add OAuth2 client configuration
- Implement token exchange
- Add user profile fetching
- Update login UI

Closes #123
```

**Rules:**
- Use present tense ("add feature" not "added feature")
- Use imperative mood ("move cursor" not "cursor moves")
- Don't capitalize first letter
- No period at the end
- Reference issues/PRs in footer

---

## 🧪 Testing Guidelines

### Test Coverage

- **Minimum**: 80% code coverage
- **Target**: 90%+ for critical code

### Test Pyramid

```
       /\
      /  \     E2E (10%)
     /____\
    /      \   Integration (30%)
   /________\
  /          \ Unit (60%)
 /____________\
```

### Writing Tests

**JavaScript/TypeScript:**
```typescript
describe('fetchUser', () => {
  it('should fetch user successfully', async () => {
    // Arrange
    const userId = '123';
    mockApi.get.mockResolvedValue({ data: mockUser });

    // Act
    const user = await fetchUser(userId);

    // Assert
    expect(user).toEqual(mockUser);
    expect(mockApi.get).toHaveBeenCalledWith('/users/123');
  });
});
```

**Python:**
```python
def test_fetch_user():
    # Arrange
    user_id = "123"
    mock_api.get.return_value = MockResponse(mock_user)

    # Act
    user = fetch_user(user_id)

    # Assert
    assert user == mock_user
    mock_api.get.assert_called_with('/users/123')
```

---

## 🏷️ Issue Labels

We use labels to organize and prioritize issues:

**Type:**
- `type: bug` - Something isn't working
- `type: feature` - New feature request
- `type: docs` - Documentation improvement
- `type: refactor` - Code restructuring
- `type: performance` - Performance improvement

**Priority:**
- `priority: critical` - Urgent, breaks core functionality
- `priority: high` - Important, should be done soon
- `priority: medium` - Normal priority
- `priority: low` - Nice to have

**Status:**
- `status: needs-triage` - Needs initial review
- `status: needs-info` - Waiting for more information
- `status: blocked` - Blocked by another issue
- `status: in-progress` - Currently being worked on
- `status: needs-review` - Ready for review

**Special:**
- `good first issue` - Good for newcomers
- `help wanted` - Community help needed
- `breaking change` - Introduces breaking changes
- `security` - Security-related issue

---

## 🎯 Development Workflow

### Branch Strategy

```
main (production)
  ↑
develop (staging)
  ↑
feature/xxx (feature branches)
```

**Branch Naming:**
- `feature/add-oauth` - New features
- `fix/memory-leak` - Bug fixes
- `docs/api-guide` - Documentation
- `refactor/auth-service` - Refactoring
- `hotfix/security-patch` - Emergency fixes

---

## ❓ Questions?

- 📧 Check the repository's README for contact information
- 💬 Open a [Discussion](https://github.com/splaxtr/.github/discussions)
- 🐛 For bugs, use [Issues](https://github.com/splaxtr/.github/issues)

---

## 🙏 Thank You!

Your contributions make these projects better for everyone. We appreciate your time and effort! ❤️

**Happy coding! 🚀**
